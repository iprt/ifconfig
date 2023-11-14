package controllers

import (
	"bytes"
	"fmt"
	"net"
)

func (mainController *MainController) GetLocationFromIP2Region() {
	ip := mainController.GetString("ip", mainController.Ctx.Input.IP())
	res, err := mainController.QueryIP2Region(ip)
	if err != nil {
		mainController.Data["Value"] = err.Error()
	} else {
		mainController.Data["Value"] = res
	}
	mainController.TplName = "value.tpl"
}

func (mainController *MainController) GetForwarded() {
	if len(mainController.Ctx.Request.Header["X-Forwarded-For"]) > 0 {
		mainController.Data["Value"] = mainController.Ctx.Request.Header["X-Forwarded-For"][0]
	}
	mainController.TplName = "value.tpl"
}

func (mainController *MainController) GetHost() {
	ip := mainController.GetString("ip", mainController.Ctx.Input.IP())
	names, err := net.LookupAddr(ip)
	if err != nil || len(names) == 0 {
		mainController.Data["Value"] = ""
	} else {
		var value string
		for _, v := range names {
			value += fmt.Sprintf("%s\n", v)
		}
		mainController.Data["Value"] = value
	}
	mainController.TplName = "value.tpl"
}

func (mainController *MainController) GetIP() {
	ip := mainController.GetString("ip", mainController.Ctx.Input.IP())
	mainController.Data["Value"] = ip
	mainController.TplName = "value.tpl"
}

func (mainController *MainController) GetPort() {
	remote_addr := []byte(mainController.Ctx.Request.RemoteAddr)
	pos := bytes.IndexByte(remote_addr, ':')
	mainController.Data["Value"] = string(remote_addr[pos+1:])
	mainController.TplName = "value.tpl"
}

func (mainController *MainController) GetVia() {
	if len(mainController.Ctx.Request.Header["Via"]) > 0 {
		mainController.Data["Value"] = mainController.Ctx.Request.Header["Via"][0]
	}
	mainController.TplName = "value.tpl"
}

func (mainController *MainController) GetMime() {
	if len(mainController.Ctx.Request.Header["Accept"]) > 0 {
		mainController.Data["Value"] = mainController.Ctx.Request.Header["Accept"][0]
	}
	mainController.TplName = "value.tpl"
}

func (mainController *MainController) GetLang() {
	if len(mainController.Ctx.Request.Header["Accept-Language"]) > 0 {
		mainController.Data["Value"] = mainController.Ctx.Request.Header["Accept-Language"][0]
	}
	mainController.TplName = "value.tpl"
}

func (mainController *MainController) GetCharset() {
	if len(mainController.Ctx.Request.Header["Charset"]) > 0 {
		mainController.Data["Value"] = mainController.Ctx.Request.Header["Charset"][0]
	}
	mainController.TplName = "value.tpl"
}

func (mainController *MainController) GetEncoding() {
	if len(mainController.Ctx.Request.Header["Accept-Encoding"]) > 0 {
		mainController.Data["Value"] = mainController.Ctx.Request.Header["Accept-Encoding"][0]
	}
	mainController.TplName = "value.tpl"
}

func (mainController *MainController) GetUserAgent() {
	mainController.Data["Value"] = mainController.Ctx.Request.UserAgent()
	mainController.TplName = "value.tpl"
}

func (mainController *MainController) GetConnection() {
	if len(mainController.Ctx.Request.Header["Connection"]) > 0 {
		mainController.Data["Value"] = mainController.Ctx.Request.Header["Connection"][0]
	}
	mainController.TplName = "value.tpl"
}

func (mainController *MainController) GetKeepAlive() {
	if len(mainController.Ctx.Request.Header["KeepAlive"]) > 0 {
		mainController.Data["Value"] = mainController.Ctx.Request.Header["KeepAlive"][0]
	}
	mainController.TplName = "value.tpl"
}
