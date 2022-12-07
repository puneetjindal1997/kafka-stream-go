package database

import (
	"disastermanagement/pb"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type manager struct {
	connection *gorm.DB
}

var TxnDb *gorm.DB

var Mgr Manager

type Manager interface {
	InsertUserLog(*pb.UserErrorRequest) error
	GetUserLogs(string) ([]*pb.UserErrorRequest, error)
}

/*
 *	init funcion to create connection with database while running the server
 */
func DbInit() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	dsn := os.Getenv("SQLUSERNAME") + ":" + os.Getenv("SQLPASSWORD") + "@tcp(" + os.Getenv("DBHOST") + ":" + os.Getenv("DBPORT") + ")/" + os.Getenv("SQLDATABASE")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("not connected")
		return
	}
	log.Println("db connected")
	TxnDb = db
	Mgr = &manager{connection: db}
}
