package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"zen/mapper"
	"zen/router"
)

func main() {

	setup()
	defer tearDown()
}

func setup() {
	//step 0, setup configurations here

	// step 1, setup data base here
	if err := mapper.Setup(); nil != err {
		println("cant connect data base, application will be quit, maybe cause: %s", err.Error())
		os.Exit(-1)
	}
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	router.Setup(engine)
	if err := engine.Run("localhost:8080"); nil != err {
		println("application listen at localhost:8080 start failed, maybe case: %s", err.Error())
		os.Exit(-1)
	}
}

func tearDown() {
	if err := mapper.TearDown(); nil != err {
		println("data base connections close failed, maybe cause: %s", err.Error())
	}
}
