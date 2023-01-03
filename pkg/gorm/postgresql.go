package gorm

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strings"

	e "github.com/deall-users/pkg/error"

	"gorm.io/gorm"
)

type DbPgSql struct {
	Db *gorm.DB
}

func (DbPgSql) PositionAt(activity string) string {
	return "Db." + activity
}

func (d DbPgSql) SQLDB() (*sql.DB, error) {
	return d.Db.DB()
}

func (d *DbPgSql) Close() error {
	sqlDB, err := d.Db.DB()
	if err != nil {
		return e.WrapErrorf(err, http.StatusInternalServerError, d.PositionAt("sqlDB"))
	}

	return sqlDB.Close()
}

func (d *DbPgSql) WrapErrorGorm(ctx context.Context, err error, location string) error {
	switch {
	case err == nil:
		return nil
	case errors.Is(err, context.DeadlineExceeded):
		return e.ErrTimeout
	case errors.Is(err, gorm.ErrRecordNotFound):
		return e.ErrDataNotFound
	case strings.Contains(err.Error(), "SQLSTATE 23503"): // handle fk violation
		return e.ErrRelationNotFound
	case strings.Contains(err.Error(), "SQLSTATE 23505"): // handle duplicate key violation
		return e.ErrDataAlreadyExist
	default:
		return e.WrapErrorf(err, http.StatusInternalServerError, d.PositionAt(location))
	}
}
