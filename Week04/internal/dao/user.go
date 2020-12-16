package dao

import (
	"database/sql"
	"fmt"

	"../model"

	"github.com/pkg/errors"
)

type dao struct {
	db *sql.DB
}

func NewDao(db *sql.DB) Dao {
	return &dao{db: db}
}

const (
	DB_NAME = "mysql_database_name"
	DB_USER = "mysql_user"
	DB_PASS = "mysql_password"
)

type Dao interface {
	UserById(id int) (*model.User, error)
}

func NewDB() (db *sql.DB, cleanup func(), err error) {
	db, err = sql.Open("mysql", fmt.Sprintf("%s/%s/%s", DB_NAME, DB_USER, DB_PASS))
	cleanup = func() {
		if err == nil {
			db.Close()
		}
	}
	return db, cleanup, nil
}

func (d *dao) UserById(id int) (*model.User, error) {
	row := d.db.QueryRow("SELECT `id` FROM `users` WHERE id=?", id)
	user := new(model.User)
	err := row.Scan(user.Id)
	if err == sql.ErrNoRows {
		return nil, errors.Wrap(errors.New("Not Found"), "No User")
	}
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get User")
	}
	return user, nil
}
