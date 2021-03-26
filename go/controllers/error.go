package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

type ResultResponse struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (c *ErrorController) Error404() {
	m := ResultResponse{
		Success: false,
		Msg:     "请求的地址不存在",
		Data:    nil,
	}
	c.Data["json"] = m
	c.ServeJSON()
}

