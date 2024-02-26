package db

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var RabbitMQConnection = "amqp://" + RMQUSER + ":" + RMQPASS + "@" + RMQHOST + ":" + RMQPORT + "/"
var RedisConnection = REDISHOST + ":" + REDISPORT

func Mongo() *mongo.Client {
	credential := options.Credential{
		Username:   MONGOUSER,
		Password:   MONGOPASS,
		AuthSource: MONGOAUTHDB,
	}
	clientOptions := options.Client().ApplyURI("mongodb://" + MONGOHOST + ":" + MONGOPORT).SetAuth(credential)
	mongoClient, err := mongo.Connect(CTX, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return mongoClient
}

func MySQLConnect() *sql.DB {

	mySQLdb, err := sql.Open("mysql", MYSQLUSER+":"+MYSQLPASS+"@tcp("+MYSQLHOST+":"+MYSQLPORT+")/"+MYSQLDB+"?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	mySQLdb.SetMaxOpenConns(20)
	mySQLdb.SetMaxIdleConns(10)

	return mySQLdb
}

var MySqlDB = MySQLConnect()
var MongoClient = Mongo()
var CTX = context.TODO()
