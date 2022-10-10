package fecades

import (
	"clean-golang/app/infrastructure/database/mongo"
	"clean-golang/app/infrastructure/database/mysql"
	"clean-golang/app/infrastructure/database/pgsql"
	"clean-golang/app/infrastructure/logger"
	"fmt"
)

func NewDb(s string) error {
	switch s {
	case "mysql":
		m := mysql.NewMysql()
		_, mErr := m.Make()
		if mErr != nil {
			return mErr
		}
	case "pgsql":
		p := pgsql.NewPgsql()
		_, pErr := p.Make()
		if pErr != nil {
			return pErr
		}
	case "mongo":
		m := mongo.NewMongo()
		_, mErr := m.Make()
		if mErr != nil {
			return mErr
		}
	default:
		textErr := fmt.Sprintf("%s is not database", s)
		logger.Error(textErr)
	}
	return nil
}
