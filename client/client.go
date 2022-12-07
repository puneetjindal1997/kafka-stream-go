/*
 *	Author: Puneet
 *	Use to start the client
 */
package main

import (
	"disastermanagement/pb"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(*gin.Context)
}
type routes struct {
	router *gin.Engine
}
type Routes []Route

var client pb.UserErrorServiceClient
var kafkaWriter *kafka.Writer
var kafkaReader *kafka.Reader

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	// Connection to internal grpc server
	conn, err := grpc.Dial(os.Getenv("SERVER_HOST")+":"+os.Getenv("SERVER_PORT"), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client = pb.NewUserErrorServiceClient(conn)
	// comment this producer initialiZer if it is not need full
	go kafkaWritterConnection()

	// route calling
	go Produce()
	go SaveLogs()
	ClientRoutes()
}

/*
 *	Function for grouping log routes
 */
func (r routes) LogGrouping(rg *gin.RouterGroup) {
	orderRouteGrouping := rg.Group("/log")
	for _, route := range GetDisasterLogs {
		switch route.Method {
		case "GET":
			orderRouteGrouping.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			orderRouteGrouping.POST(route.Pattern, route.HandlerFunc)
		case "OPTIONS":
			orderRouteGrouping.OPTIONS(route.Pattern, route.HandlerFunc)
		case "PUT":
			orderRouteGrouping.PUT(route.Pattern, route.HandlerFunc)
		case "DELETE":
			orderRouteGrouping.DELETE(route.Pattern, route.HandlerFunc)
		default:
			orderRouteGrouping.GET(route.Pattern, func(c *gin.Context) {
				c.JSON(200, gin.H{
					"result": "Specify a valid http method with this route.",
				})
			})
		}
	}
}

// append routes with versions
func ClientRoutes() {
	r := routes{
		router: gin.Default(),
	}
	v1 := r.router.Group(os.Getenv("API_VERSION"))
	r.LogGrouping(v1)
	if err := r.router.Run(":" + os.Getenv("CLIENT_PORT")); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

// intialize the writer with the broker addresses, and the topic
func kafkaWritterConnection() {
	kafkaWriter = kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{os.Getenv("KAFKA_HOST") + ":" + os.Getenv("KAFKA_PORT")},
		Topic:   "testError",
		// wait for at most 3 seconds before receiving new data
		WriteTimeout: 3 * time.Second,
	})
}

// initialize a new reader with the brokers and topic
func kafkaReaderConnection(topicName string) {
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	kafkaReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{os.Getenv("KAFKA_HOST") + ":" + os.Getenv("KAFKA_PORT")},
		Topic:   topicName,
		GroupID: "my-group",
	})
}
