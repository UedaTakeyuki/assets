package main

import (
	//	"log"

	//	"github.com/gin-gonic/autotls"
	//"fmt"
	//	"fmt"
	//	"log"
	"flag"

	"github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/pelletier/go-toml"
)

//const DataFileJson = "idTable.json"
//const DataFileCSV = "idTable.csv"
type Mode struct {
	// SQLite recovery mode: Fetch from faunaDB in case SQLite NOT return data
	isSQLiteRecovery bool // SQLite recovery mode:
}

type Grobal struct {
	mode Mode
}

var global = Grobal{mode: Mode{isSQLiteRecovery: false}}

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {

	// idTable, idCount
	//	idCount, _ := NewPersistentInt("idTableCount")
	//	idTable, _ := NewIDTable("idTable", idCount)

	// CLI options
	// refer https://gobyexample.com/command-line-flags
	configfilePtr := flag.String("configfile", "assets.toml", "a string")

	// config
	// https://godoc.org/github.com/pelletier/go-toml#LoadFile
	config, _ := toml.LoadFile(*configfilePtr)
	//	config, _ := toml.LoadFile("ml.toml")
	runConfig := config.Get("Run").(*toml.Tree)

	// routes
	r := gin.Default()
	//r.Static("/assets", "./assets")
	r.Use(static.Serve("/", static.LocalFile("./assets", false)))
	//	r.LoadHTMLGlob("templates/*.html")
	r.Use(cors.Default())
	/*	r.Use(cors.New(cors.Config{
			AllowOrigins: []string{"http://localhost:8080", "https://ml.uedasoft.com"},
			//		AllowMethods:     []string{"PUT", "PATCH"},
			//		AllowHeaders:     []string{"Origin"},
			//		ExposeHeaders:    []string{"Content-Length"},
			//		AllowCredentials: true,
			//		AllowOriginFunc: func(origin string) bool {
			//			return origin == "https://github.com"
			//		},
			//		MaxAge: 12 * time.Hour,
		}))
	*/
	r.Use(helmet.NoCache())

	r.GET("/ping", func(c *gin.Context) {
		pong(c)
	})

	r.Run(runConfig.Get("port").(string))
}
