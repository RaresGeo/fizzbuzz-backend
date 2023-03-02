package fizzbuzz

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterFizzBuzzEndpoint(router *gin.RouterGroup) {
	// endpoint that tikes a number and returns fizzbuzz
	// If the number is missing, return status code 400 (bad request)
	router.GET("/:number", func(c *gin.Context) {
		number := c.Param("number")
		input, err := strconv.ParseInt(number, 10, 32)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"result": fizzbuzz(input)})
	})
}

func fizzbuzz(input int64) string {
	if input%15 == 0 {
		return os.Getenv("DIVISIBLE_BY_3_AND_5")
	}
	if input%3 == 0 {
		return os.Getenv("DIVISIBLE_BY_3")
	}
	if input%5 == 0 {
		return os.Getenv("DIVISIBLE_BY_5")
	}
	return os.Getenv("DIVISIBLE_BY_NONE")
}
