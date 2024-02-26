package activity

import (
	"encoding/json"
	"strconv"

	"github.com/Lanquill/Forge/pkg/db"
	"github.com/gomodule/redigo/redis"
	"github.com/nitishm/go-rejson"
)

type RedisDocIdError struct{}

func (m *RedisDocIdError) Error() string {
	return "Error"
}

func SetEntityInfoRedis(entityId int64, entityDetails EntityDetailsTemp) error {

	var conn, _ = redis.Dial("tcp", db.RedisConnection)
	defer conn.Close()

	rh := rejson.NewReJSONHandler()
	rh.SetRedigoClient(conn)

	// Store in Redis
	_, err := rh.JSONSet(strconv.Itoa(int(entityId)), ".", entityDetails)
	return err
}

func GetEntityInfoRedis(entityId int64) (EntityDetailsTemp, error) {

	var conn, _ = redis.Dial("tcp", db.RedisConnection)
	defer conn.Close()

	rh := rejson.NewReJSONHandler()
	rh.SetRedigoClient(conn)

	entityDetails := EntityDetailsTemp{}

	res, err := rh.JSONGet(strconv.Itoa(int(entityId)), ".")
	if err != nil {

		return entityDetails, err
	}

	if res == nil {

		return entityDetails, &RedisDocIdError{}
	}

	err = json.Unmarshal(res.([]byte), &entityDetails)
	if err != nil {
		return entityDetails, err
	}

	return entityDetails, err
}
