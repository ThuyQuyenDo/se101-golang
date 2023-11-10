package app

import (
	"flag"
	"fmt"
	"go-ecommerce/config"
	"go-ecommerce/db"
	usrapis "go-ecommerce/internal/user/apis"
	"go-ecommerce/middleware"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Run(flagArg *FlagArg) {
	flag.Parse()

	if err := os.Setenv("CONFIG_PATH", *flagArg.ConfigPath); err != nil {
		panic(err)
	}

	if err := os.Setenv("CONFIG_ENV", *flagArg.Env); err != nil {
		panic(err)
	}

	if err := os.Setenv("CONFIG_ALLLOW_MIGRATION",
		strconv.FormatBool(*flagArg.Upgrade)); err != nil {
		panic(err)
	}

	var appConfig config.AppConfig

	envValue := os.Getenv("CONFIG_ALLLOW_MIGRATION")
	allowUpgrade, err := strconv.ParseBool(envValue)
	if err != nil {
		panic(err)
	}

	config.InitLoadAppConf(&appConfig)

	dbInstance := db.InitDatabase(allowUpgrade, appConfig)

	appRouter := gin.Default()
	appRouter.Use(middleware.Recover()) // panic must have recover to show error

	usrapis.SetupRouter(appRouter, dbInstance)

	//appRouter.Run(":3000") -> available :
	if err := appRouter.Run(fmt.Sprintf(":%d", appConfig.ServicePort)); err != nil {
		panic(err)
	}

	//fmt.Print(appConfig)
}
