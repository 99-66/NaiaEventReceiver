package main

import (
	"github.com/99-66/NaiaArticleEventReceiver/controllers"
	_ "github.com/99-66/NaiaArticleEventReceiver/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
)


// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /
// @query.collection.format multi

// @title Article Event Receiver API
// @description Article Event Receiver API
// @schemes http https

func main() {
	kafkaClient, err := controllers.NewKafkaClient()
	if err != nil {
		panic(err)
	}
	defer kafkaClient.Producer.Close()

	r := initRoutes(kafkaClient)
	log.Fatal(r.Run())
}

func initRoutes(kafkaClient *controllers.KafkaClient) *gin.Engine {
	r := gin.Default()

	// CORS allow all origins
	conf := cors.DefaultConfig()
	conf.AllowAllOrigins = true
	r.Use(cors.New(conf))

	// Kafka routes
	r.GET("/", index)
	r.POST("/event/recv", kafkaClient.POST)

	// Swagger Settings
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, "/")
}