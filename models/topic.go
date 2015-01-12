package models

import (
     "time"
     "strconv"
     "github.com/astaxie/beego"
)

type Topic struct {
     Id                int64
     Cid               int64
     Title             string
     Category          string
     Content           string  `xorm: text`
     Attachment        string
     Created           time.Time
     Updated           time.Time
     Views             int64
     Author            string
     ReplyTime         time.Time
     ReplyCount        int64
}

func init() {
   tables = append(tables,new(Topic))
}

func AddTopic(title, category,content,attachment string) error {
     topic := &Topic{
           Title:         title,
           Category:      category,
           Content:       content,
           Attachment:    attachment,
           Created:       time.Now(),
           Updated:       time.Now(),
     }

    _, err := x.Insert(topic)
    if err != nil {
       return err
    }

    cate := &Category{Title: category}
    has, err := x.Get(cate)
    if err != nil {
       return nil
    }

    if has {
       cate.TopicCount++
       _, err = x.Update(cate)
    }
    return err
}

func GetTopic(tid string)(*Topic ,error) {
   tidNum, err := strconv.ParseInt(tid,10,64)
   if err != nil {
        return nil, err
   }

   topic := &Topic{Id: tidNum}
   has, err := x.Get(topic)
   if err != nil {
      return nil, err
   }

   if !has {
      return nil, err
   }

   topic.Views++
   _, err = x.Update(topic)

  return topic, nil
}

func ModifyTopic(tid,title,category,content,attachment string) error {
     tidNum, err := strconv.ParseInt(tid,10,64)
     if err != nil {
        return err
     }

    var oldcate, oldattach string
    topic := &Topic{Id: tidNum}
    has, err := x.Get(topic)
    if err != nil {
       return err
    }

    if has {
       oldcate = topic.Category
       oldattach = topic.Attachment
       topic.Title = title
       topic.Category = category
       topic.Content = content
       topic.Attachment = attachment
       topic.Updated = time.Now()
       _, err = x.Update(topic) 
       if err != nil {
         return err
       }
    }

    if len(oldcate) > 0 {
       cate := &Category{Title: oldcate}
       has, err = x.Get(cate)
       if err != nil {
         return err
       }
       if has {
             cate.TopicCount--
             _,err = x.Update(cate)
       }
    }

   cate := &Category{Title: category}
   has, err = x.Get(cate)
   beego.Debug(oldattach)
   if err != nil {
      return err
   }

   if has {
     cate.TopicCount++
     _, err = x.Update(cate)
   }

   return nil
}

func DeleteTopic(tid string) error {
     tidNum, err := strconv.ParseInt(tid,10,64)
     if err != nil {
        return err
     }
     var oldcate string
     topic := &Topic{Id: tidNum}
     has, err := x.Get(topic)
     if err != nil {
        return err
     }

     if has {
        oldcate = topic.Category
        x.Delete(topic)
     }

     if len(oldcate) > 0 {
        cate := &Category{Title: oldcate}
        has, err := x.Get(cate)
        if err != nil {
           return err
        }
        if has {
           cate.TopicCount--
           _, err = x.Update(cate)
        }
     }
     return nil
}

func GetAllTopics(category string,ishome bool)([]*Topic, error) {
    var err error
    topics := make([]*Topic, 0)
    if ishome {
      err = x.OrderBy("Created").Find(&topics)
    } else {
      err = x.Find(&topics)
    }
    return topics, err
}
