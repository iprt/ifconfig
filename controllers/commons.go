package controllers

import (
	"errors"
	"github.com/beego/beego/v2/server/web"
	"github.com/zu1k/nali/pkg/geoip"
	"github.com/zu1k/nali/pkg/ip2region"
	"github.com/zu1k/nali/pkg/ipip"
	"github.com/zu1k/nali/pkg/qqwry"
)

const (
	QQWryPath = "qqwry.dat"
	// ZXIPv6WryPath    = "zxipv6wry.db"
	GeoLite2CityPath = "GeoLite2-City.mmdb"
	IPIPFreePath     = "ipipfree.ipdb"
	Ip2RegionPath    = "ip2region.db"
)

var (
	geoip2Instance    *geoip.GeoIP
	qqwryInstance     *qqwry.QQwry
	ipipInstance      *ipip.IPIPFree
	ip2regionInstance *ip2region.Ip2Region
)

func init() {
	geoip2Instance, _ = geoip.NewGeoIP(GeoLite2CityPath)
	qqwryInstance, _ = qqwry.NewQQwry(QQWryPath)
	ipipInstance, _ = ipip.NewIPIP(IPIPFreePath)
	ip2regionInstance, _ = ip2region.NewIp2Region(Ip2RegionPath)
}

type MainController struct {
	web.Controller
}

func (mainController *MainController) QueryGeoip2(ip string) (string, error) {
	if geoip2Instance == nil {
		return "", errors.New("Geoip2 service not available")
	}
	res, err := geoip2Instance.Find(ip)
	if err != nil {
		return "", err
	}
	return res.String(), nil
}

func (mainController *MainController) GetLocationFromGeoIP2() {
	ip := mainController.GetString("ip", mainController.Ctx.Input.IP())
	res, err := mainController.QueryGeoip2(ip)
	if err != nil {
		mainController.Data["Value"] = err.Error()
	} else {
		mainController.Data["Value"] = res
	}
	mainController.TplName = "value.tpl"
}

func (mainController *MainController) QueryQQWry(ip string) (string, error) {
	if qqwryInstance == nil {
		return "", errors.New("QQWry service not available")
	}
	res, err := qqwryInstance.Find(ip)
	if err != nil {
		return "", err
	}
	return res.String(), nil
}

func (mainController *MainController) GetLocationFromQQWry() {
	ip := mainController.GetString("ip", mainController.Ctx.Input.IP())
	res, err := mainController.QueryQQWry(ip)
	if err != nil {
		mainController.Data["Value"] = err.Error()
	} else {
		mainController.Data["Value"] = res
	}
	mainController.TplName = "value.tpl"
}

func (mainController *MainController) QueryIPIPFree(ip string) (string, error) {
	if ipipInstance == nil {
		return "", errors.New("IPIP free service not available")
	}
	res, err := ipipInstance.Find(ip)
	if err != nil {
		return "", err
	}
	return res.String(), nil
}

func (mainController *MainController) GetLocationFromIPIP() {
	ip := mainController.GetString("ip", mainController.Ctx.Input.IP())
	res, err := mainController.QueryIPIPFree(ip)
	if err != nil {
		mainController.Data["Value"] = err.Error()
	} else {
		mainController.Data["Value"] = res
	}
	mainController.TplName = "value.tpl"
}

func (mainController *MainController) QueryIP2Region(ip string) (string, error) {
	if ip2regionInstance == nil {
		return "", errors.New("IP2Region service not available")
	}
	res, err := ip2regionInstance.Find(ip)
	if err != nil {
		return "", err
	}
	return res.String(), nil
}

type ifconfig struct {
	Email      string
	QQWry      string
	Geoip2     string
	IPIP       string
	IP2Region  string
	UserAgent  string
	Host       string
	IP         string
	Port       string
	Method     string
	Encoding   string
	Mime       string
	Connection string
	Via        string
	Charset    string
	Keepalive  string
	Forwarded  string
	Lang       string
	Referer    string
}
