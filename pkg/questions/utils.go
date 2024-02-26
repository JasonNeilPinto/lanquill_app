package questions

import (
	"database/sql"
	"errors"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"github.com/Lanquill/Forge/pkg/db"
	"github.com/Lanquill/Forge/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Get_words(text string) int {
	words := strings.Fields(text)
	return len(words)
}

func IsNull(ns sql.NullString) sql.NullString {
	if !ns.Valid {
		return sql.NullString{}
	}
	return ns
}

func TextSize(s string) error {
	maxCharacters := 3000
	if len(s) > maxCharacters {
		return errors.New(ErrWordLimit)
	}
	return nil
}

func joinStrings(delimiter string, strs []string) string {
	result := ""
	for i, str := range strs {
		if i > 0 {
			result += delimiter
		}
		result += str
	}
	return result
}

func CreateFileName() string {
	temp := time.Now().Format("2006-01-02 15:04:05")
	str := strings.Replace(strings.Replace(strings.Replace(temp, " ", "_", -1), "-", "", -1), ":", "", -1)
	return str

}

func RemoveAudioFile(id int64) error {

	var fileName string
	stmt, err := db.MySqlDB.Prepare(SELECT_AUDIOFILE_PATH)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&fileName)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}
	err = os.Remove(utils.STATICPATH + "/audio_files/" + fileName)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	return nil
}

func GetSerialNumber(coll *mongo.Collection) (int32, error) {
	grammar := Grammar{}

	findOptions := options.Find()
	findOptions.SetProjection(bson.M{"serial_number": 1, "_id": 0})
	findOptions.SetSort(bson.D{{Key: "serial_number", Value: -1}})
	findOptions.SetLimit(1)

	cursor, err := coll.Find(db.CTX, bson.M{}, findOptions)
	if err != nil {
		log.Println(err)
	}
	for cursor.Next(db.CTX) {
		if err = cursor.Decode(&grammar); err != nil {
			log.Println("ERROR: ", err)
		}
	}
	if err := cursor.Err(); err != nil {
		log.Println("ERROR: ", err)
	}

	return grammar.SerialNumber, nil
}

func GenerateRandomString(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		randomIndex := r.Intn(len(charset))
		result[i] = charset[randomIndex]
	}

	return string(result)
}

func RemoveSpeechAudio(fileName string) error {
	err := os.Remove(utils.STATICPATH + "/speech_audio/" + fileName)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	return nil
}

func AddAudioFile(file multipart.File, filename string, AudioFor string) (string, error) {
	if file != nil {
		if AudioFor == "listen" {
			filename = CreateFileName() + "." + filename
			audioFilePath := utils.STATICPATH + "/audio_files/" + filename

			f, err := os.OpenFile(audioFilePath, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				log.Println(err)
				return "", err
			}
			defer f.Close()
			io.Copy(f, file)
			return filename, nil
		} else if AudioFor == "speech" {
			filename = CreateFileName() + "." + filename
			audioFilePath := utils.STATICPATH + "/speech_audio/" + filename

			f, err := os.OpenFile(audioFilePath, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				log.Println(err)
				return "", err
			}
			defer f.Close()
			io.Copy(f, file)
			return filename, nil
		}
	}
	return "", nil

}
