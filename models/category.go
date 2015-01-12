package models

import (
     "time"
     "strconv"
)

type Category struct {
     Id             int64
     Title          string
     Created        time.Time
     Views          int64
     TopicTime      time.Time
     TopicCount     int64
}


func AddCategory(name string) error {
    cate := &Category{Title: name}
    has, err := x.Get(cate)
    if err != nil {
       return err
    }
    if !has {
       _, err = x.Insert(cate)
       if err != nil {
          return err
       }
    }
    return nil
}

func DeleteCategory(id string) error {
   cid, err := strconv.ParseInt(id,10,64)
   if err != nil {
      return err
   }
   cate := &Category{Id: cid}
   _, err = x.Delete(cate)
   if err != nil {
      return err
   }
   return nil
}

func GetAllCategories() ([]*Category, error) {
   cates := make([]*Category, 0)
   err := x.Find(&cates)
   return cates, err
}

func init() {
   tables = append(tables,new(Category))
}
