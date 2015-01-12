package models

import (
    "github.com/astaxie/beego"
    _ "github.com/go-sql-driver/mysql"
    "github.com/go-xorm/xorm"
)

var (
   x *xorm.Engine
)

func init() {
  var err error
  datasource := beego.AppConfig.String("mysql")
  beego.Debug(datasource)
  if datasource == "" {
     datasource = "root:123456@/seraphim?loc=PRC"
  }
  x, err = xorm.NewEngine("mysql",datasource)
  if err != nil {
     beego.Error(err)
  }
}

var tables = []interface{}{}

func InitDB() (err error) {
    return x.Sync(tables...)
}
