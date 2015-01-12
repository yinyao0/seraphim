
package controllers

import (
    "github.com/astaxie/beego"
    //"seraphim/models"
)

type TopicController struct {
     beego.Controller
}

func (this *TopicController) Get() {
    this.TplNames = "topic.html"
}

func (this *TopicController) Add() {
    this.TplNames = "topic_add.html"
}
