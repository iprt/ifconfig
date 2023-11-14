package controllers

import (
	"bytes"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"net"
)

func (mainController *MainController) GetAll() {
	mainController.Data["Email"] = web.AppConfig.DefaultString("email", "")
	mainController.Data["UserAgent"] = mainController.Ctx.Request.UserAgent()
	ip := mainController.GetString("ip", mainController.Ctx.Input.IP())
	names, err := net.LookupAddr(ip)
	if err != nil || len(names) == 0 {
		mainController.Data["Host"] = ""
	} else {
		var value string
		for _, v := range names {
			value += fmt.Sprintf("%s\n", v)
		}
		mainController.Data["Host"] = value
	}
	mainController.Data["IP"] = ip
	mainController.Data["Geoip2"], _ = mainController.QueryGeoip2(ip)
	mainController.Data["IPIP"], _ = mainController.QueryIPIPFree(ip)
	mainController.Data["QQWry"], _ = mainController.QueryQQWry(ip)
	mainController.Data["IP2Region"], _ = mainController.QueryIP2Region(ip)
	remote_addr := []byte(mainController.Ctx.Request.RemoteAddr)
	pos := bytes.IndexByte(remote_addr, ':')
	mainController.Data["Port"] = string(remote_addr[pos+1:])
	mainController.Data["Method"] = mainController.Ctx.Request.Method
	if len(mainController.Ctx.Request.Header["Accept-Encoding"]) > 0 {
		mainController.Data["Encoding"] = mainController.Ctx.Request.Header["Accept-Encoding"][0]
	}
	if len(mainController.Ctx.Request.Header["Accept"]) > 0 {
		mainController.Data["Mime"] = mainController.Ctx.Request.Header["Accept"][0]
	}
	if len(mainController.Ctx.Request.Header["Connection"]) > 0 {
		mainController.Data["Connection"] = mainController.Ctx.Request.Header["Connection"][0]
	}
	if len(mainController.Ctx.Request.Header["Via"]) > 0 {
		mainController.Data["Via"] = mainController.Ctx.Request.Header["Via"][0]
	}
	if len(mainController.Ctx.Request.Header["Charset"]) > 0 {
		mainController.Data["Charset"] = mainController.Ctx.Request.Header["Charset"][0]
	}
	if len(mainController.Ctx.Request.Header["KeepAlive"]) > 0 {
		mainController.Data["Keepalive"] = mainController.Ctx.Request.Header["KeepAlive"][0]
	}
	if len(mainController.Ctx.Request.Header["X-Forwarded-For"]) > 0 {
		mainController.Data["Forwarded"] = mainController.Ctx.Request.Header["X-Forwarded-For"][0]
	}
	if len(mainController.Ctx.Request.Header["Accept-Language"]) > 0 {
		mainController.Data["Lang"] = mainController.Ctx.Request.Header["Accept-Language"][0]
	}
	mainController.Data["Referer"] = mainController.Ctx.Input.Refer()

	mainController.TplName = "all.tpl"
}
