package db

import (
	"diagnofish/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct{}

func (p *Postgres) Connect(creds *model.Credential) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", creds.Host, creds.Username, creds.Password, creds.DatabaseName, creds.Port)

	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

func NewDB() *Postgres {
	return &Postgres{}
}

// Reset table on testing
// func (p *Postgres) Reset(db *gorm.DB, table string) error {
// 	return db.Transaction(func(tx *gorm.DB) error {
// 		if err := tx.Exec("TRUNCATE " + table).Error; err != nil {
// 			return err
// 		}

// 		if err := tx.Exec("ALTER SEQUENCE " + table + "_id_seq RESTART WITH 1").Error; err != nil {
// 			return err
// 		}

// 		return nil
// 	})
// }
