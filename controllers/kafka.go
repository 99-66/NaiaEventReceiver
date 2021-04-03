package controllers

import (
	"encoding/json"
	"errors"
	"github.com/99-66/NaiaArticleEventReceiver/models"
	"github.com/Shopify/sarama"
	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"net/http"
)

type KafkaClient struct {
	Producer sarama.AsyncProducer
	Topic string
}

type KafkaConfig struct {
	Broker string `env:"BROKER"`
	//Topic string `env:"TOPIC"`
}

func NewKafkaClient() (*KafkaClient, error) {
	// env에서 broker와 topic을 읽어서 클라이언트를 생성한다
	conf := KafkaConfig{}
	if err := env.Parse(&conf); err != nil {
		return nil, errors.New("cloud not load kafka configuration")
	}

	// 프로듀서를 생성한다
	saramaCfg := sarama.NewConfig()
	producer, err := sarama.NewAsyncProducer([]string{conf.Broker}, saramaCfg)
	if err != nil {
		return nil, err
	}

	kafka := KafkaClient{
		Producer: producer,
		//Topic: conf.Topic,
	}

	return &kafka, nil
}

// Create godoc
// @Summary POST Todo
// @Description Event Articles Send to Kafka Todo
// @Tags Kafka
// @Accept application/json
// @Produce application/json
// @Param Article body models.Article true "Create Article"
// @Success 200 {string} {}
// @Failure 404 {object} config.APIError
// @Router /evt/article [post]
func (k *KafkaClient) POST(c *gin.Context) {
	// JSON Body를 변환한다(required field check)
	var article models.Article
	err := c.BindJSON(&article)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	// Kafka로 전달하기 위해 마샬링한다
	valJson, err := json.Marshal(article)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed Json UnMarshaling"})
		return
	}

	// Kafka로 메시지를 전달한다
	msg := &sarama.ProducerMessage{
		Topic: article.Tag,
		Value: sarama.ByteEncoder(valJson),
	}
	k.Producer.Input() <- msg

	c.JSON(http.StatusOK, "")
}