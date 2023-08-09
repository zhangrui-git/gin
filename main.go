package main

import (
	"by/user/bootstrap"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//defer profile.Start(profile.CPUProfile, profile.ProfilePath("./profile")).Stop()
	//defer profile.Start(profile.MemProfile, profile.ProfilePath("./profile")).Stop()
	//defer profile.Start(profile.GoroutineProfile, profile.ProfilePath("./profile")).Stop()

	e := gin.Default()
	bootstrap.Boot(e)
	log.Fatal(e.Run(":81"))
}
