package main

import (
	"flag"
	"log"
	"personal/fizzbuzz/backend/src/fizzbuzz"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Equivalent of node_env
	production := flag.Bool("production", false, "production mode")
	stage := flag.Bool("stage", false, "stage mode")

	flag.Parse()

	var err error

	switch true {
	case *production:
		err = godotenv.Load(".env.production")
		log.Println("Production mode")
	case *stage:
		err = godotenv.Load(".env.stage")
		log.Println("stage mode")
	default:
		err = godotenv.Load()
		log.Println("Development mode")
	}

	// Realistically this will only use .env, seeing as to how there are only three environment variables
	// And they will most likely be the same across the different environments
	// But assuming I was asked to use environment variables to show that I can, here you go :)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create a new Gin router
	router := gin.Default()

	// set trusted proxies to nil
	router.SetTrustedProxies(nil)

	// Register the fizzbuzz router
	fizzbuzzRouter := router.Group("/fizzbuzz")
	fizzbuzz.RegisterFizzBuzzEndpoint(fizzbuzzRouter)

	// Run the server on port 5001
	router.Run(":5001")
}
