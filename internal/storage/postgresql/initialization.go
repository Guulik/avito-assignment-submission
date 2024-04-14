package postgresql

import (
	"Avito_trainee_assignment/config"
	"fmt"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func InitPostgres(c *config.Config) (*sqlx.DB, error) {
	connectionUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Postgres.Host, c.Postgres.SQLPort, c.Postgres.User, c.Postgres.Password, c.Postgres.DBName, c.Postgres.SslMode)

	db, err := sqlx.Connect(c.Postgres.Driver, connectionUrl)
	if err != nil {
		fmt.Println(err)
	}
	return db, err
}

func CreateTable(db *sqlx.DB) error {
	var (
		query = `
		CREATE TABLE IF NOT EXISTS "banner"
		(
			banner_id       serial       not null primary key,
			feature_id   	bigint       not null,
			tag_ids   		bigint[]	 not null,
			content 		JSON		 not null,
			is_active  		boolean 	 not null default true,
			created_at      timestamp	 not null,
			updated_at      timestamp	 not null
		);
		CREATE INDEX IF NOT EXISTS index_id ON banner (banner_id);
		`
		queryRelation = `
		CREATE TABLE IF NOT EXISTS "banner_definition" (
		banner_id BIGINT NOT NULL,
		feature_id BIGINT NOT NULL,
		tag_id BIGINT NOT NULL,
		PRIMARY KEY (feature_id, tag_id)
		);
		ALTER TABLE banner_definition DROP CONSTRAINT IF EXISTS fk_banner_id;

		ALTER TABLE banner_definition
		ADD CONSTRAINT fk_banner_id
		FOREIGN KEY (banner_id) REFERENCES "banner" (banner_id)
		ON DELETE CASCADE;

		CREATE UNIQUE INDEX IF NOT EXISTS index_feature_tag ON banner_definition (feature_id, tag_id);
		`
	)
	if _, err := db.Exec(query); err != nil {
		return err
	}
	if _, err := db.Exec(queryRelation); err != nil {
		return err
	}

	return nil
}
