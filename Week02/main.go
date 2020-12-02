package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

type User struct {
	Name string
}

func (u *User) Dao(Name string) (*User, error) {
	//select user from DB...
	err := sql.ErrNoRows
	return nil, errors.Errorf("No user info with error userName%d,err=%v", Name, err)
}

func Biz(Name string) (*User, error) {
	u := &User{Name}
	if _, err := u.Dao(Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.WithMessage(err, "no user found")
		}
		//其他错误
		if err != nil {
			return nil, errors.Wrap(err, " sql error")
		}
	}

	return u, nil
}

func main() {
	name := "Dylan"
	u, err := Biz(name)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("data info: %+v\n", u)
}
