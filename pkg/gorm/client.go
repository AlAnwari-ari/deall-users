package gorm

import (
	"strings"
	"time"

	cfg "github.com/deall-users/config"
	"github.com/deall-users/pkg/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresClient() (*DbPgSql, error) {
	db, err := pgConnection()
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return &DbPgSql{db}, nil
}

func pgConnection() (*gorm.DB, error) {
	dsn := "host=" + cfg.DB_HOST +
		" user=" + cfg.DB_USER +
		" dbname=" + cfg.DB_NAME +
		" sslmode=disable password=" + cfg.DB_PASSWORD +
		" port=" + cfg.DB_PORT +
		" connect_timeout=" + cfg.DB_TIMEOUT +
		" TimeZone=" + cfg.TIME_LOCATION

	// customization gorm
	gormConfig := &gorm.Config{
		NowFunc: utils.TimeNow,
		// Logger:  logger.Default.LogMode(logger.Info),
	}

	if !strings.EqualFold(cfg.RUN_MODE, "debug") {
		gormConfig.Logger = logger.Default.LogMode(logger.Silent)
	}

	return gorm.Open(postgres.Open(dsn), gormConfig)
}
