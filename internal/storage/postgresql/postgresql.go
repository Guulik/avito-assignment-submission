package postgresql

import (
	"Avito_trainee_assignment/internal/config/constants"
	"Avito_trainee_assignment/internal/domain/model"
	sl "Avito_trainee_assignment/internal/lib/logger/slog"
	"Avito_trainee_assignment/internal/storage"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"strings"
	"time"
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

func (s Storage) UserBannerDB(featureId int, tagId int) ([]byte, error) {
	const op = "Repo.UserBannerDB"

	log := s.log.With(
		slog.String("op", op),
	)
	log.Info(fmt.Sprintf("fID: %v, tID%v", featureId, tagId))
	var (
		data  []byte
		query = fmt.Sprintf(`
		SELECT content
		FROM %s 
		WHERE feature_id = $1 AND $2 = ANY (tag_ids) AND is_active = true;
		`, constants.BannerTable)

		values = []any{featureId, tagId}
	)

	log.Info(fmt.Sprintf("sql query: %v", query))

	if err := s.db.Get(&data, query, values...); err != nil {
		return nil, err
	}

	return data, nil
}

func (s Storage) UserBannerCached(featureId int, tagId int) ([]byte, error) {
	//TODO implement me
	return nil, storage.ErrNotFound
}

func (s Storage) Banners(featureId int, tagIg int, limit int, offset int) (*model.Banner, error) {
	//TODO implement me
	return nil, storage.ErrNotFound
}

func (s Storage) Save(featureId int, tagsIds []int, content []byte, isActive bool) (int, error) {
	const op = "Repo.Save"

	log := s.log.With(
		slog.String("op", op),
	)

	var (
		query = fmt.Sprintf(`
		INSERT INTO %s (feature_id, content, is_active, tag_ids, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING banner_id
		`, constants.BannerTable)

		values = []any{featureId, content, isActive,
			"{" + strings.Trim(strings.Replace(fmt.Sprint(tagsIds), " ", ", ", -1), "[]") + "}",
			time.Now(), time.Now(),
		}
		id int
	)
	log.Info(fmt.Sprintf("sql query: %v", query))

	log.Info("beginning transaction")
	tx, err := s.db.Begin()
	if err != nil {
		log.Error("failed to begin transaction", sl.Err(err))
		return -1, err
	}
	defer tx.Rollback()

	row := tx.QueryRow(query, values...)
	log.Info("row:", row)
	err = row.Scan(&id)

	if err != nil {
		log.Error("failed to scan row", sl.Err(err))
		return -1, err
	}

	log.Info("trying to commit transaction")
	err = tx.Commit()
	if err != nil {
		log.Error("failed to commit transaction", sl.Err(err))
	}
	return id, nil
}

func (s Storage) Patch(bannerId int, tagsId []int, featureId int, content []byte, isActive bool) error {
	//TODO implement me
	return storage.ErrSaveFail
}

func (s Storage) Delete(bannerId int) error {
	//TODO implement me
	return storage.ErrSaveFail
}
