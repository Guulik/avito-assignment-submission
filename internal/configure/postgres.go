package configure

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"net/url"
)

func NewPostgres(c *Config) *sqlx.DB {
	db := sqlx.MustConnect(c.Postgres.Driver, c.connectionString())
	return db
}

func (c *Config) MigrateUp(url ...string) error {
	var sourceURL string
	if url == nil {
		sourceURL = "file://internal/migrations/up"
	} else {
		sourceURL = url[0]
	}
	fmt.Println(c.connectionString())
	m, err := migrate.New(sourceURL, c.connectionString())
	if err != nil {
		return err
	}
	if err = m.Up(); err != nil {
		return err
	}

	return nil
}

func (c *Config) connectionString() string {
	u := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(c.Postgres.User, c.Postgres.Password),
		Host:   fmt.Sprintf("%s:%d", c.Postgres.Host, c.Postgres.SQLPort),
		Path:   c.Postgres.DBName,
	}

	q := u.Query()
	q.Set("sslmode", "disable")

	u.RawQuery = q.Encode()

	return u.String()
}
