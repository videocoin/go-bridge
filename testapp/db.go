package testapp

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Open(url string) (*DB, error) {
	sdb, err := gorm.Open("sqlite3", url)
	if err != nil {
		return nil, err
	}
	// sqlite's locking protocol allows for many readers and exactly one writer at any given time
	// https://github.com/mattn/go-sqlite3/issues/274
	// using single connection is still fast
	sdb.DB().SetMaxOpenConns(1)
	db := &DB{db: sdb}
	db.Migrate()
	return db, nil
}

type DB struct {
	db *gorm.DB
}

func (db *DB) Migrate() {
	db.db.AutoMigrate(&Transfer{})
}

func (db *DB) SaveTransfer(tr *Transfer) error {
	for _, err := range db.db.Save(tr).GetErrors() {
		if err != nil {
			return err
		}
	}
	return nil
}
