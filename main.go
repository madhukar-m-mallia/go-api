package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madhukar-m-mallia/go-api/controller"
	"github.com/madhukar-m-mallia/go-api/service"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	keyValueService    service.KeyValueService       = service.New()
	keyValueController controller.KeyValueController = controller.New(keyValueService)
)

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func handleRequests() {
	router := gin.Default()
	router.GET("/", healthCheckAPI)
	router.GET("/get/:key", getValueOfKeyAPI)
	router.POST("/set", setKeyValueAPI)
	router.GET("/search", searchKeyValueAPI)
	router.GET("/metrics", prometheusHandler())
	log.Fatal(router.Run())
}

func main() {
	handleRequests()
}

func healthCheckAPI(c *gin.Context) {
	log.Println("Endpoint Hit: healthCheckAPI")
	c.JSON(http.StatusOK, gin.H{
		"message": "Healthy",
	})
}

func getValueOfKeyAPI(c *gin.Context) {
	key := c.Param("key")
	value, err := keyValueController.FindOne(key)

	if err != nil {
		c.String(http.StatusNotFound, "Could not get value for key '%s'. Error: %s", key, err.Error())
	} else {
		c.String(http.StatusOK, "%s", value)
	}
}

func setKeyValueAPI(c *gin.Context) {
	appCache, err := keyValueController.Set(c)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		c.String(http.StatusOK, "Updates %s: %s", appCache.Key, appCache.Value)
	}
}

func searchKeyValueAPI(c *gin.Context) {
	searchResult, err := keyValueController.Search(c)
	if err != nil {
		c.String(http.StatusNotFound, "Could not search value. Error: %s", err.Error())
	} else {
		c.String(http.StatusOK, "%s", searchResult)
	}
}
