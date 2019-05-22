package controllers

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"

    "student-info-collect/models"
)

type BaseController struct {
    beego.Controller
    MyOrmer orm.Ormer
}

func (base *BaseController) Prepare() {
    base.MyOrmer = orm.NewOrm()
}

func (base *BaseController) ServeResult(status int, msg string) {
    res := &models.ServeResult{status, msg, nil}
    base.Data["json"] = res
    base.ServeJSON()
}

func (base *BaseController) ServeDataResult(status int, msg string, data interface{}) {
    res := &models.ServeResult{status, msg, data}
    base.Data["json"] = res
    base.ServeJSON()
}
