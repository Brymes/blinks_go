package main

import (
	"StickyLabsBlinks/app"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	//TODO action.json
	router.GET("/actions.json", app.ActionsRulesHAndler)
	router.GET("/api/actions/mint_nft", app.GetActionsHandler)
	router.OPTIONS("/api/actions/mint_nft", app.OptionsHandler)
	router.POST("/api/actions/mint_nft", app.PostHandler)

	log.Println("StickyLabs Blink Active 🚀")
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
		return
	}
}
