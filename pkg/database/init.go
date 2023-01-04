package database

import (
	"fmt"
	"github.com/fat-garage/wordblock-backend/pkg/conf"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"log"

	"github.com/fat-garage/wordblock-backend/models"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

// Init .
func Init() {
	dbConf := conf.Cfg.Database
	var err error
	switch dbConf.Dialect {
	case "mysql":
		Db, err = mysqlInit()
	case "pgsql":
		Db, err = pgsqlInit()
	default:
		log.Fatalln("database dialect not support")
	}

	if err != nil {
		log.Fatalln(err)
	}

	if dbConf.Debug {
		Db = Db.Debug()
	}
	if dbConf.AutoMigrate {
		autoMigrate()
	}
}

func mysqlInit() (*gorm.DB, error) {
	dbConf := conf.Cfg.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Dbname)

	mysqlConfig := mysql.Config{
		DSN:                       dsn,  // DSN data source name
		DisableDatetimePrecision:  true, // Datetime precision is disabled. Databases before MySQL 5.6do not support it.
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}
	db, err := gorm.Open(mysql.New(mysqlConfig))
	if err != nil {
		return nil, errors.WithMessage(err, "mysql connect failed")
	}
	return db, nil
}

func pgsqlInit() (*gorm.DB, error) {
	dbConf := conf.Cfg.Database
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		dbConf.Host, dbConf.Port, dbConf.User, dbConf.Dbname, dbConf.Password, dbConf.SSLMode)

	pgsqlConfig := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), pgsqlConfig)
	if err != nil {
		return nil, errors.WithMessage(err, "pgsql connect failed")
	}
	return db, nil
}

func autoMigrate() {
	if err := Db.AutoMigrate(
		&models.Block{},
	); err != nil {
		log.Fatal("AutoMigrate error", err)
	}
	logrus.Println("All table AutoMigrate finish.")
}
