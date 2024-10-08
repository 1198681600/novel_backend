package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type (
	// gormPgSqlConfig is the configuration for the GormPgSql driver.
	gormPgSqlConfig struct {
		DBName       string
		DSN          string
		MaxIdleConns int
		MaxOpenConns int
	}
)

var (
	DB *gorm.DB
)

func initDB() {
	config := &gormPgSqlConfig{
		DBName:       "backend",
		DSN:          viper.GetString("database.dsn"),
		MaxIdleConns: viper.GetInt("database.max_idle_conns"),
		MaxOpenConns: viper.GetInt("database.max_open_conns"),
	}
	DB = gormPgSql(config)

}

func gormPgSql(config *gormPgSqlConfig) *gorm.DB {
	pgsqlConfig := postgres.Config{DSN: config.DSN, PreferSimpleProtocol: false}
	if db, err := gorm.Open(postgres.New(pgsqlConfig), gormConfig()); err != nil {
		zap.L().Panic("failed to connect database", zap.Error(err))
		return nil
	} else {
		sqlDB, _ := db.DB()
		if config.MaxIdleConns != 0 {
			sqlDB.SetMaxIdleConns(config.MaxIdleConns)
		} else {
			sqlDB.SetMaxIdleConns(5)
		}

		if config.MaxOpenConns != 0 {
			sqlDB.SetMaxOpenConns(config.MaxOpenConns)
		} else {
			sqlDB.SetMaxOpenConns(10)
		}

		return db
	}
}

func gormConfig() *gorm.Config {
	return &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Info),
	}
}
