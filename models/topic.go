package models

import (
     "os"
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

type Reply struct {
     Id             int64
     Tid            int64
     Name           string
     Content        string
     Created        time.Time
}

func init() {
   tables = append(tables,new(Topic),new(Reply))
}

func AddTopic(title,category,content,author,attachment string) error {
     topic := &Topic{
           Title:         title,
           Category:      category,
           Content:       content,
           Author:        author,
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
       _, err = x.Id(cate.Id).Update(cate)
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
   _, err = x.Id(topic.Id).Update(topic)

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
       _, err = x.Id(topic.Id).Update(topic)
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
             _,err = x.Id(cate.Id).Update(cate)
       }
    }

   if len(oldattach) >0 {
      os.Remove("static/attachment/"+oldattach)
   }
   cate := &Category{Title: category}
   has, err = x.Get(cate)
   beego.Debug(oldattach)
   if err != nil {
      return err
   }

   if has {
     cate.TopicCount++
     _, err = x.Id(cate.Id).Update(cate)
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
           _, err = x.Id(cate.Id).Update(cate)
        }
     }
     return nil
}

func GetAllTopics(category string,ishome bool)([]*Topic, error) {
    var err error
    topics := make([]*Topic, 0)
    if ishome {
      if len(category) > 0 {
         err := x.Where("category=?",category).OrderBy("-created").Find(&topics)
         return topics, err
      }
      err = x.OrderBy("-created").Find(&topics)
    } else {
      err = x.Find(&topics)
    }
    return topics, err
}

func GetAllTopicsByUser(name string) ([]*Topic, error) {
   topics := make([]*Topic, 0)
   err := x.Where("Author=?",name).Find(&topics)
   if err != nil {
      return nil, err
   }
   return topics, err
}

func GetAuthorById(id string) (string,error) {
  tid, err := strconv.ParseInt(id,10,64)
  if err != nil {
     return "", err
  }
  topic := &Topic{Id: tid}
  has, err := x.Get(topic)
  if has {
     if err != nil {
        return "", err
     }
     return topic.Author, err
  }
  return "", err
}

//reply
func AddReply(tid, nickname, content string) error {
     tidNum, err := strconv.ParseInt(tid,10,64)
     if err != nil {
        return err
     }

     reply := &Reply{
           Tid:      tidNum,
           Name:     nickname,
           Content:  content,
           Created:  time.Now(),
     }
     _, err = x.Insert(reply)
     if err != nil {
        return err
     }

     topic := &Topic{Id: tidNum}
     has, err := x.Get(topic)
     if has {
        topic.ReplyTime = time.Now()
        topic.ReplyCount++
        _, _ = x.Id(topic.Id).Update(topic)
     }
     return err
}

func GetAllReplies(tid string) ([]*Reply, error) {
     tidNum, err := strconv.ParseInt(tid,10,64)
     if err != nil {
        return nil, err
     }

     replies := make([]*Reply, 0)
     err = x.Where("Tid=?",tidNum).Find(&replies)
     return replies, err
}

func DeleteReply(rid string) error {
     ridNum, err := strconv.ParseInt(rid,10,64)
     if err != nil {
        return err
     }

     var tid  int64
     reply := &Reply{Id: ridNum}
     has, err := x.Get(reply)
     if has {
        tid = reply.Tid
        x.Delete(reply)
        if err != nil {
           return err
        }
     }

    replies := make([]*Reply, 0)
    err = x.Where("Tid=?",tid).Find(&replies)
    if err != nil {
       return nil
    }

    topic := &Topic{Id: tid}
    has, err = x.Get(topic)
    if has {
       topic.ReplyCount = int64(len(replies))
       if err != nil {
          return err
       }
       _ , err = x.Id(tid).Update(topic)
    }
    return err
}
