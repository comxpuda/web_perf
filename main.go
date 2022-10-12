package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"runtime"
)

func main() {
	ConfigRuntime()
	startStats()
	StartGin()
}

func ConfigRuntime() {
	nCpu := runtime.NumCPU()
	runtime.GOMAXPROCS(nCpu)
	fmt.Printf("Running with %d CPUs \n", nCpu)
}

func startStats() {
	go statsWorker()
}

func StartGin() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/test", testGet)
	r.GET("/insert", insertDoctorDB)
	r.GET("/query", queryDoctorDB)

	r.GET("/stats", getStats)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}

}
