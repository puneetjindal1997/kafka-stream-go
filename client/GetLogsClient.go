package main

import (
	"context"
	"disastermanagement/constants"
	"disastermanagement/pb"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

/*
 *	Function to get the logs
 *
 *	return c.JSON response
 */
func GetLogs(c *gin.Context) {
	finalLog := []interface{}{}
	response, serverErr := client.GetUserLogs(c, &pb.Null{})
	if serverErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": constants.InternalServerError,
			"data":    nil,
		})
		return
	}
	// unmarshalling single single log
	for _, singleLog := range response.UserResp {
		// binding
		finalLog = append(finalLog, map[string]interface{}{"id": singleLog.Id, "data": singleLog.Message, "updated_at": singleLog.UpdatedAt})
	}
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": constants.Success,
		"data": map[string]interface{}{
			"data":     finalLog,
			"log_type": "user",
		},
	})
}

func Produce() {
	log.Println("In producer")
	// initialize a counter
	i := 0

	for {
		// each kafka message has a key and value. The key is used
		// to decide which partition (and consequently, which broker)
		// the message gets published on
		err := kafkaWriter.WriteMessages(context.Background(), kafka.Message{
			Key: []byte(strconv.Itoa(i)),
			// create an arbitrary message payload for the value
			Value: []byte("{message:'this is message" + strconv.Itoa(i) + "'}"),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

		// log a confirmation once the message is written
		fmt.Println("writes:", i)
		i++
		// sleep for a second
		time.Sleep(time.Second)
		if i == 10 {
			break
		}
	}
}

func SaveLogs() {
	log.Println("In consumer")
	// consumer initializer
	kafkaReaderConnection("testError")
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, kafkaErr := kafkaReader.ReadMessage(context.Background())
		if kafkaErr != nil {
			log.Println("There is something wrong while reading", kafkaErr)
		}
		response, serverErr := client.SaveUserLogs(context.Background(),
			&pb.UserErrorRequest{
				Type:      constants.LogType(1).String(),
				Message:   string(msg.Value),
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			})
		if serverErr != nil {
			log.Println("There is something wrong in the server", serverErr)
		}
		if response.Status {
			log.Println("Logs are saved successfully. :D")
		}
	}
}
