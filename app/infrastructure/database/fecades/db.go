package fecades

import (
	"clean-golang/app/infrastructure/database/mysql"
)

func NewDb(s string) error {
	if s == "mysql" {
		q := mysql.NewMysql()
		_, err := q.Make()
		if err != nil {
			return err
		}
	}
	return nil
}
