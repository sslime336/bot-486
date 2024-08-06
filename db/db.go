package db

import (
	"github.com/glebarez/sqlite"
	"github.com/sslime336/bot-486/config"
	"github.com/sslime336/bot-486/db/orm"
	"github.com/sslime336/bot-486/logging"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Sqlite *gorm.DB

func Init() {
	db, err := gorm.Open(sqlite.Open(config.App.Database.Sqlite.Path), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		logging.Fatal("failed to connect to database", zap.Error(err))
	}
	Sqlite = db
	if err := Sqlite.AutoMigrate(
		&orm.GroupImpart{}, &orm.Hunter{}, &orm.Impart{},
	); err != nil {
		logging.Fatal("migration failed", zap.Error(err))
	}
}
