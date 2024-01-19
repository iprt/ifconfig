package controllers

import (
	"bytes"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func (mainController *MainController) Get() {
	if noweb := os.Getenv("NOWEB"); noweb == "1" {
		mainController.Abort("404")
		return
	}
	ip := mainController.GetString("ip", mainController.Ctx.Input.IP())
	mainController.Data["IP"] = ip
	if strings.Contains(mainController.Ctx.Request.UserAgent(), "curl") {
		mainController.TplName = "iponly.tpl"
	} else {
		mainController.fillIndexData()
		mainController.TplName = "index.tpl"
	}
}

func (mainController *MainController) GetGeo() {
	if noweb := os.Getenv("NOWEB"); noweb == "1" {
		mainController.Abort("404")
		return
	}
	mainController.fillGeoData()

	if strings.Contains(mainController.Ctx.Request.UserAgent(), "curl") {
		mainController.TplName = "geo.tpl"
	} else {
		mainController.fillIndexData()
		mainController.TplName = "index.tpl"
	}
}

func (mainController *MainController) fillGeoData() {
	ip := mainController.GetString("ip", mainController.Ctx.Input.IP())

	mainController.Data["IP"] = ip
	mainController.Data["Geoip2"], _ = mainController.QueryGeoip2(ip)
	mainController.Data["IPIP"], _ = mainController.QueryIPIPFree(ip)
	mainController.Data["QQWry"], _ = mainController.QueryQQWry(ip)
	mainController.Data["IP2Region"], _ = mainController.QueryIP2Region(ip)
}

// fillIndexData populates the data for the index page.
//
// It retrieves the IP address from the request, and if no "ip" parameter is provided,
// it uses the client's IP address. Then, it sets the "BaseUrl" value to the hostname of the request.
//
// If the "record.enable" configuration is set to true, it retrieves the default record or the record
// associated with the current host, and sets it in the "Record" field of the data. Otherwise, it sets
// the "Record" field to ".".
//
// The "Email" field is set to the configured email address.
// The "UserAgent" field is set to the user agent of the request.
// The "CopyrightBegin" field is set to "2012".
// The "CopyrightEnd" field is set to the current year.
//
// It performs a reverse DNS lookup with the retrieved IP address, and sets the "Host" field with the
// result if the lookup is successful.
//
// It queries various IP geolocation providers (Geoip2, IPIPFree, QQWry, IP2Region) with the IP address,
// and sets the respective fields in the data with the results.
//
// It retrieves the remote address and extracts the port number, setting it in the "Port" field.
//
// It sets the "Method" field with the request method.
// It sets the "Encoding" field with the first value of the "Accept-Encoding" header, if present.
// It sets the "Mime" field with the first value of the "Accept" header, if present.
// It sets the "Connection" field with the first value of the "Connection" header, if present.
// It sets the "Via" field with the first value of the "Via" header, if present.
// It sets the "Charset" field with the first value of the "Charset" header, if present.
// It sets the "Keepalive" field with the first value of the "KeepAlive" header, if present.
// It sets the "Forwarded" field with the first value of the "X-Forwarded-For" header, if present.
// It sets the "Lang" field with the first value of the "Accept-Language" header, if present.
//
// Finally, it sets the "Referer" field with the referrer URL from the request.
func (mainController *MainController) fillIndexData() {
	ip := mainController.GetString("ip", mainController.Ctx.Input.IP())
	//mainController.Data["BaseUrl"] = web.AppConfig.DefaultString("baseurl", "ipcrystal.com")
	host := mainController.Ctx.Request.Host

	// 配置 BaseUrl
	mainController.Data["BaseUrl"] = host

	recordEnable := web.AppConfig.DefaultBool("record.enable", false)
	if recordEnable {
		// 不同域名下的备案不同的映射
		defaultRecord := web.AppConfig.DefaultString("record", "")
		recordWithHost := web.AppConfig.DefaultString("record."+host, "")
		if recordWithHost == "" || recordWithHost == "." {
			mainController.Data["Record"] = defaultRecord
		} else {
			mainController.Data["Record"] = recordWithHost
		}
	} else {
		mainController.Data["Record"] = "."
	}

	// 配置 Email
	mainController.Data["Email"] = web.AppConfig.DefaultString("email", "")

	// 配置 UserAgent
	mainController.Data["UserAgent"] = mainController.Ctx.Request.UserAgent()

	// 配置 CopyrightBegin
	mainController.Data["CopyrightBegin"] = "2012"

	// 配置 CopyrightEnd
	mainController.Data["CopyrightEnd"] = strconv.Itoa(time.Now().Year())

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
	mainController.Data["Geoip2"], _ = mainController.QueryGeoip2(ip)
	mainController.Data["IPIP"], _ = mainController.QueryIPIPFree(ip)
	mainController.Data["QQWry"], _ = mainController.QueryQQWry(ip)
	mainController.Data["IP2Region"], _ = mainController.QueryIP2Region(ip)
	remoteAddr := []byte(mainController.Ctx.Request.RemoteAddr)
	pos := bytes.IndexByte(remoteAddr, ':')
	mainController.Data["Port"] = string(remoteAddr[pos+1:])
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
}
