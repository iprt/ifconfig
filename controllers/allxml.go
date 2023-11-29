package controllers

import (
	"bytes"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"log"
	"net"
)

func (mainController *MainController) GetAllXML() {
	thisData := ifconfig{}
	thisData.Email = web.AppConfig.DefaultString("email", "")
	thisData.UserAgent = mainController.Ctx.Request.UserAgent()

	ip := mainController.GetString("ip", mainController.Ctx.Input.IP())
	names, err := net.LookupAddr(ip)
	if err != nil || len(names) == 0 {
		thisData.Host = ""
	} else {
		var value string
		for _, v := range names {
			value += fmt.Sprintf("%s\n", v)
		}
		thisData.Host = value
	}

	thisData.IP = ip
	thisData.Geoip2, _ = mainController.QueryGeoip2(ip)
	thisData.IPIP, _ = mainController.QueryIPIPFree(ip)
	thisData.QQWry, _ = mainController.QueryQQWry(ip)
	thisData.IP2Region, _ = mainController.QueryIP2Region(ip)
	log.Println(thisData.IP2Region)
	remoteAddr := []byte(mainController.Ctx.Request.RemoteAddr)
	pos := bytes.IndexByte(remoteAddr, ':')
	thisData.Port = string(remoteAddr[pos+1:])
	thisData.Method = mainController.Ctx.Request.Method
	if len(mainController.Ctx.Request.Header["Accept-Encoding"]) > 0 {
		thisData.Encoding = mainController.Ctx.Request.Header["Accept-Encoding"][0]
	}
	if len(mainController.Ctx.Request.Header["Accept"]) > 0 {
		thisData.Mime = mainController.Ctx.Request.Header["Accept"][0]
	}
	if len(mainController.Ctx.Request.Header["Connection"]) > 0 {
		thisData.Connection = mainController.Ctx.Request.Header["Connection"][0]
	}
	if len(mainController.Ctx.Request.Header["Via"]) > 0 {
		thisData.Via = mainController.Ctx.Request.Header["Via"][0]
	}
	if len(mainController.Ctx.Request.Header["Charset"]) > 0 {
		thisData.Charset = mainController.Ctx.Request.Header["Charset"][0]
	}
	if len(mainController.Ctx.Request.Header["KeepAlive"]) > 0 {
		thisData.Keepalive = mainController.Ctx.Request.Header["KeepAlive"][0]
	}
	if len(mainController.Ctx.Request.Header["X-Forwarded-For"]) > 0 {
		thisData.Forwarded = mainController.Ctx.Request.Header["X-Forwarded-For"][0]
	}
	if len(mainController.Ctx.Request.Header["Accept-Language"]) > 0 {
		thisData.Lang = mainController.Ctx.Request.Header["Accept-Language"][0]
	}
	thisData.Referer = mainController.Ctx.Input.Refer()

	mainController.Data["xml"] = thisData
	serverArr := mainController.ServeXML()
	if serverArr != nil {
		return
	}
}
