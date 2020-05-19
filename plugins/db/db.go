package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"strconv"
)
const (
	//mysql Config
	MysqlHost     = "114.55.100.116"
	MysqlPort     = 3306
	MysqlDBName   = "ree"
	MysqlUserName = "root"
	MysqlPassword = "123456"
)


func GetEngine() (*xorm.Engine, error) {
	dburl := MysqlUserName + ":" + MysqlPassword + "@tcp(" + MysqlHost + ":" + strconv.Itoa(MysqlPort) + ")/" + MysqlDBName + "?charset=utf8"
	orm, err := xorm.NewEngine("mysql", dburl)
	if err != nil {
		return nil, err
	}
	orm.ShowSQL(true)
	return orm, nil
}
