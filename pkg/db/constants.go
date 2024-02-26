package db

import "os"

var MONGOHOST = os.Getenv("MONGO_ADMIN_HOST")
var MONGOPORT = os.Getenv("MONGO_ADMIN_PORT")
var MONGOUSER = os.Getenv("MONGO_ADMIN_USER")
var MONGOPASS = os.Getenv("MONGO_ADMIN_PASS")
var MONGOAUTHDB = os.Getenv("MONGO_ADMIN_AUTH_DB")

var MYSQLHOST = os.Getenv("MYSQL_ADMIN_HOST")
var MYSQLPORT = os.Getenv("MYSQL_ADMIN_PORT")
var MYSQLUSER = os.Getenv("MYSQL_ADMIN_USER")
var MYSQLPASS = os.Getenv("MYSQL_ADMIN_PASS")
var MYSQLDB = os.Getenv("MYSQL_ADMIN_AUTHDB")

var RMQHOST = os.Getenv("RMQ_HOST")
var RMQPORT = os.Getenv("RMQ_PORT")
var RMQUSER = os.Getenv("RMQ_USER")
var RMQPASS = os.Getenv("RMQ_PASS")

var REDISHOST = os.Getenv("REDIS_HOST")
var REDISPORT = os.Getenv("REDIS_PORT")
