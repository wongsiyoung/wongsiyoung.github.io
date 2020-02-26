package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/goconfig"
)

type ServiceConfig struct {
	TimeServerType string
	TimeServerIP   string
	TimeServerPort string
	TimeZone       string
	RCSServerIP    string
	RCSServerPort  string
}

type TimeZone struct {
	TimeServerName    string
	TimeServerAddress string
	TimeServerPort    string
	TimeZone          string
}

func SaveServiceConfig(c *gin.Context) {

	var info ServiceConfig
	err := c.Bind(&info)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{
			"status_code": 500,
			"errMsg":      "server error",
		})
		return
	}

	cfg, err := goconfig.LoadConfigFile("./conf/config.ini")
	if err != nil {
		log.Fatalf("cannot load config file: %s", err)
		c.JSON(500, gin.H{
			"status_code": 500,
			"errMsg":      "server error",
		})
		return
	}
	cfg.SetValue("ServiceConfig", "TimeServerType", info.TimeServerType)
	cfg.SetValue("ServiceConfig", "TimeServerIP", info.TimeServerIP)
	cfg.SetValue("ServiceConfig", "TimeServerPort", info.TimeServerPort)
	cfg.SetValue("ServiceConfig", "TimeZone", info.TimeZone)
	cfg.SetValue("ServiceConfig", "RCSServerIP", info.RCSServerIP)
	cfg.SetValue("ServiceConfig", "RCSServerPort", info.RCSServerPort)
	err = goconfig.SaveConfigFile(cfg, "./conf/config.ini")
	fmt.Println("save success")

	ifrefresh := c.MustGet("ifrefresh").(bool)
	token := ""
	code := 200
	if ifrefresh {
		token = c.MustGet("token").(string)
		code = 201
	}
	c.JSON(code, gin.H{
		"status_code": code,
		"errMsg":      "success",
		"token":       token,
	})
}

func ShowServiceConfig(c *gin.Context) {

	cfg, err := goconfig.LoadConfigFile("./conf/config.ini")
	if err != nil {
		log.Fatalf("cannot load config file: %s", err)
		c.JSON(500, gin.H{
			"status_code": 500,
			"errMsg":      "server error",
		})
		return
	}

	ifrefresh := c.MustGet("ifrefresh").(bool)
	token := ""
	code := 200
	if ifrefresh {
		token = c.MustGet("token").(string)
		code = 201
	}
	c.JSON(code, gin.H{
		"status_code":    code,
		"errMsg":         "Success",
		"token":          token,
		"TimeServerType": cfg.MustValue("ServiceConfig", "TimeServerType"),
		"TimeServerIP":   cfg.MustValue("ServiceConfig", "TimeServerIP"),
		"TimeServerPort": cfg.MustValue("ServiceConfig", "TimeServerPort"),
		"TimeZone":       cfg.MustValue("ServiceConfig", "TimeZone"),
		"RCSServerIP":    cfg.MustValue("ServiceConfig", "RCSServerIP"),
		"RCSServerPort":  cfg.MustValue("ServiceConfig", "RCSServerPort"),
	})
}

func SaveTimeZone(c *gin.Context) {

	var info TimeZone
	err := c.Bind(&info)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{
			"status_code": 500,
			"errMsg":      "server error",
		})
		return
	}

	cfg, err := goconfig.LoadConfigFile("./conf/config.ini")
	if err != nil {
		log.Fatalf("cannot load config file: %s", err)
		c.JSON(500, gin.H{
			"status_code": 500,
			"errMsg":      "server error",
		})
		return
	}
	cfg.SetValue("ServiceConfig", "TimeServerType", info.TimeServerName)
	cfg.SetValue("ServiceConfig", "TimeServerIP", info.TimeServerAddress)
	cfg.SetValue("ServiceConfig", "TimeServerPort", info.TimeServerPort)
	cfg.SetValue("ServiceConfig", "TimeZone", info.TimeZone)
	err = goconfig.SaveConfigFile(cfg, "./conf/config.ini")
	fmt.Println("save success")

	ifrefresh := c.MustGet("ifrefresh").(bool)
	token := ""
	code := 200
	if ifrefresh {
		token = c.MustGet("token").(string)
		code = 201
	}
	c.JSON(code, gin.H{
		"status_code": code,
		"errMsg":      "success",
		"token":       token,
	})
}
