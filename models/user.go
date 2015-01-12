package models

//import (
//  "github.com/astaxie/beego"
//)

type Users struct {
  Id          int64       `xorm:"autoincr"`
  Name        string
  Password    string
}


func init() {
    tables = append(tables, new(Users))
}

func (u *Users) Save() (err error) {
   _, err = x.Insert(u)
   return
}

func GetUserByName(name string) (*Users) {
   u := new(Users)
   isExist, err := x.Where("Name=?",name).Get(u)
   if err != nil {
      return nil
   }
   if !isExist {
      return nil
   }
   return u
}
