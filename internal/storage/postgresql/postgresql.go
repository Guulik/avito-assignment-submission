package postgresql

import (
	"Avito_trainee_assignment/internal/constants"
	"Avito_trainee_assignment/internal/domain/model"
	sl "Avito_trainee_assignment/internal/lib/logger/slog"
	"Avito_trainee_assignment/internal/storage"
	"fmt"
	"log/slog"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

var _ storage.BannerStorage = (*Storage)(nil)

type Storage struct {
	log *slog.Logger
	db  *sqlx.DB
}

func New(log *slog.Logger, db *sqlx.DB) *Storage {
	return &Storage{
		db:  db,
		log: log,
	}
}

const (
	SQLDuplicateError = "23505"
)

func (s Storage) UserBannerDB(featureId int64, tagId int64) ([]byte, error) {
	const op = "Repo.UserBannerDB"

	log := s.log.With(
		slog.String("op", op),
	)
	var (
		content []byte
		query   = fmt.Sprintf(`
		SELECT b.content
		FROM %s b
		JOIN %s bdef ON bdef.banner_id = b.banner_id
		WHERE bdef.feature_id = $1 AND $2 = bdef.tag_id AND b.is_active = true;`,
			constants.BannerTable,
			constants.BannerDefinitionTable,
		)

		values = []any{featureId, tagId}
	)

	log.Debug(fmt.Sprintf("sql query: %v", query))

	if err := s.db.Get(&content, query, values...); err != nil {
		log.Error("failed to get user banner", err)
		return nil, echo.NewHTTPError(http.StatusNotFound, "The banner for the user was not found")
	}

	return content, nil
}

func (s Storage) Banners(limit int64, offset int64) ([]model.BannerDB, error) {
	const op = "Repo.Banners"
	log := s.log.With(
		slog.String("op", op),
	)
	var (
		banners []model.BannerDB

		query = fmt.Sprintf(`
		SELECT *
		FROM %s`,
			constants.BannerTable)
	)
	if limit != -1 {
		query = query + fmt.Sprintf(" LIMIT %d", limit)
	}
	if offset != -1 {
		query = query + fmt.Sprintf(" OFFSET %d", offset)
	}
	log.Debug(fmt.Sprintf("query = %v", query))

	err := s.db.Select(&banners, query)
	if err != nil {
		log.Error("failed to SELECT banners", err)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return banners, nil
}

func (s Storage) FilteredBanners(featureId int64, tagIg int64, limit int64, offset int64) ([]model.BannerDB, error) {
	const op = "Repo.FilteredBanners"
	log := s.log.With(
		slog.String("op", op),
	)
	var (
		banners []model.BannerDB

		query = fmt.Sprintf(`
		SELECT b.*
		FROM %s b
		JOIN %s bdef ON bdef.banner_id = b.banner_id
		WHERE 
		`,
			constants.BannerTable,
			constants.BannerDefinitionTable)
	)
	switch {
	case featureId > -1 && tagIg > -1:
		query = query + fmt.Sprintf(" bdef.feature_id=%d AND %d=bdef.tag_id", featureId, tagIg)
	case featureId > -1:
		query = query + fmt.Sprintf(" bdef.feature_id=%d", featureId)
	case tagIg > -1:
		query = query + fmt.Sprintf(" %d=bdef.tag_id", tagIg)
	}
	query = query + " GROUP BY b.banner_id"
	if limit > -1 {
		query = query + fmt.Sprintf(" LIMIT %d", limit)
	}
	if offset > -1 {
		query = query + fmt.Sprintf(" OFFSET %d", offset)
	}
	log.Debug(fmt.Sprintf("query = %v", query))

	err := s.db.Select(&banners, query)
	if err != nil {
		log.Error("failed to SELECT banners with filters", err)
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return banners, nil
}

func (s Storage) Save(featureId int64, tagIds []int64, content []byte, isActive bool) (int64, error) {
	const op = "Repo.Save"

	log := s.log.With(
		slog.String("op", op),
	)

	var (
		query = fmt.Sprintf(`
		WITH inserted_banner AS (
			INSERT INTO %s (feature_id, content, is_active, tag_ids, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING banner_id
		)
		INSERT INTO %s (banner_id, feature_id, tag_id)
		SELECT
			banner_id,
			$1 AS feature_id,
			tag_id
		FROM
			inserted_banner
		CROSS JOIN
			UNNEST($4::bigint[]) AS tag_id
		RETURNING banner_id;`,
			constants.BannerTable,
			constants.BannerDefinitionTable)

		values = []any{
			featureId, content, isActive,
			"{" + strings.Trim(strings.Replace(fmt.Sprint(tagIds), " ", ", ", -1), "[]") + "}",
			time.Now(), time.Now(),
		}
		id int64
	)
	log.Debug("beginning transaction")
	tx, err := s.db.Begin()
	if err != nil {
		log.Error("failed to begin transaction", sl.Err(err))
		return -1, echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	defer tx.Rollback()

	row := tx.QueryRow(query, values...)
	err = row.Scan(&id)
	if err != nil {
		// TODO: find way to get sqlerror code explicitly
		if strings.Contains(err.Error(), "23505") {
			log.Error("Bad request", sl.Err(err))
			return -1, echo.NewHTTPError(http.StatusBadRequest,
				"The tag or feature of the banner overlaps with an existing banner")
		}
		log.Error("failed to scan row", sl.Err(err))
		return -1, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	log.Debug("trying to commit transaction")
	err = tx.Commit()
	if err != nil {
		log.Error("failed to commit transaction", sl.Err(err))
		return -1, echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return id, nil
}

func (s Storage) Patch(bannerId int64, tagIds []int64, featureId int64, content []byte, isActive bool) error {
	const op = "Repo.Patch"

	log := s.log.With(
		slog.String("op", op),
	)

	var (
		queryGetByID = fmt.Sprintf(`SELECT * 
		FROM %s 
		WHERE banner_id = $1;`,
			constants.BannerTable)

		queryUpdateFeatureID = fmt.Sprintf(`
		WITH updated_banner AS (
			UPDATE %s
			SET 
				feature_id = $2
			WHERE 
				banner_id = $1
			RETURNING banner_id
		)
		UPDATE 
			%s br
		SET 
			feature_id = $2
		FROM 
			updated_banner ub
		WHERE 
			br.banner_id = ub.banner_id;
		`,
			constants.BannerTable,
			constants.BannerDefinitionTable,
		)

		queryDeleteTagIDs = fmt.Sprintf(
			`
			DELETE FROM %s WHERE banner_id = $1;
			`,
			constants.BannerDefinitionTable,
		)

		queryUpdateTagIDs = fmt.Sprintf(`
		WITH updated_banner AS (
			UPDATE %s
			SET 
				tag_ids = $2
			WHERE 
				banner_id = $1
			RETURNING banner_id, feature_id
		)
		INSERT INTO 
			%s (banner_id, feature_id, tag_id)
		SELECT 
			ub.banner_id,
			ub.feature_id,
			tag_id
		FROM 
			updated_banner ub
		CROSS JOIN 
			UNNEST($2::bigint[]) AS tag_id;
		`,
			constants.BannerTable,
			constants.BannerDefinitionTable,
		)

		queryUpdateIsActive = fmt.Sprintf(`
		UPDATE %s SET is_active=$2 WHERE banner_id=$1;`,
			constants.BannerTable)

		queryUpdateContent = fmt.Sprintf(`
		UPDATE %s SET "content"=$2 WHERE banner_id=$1;`,
			constants.BannerTable)

		queryUpdatedTime = fmt.Sprintf(`
		UPDATE %s SET "updated_at"=$2 WHERE banner_id=$1;`,
			constants.BannerTable)
	)

	log.Debug("beginning transaction")
	tx, err := s.db.Begin()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	defer tx.Rollback()

	var contentJSON []byte
	var banner model.BannerDB
	log.Debug("trying to get banner")
	row := tx.QueryRow(queryGetByID, bannerId)
	err = row.Scan(
		&banner.ID,
		&banner.FeatureId,
		&banner.TagIds,
		&contentJSON,
		&banner.IsActive,
		&banner.CreatedAt,
		&banner.UpdatedAt,
	)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			log.Error("banner not found", sl.Err(err))
			return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Banner with id:%d not found", bannerId))
		}
		log.Error("failed to scan", sl.Err(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	log.Info("updating...")
	if featureId != banner.FeatureId {
		if _, err = tx.Exec(queryUpdateFeatureID, bannerId, featureId); err != nil {
			log.Error("failed to update feature", sl.Err(err))
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}
	bannerTags, err := model.ParseIntArrayFromString(banner.TagIds)
	if err != nil {
		log.Error("failed to parse tags", sl.Err(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if !slices.Equal(tagIds, bannerTags) {
		_, err = tx.Exec(queryDeleteTagIDs, bannerId)
		if err != nil {
			log.Error("failed to delete tags", sl.Err(err))
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		_, err = tx.Exec(
			queryUpdateTagIDs,
			bannerId,
			"{"+strings.Trim(strings.Replace(fmt.Sprint(tagIds), " ", ", ", -1), "[]")+"}",
		)
		if err != nil {
			log.Error("failed to update tags", sl.Err(err))
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}
	if isActive != banner.IsActive {
		if _, err = tx.Exec(queryUpdateIsActive, bannerId, isActive); err != nil {
			log.Error("failed to update active status", sl.Err(err))
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}
	if _, err = tx.Exec(queryUpdateContent, bannerId, content); err != nil {
		log.Error("failed to update content", sl.Err(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if _, err = tx.Exec(queryUpdatedTime, bannerId, time.Now()); err != nil {
		log.Error("failed to update updated_at", sl.Err(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	log.Info("update completed!")

	log.Debug("trying to commit transaction")
	if err = tx.Commit(); err != nil {
		log.Error("failed to commit transaction", sl.Err(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return nil
}

func (s Storage) Delete(bannerId int64, featureId int64, tagId int64) error {
	const op = "Repo.Delete"

	log := s.log.With(
		slog.String("op", op),
	)
	var (
		query = fmt.Sprintf(`
		DELETE FROM %s 
		WHERE
		`, constants.BannerTable)
		_ = fmt.Sprintf(`
		SELECT COUNT(*) 
		FROM %s
		WHERE
		`, constants.BannerTable)
	)

	var err error
	if bannerId > 0 {
		query = query + fmt.Sprintf(" banner_id = %d;", bannerId)
		if err = deleteBy(s.db, query); err != nil {
			log.Error("failed to delete by  Id", sl.Err(err))
			return err
		}
	}
	if featureId > 0 && tagId > 0 {

		query = query + fmt.Sprintf(" feature_id = %d AND %d = ANY(tag_ids);", featureId, tagId)
		if err = deleteBy(s.db, query); err != nil {
			log.Error("failed to delete by feature AND tag", sl.Err(err))
		}

	} else {
		if featureId > 0 {
			query = query + fmt.Sprintf(" feature_id = %d;", featureId)
			log.Info(query)
			if err = deleteBy(s.db, query); err != nil {
				log.Error("failed to delete by feature", sl.Err(err))
			}
		}
		if tagId > 0 {
			query = query + fmt.Sprintf(" %d = ANY(tag_ids);", tagId)
			log.Info(query)
			if err = deleteBy(s.db, query); err != nil {
				log.Error("failed to delete by tag", sl.Err(err))
			}
		}
	}
	return nil
}

func deleteBy(db *sqlx.DB, queryBanner string) error {
	log.Info("beginning transaction")

	tx, err := db.Begin()
	if err != nil {
		log.Error("failed to begin transaction", sl.Err(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	defer tx.Rollback()

	var affect int64

	log.Info("trying to delete banner")
	ct, err := tx.Exec(queryBanner)
	if err != nil {
		log.Error("failed to delete banner", sl.Err(err))
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if affect, err = ct.RowsAffected(); err != nil {
		return err
	}
	if affect < 1 {
		return echo.NewHTTPError(http.StatusNotFound, "nothing to delete")
	}

	log.Info(fmt.Sprintf("affected rows %v", affect))
	log.Info("trying to commit transaction")
	err = tx.Commit()
	if err != nil {
		log.Error("failed to commit transaction", sl.Err(err))
		return err
	}
	return nil
}

// GetBannerById is auxiliary function for quick access to the exact banner.
func (s Storage) GetBannerById(bannerId int64) (*model.BannerDB, error) {
	const op = "Repo.UserBannerDB"

	log := s.log.With(
		slog.String("op", op),
	)
	var (
		banner model.BannerDB
		query  = fmt.Sprintf(`
		SELECT *
		FROM %s
		WHERE banner_id = $1`,
			constants.BannerTable,
		)

		values = []any{bannerId}
	)

	log.Debug(fmt.Sprintf("sql query: %v", query))

	if err := s.db.Get(&banner, query, values...); err != nil {
		log.Error("failed to get user banner", err)
		return nil, echo.NewHTTPError(http.StatusNotFound, "Banner not found")
	}

	return &banner, nil
}
