package questions

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/Lanquill/Forge/pkg/auth"
	"github.com/Lanquill/Forge/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*-------------------------ADDING SPEECH SCRIPTS TO MONGODB---------------------------------*/
func StoreSpeechscript(tknData *auth.JWTData, speechScript SpeechScripts, words int) error {
	coll := db.MongoClient.Database("SLA_Content").Collection("speech_content")

	if err := TextSize(speechScript.ScriptText); err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	speechScript.AddedBy = "admin"

	insertData := bson.M{"User_Email": tknData.UserEmail,
		"User_ID":     tknData.UserID,
		"User_Type":   tknData.UserType,
		"title":       speechScript.Title,
		"script_text": speechScript.ScriptText,
		"words":       words,
		"added_by":    speechScript.AddedBy,
		"status":      "present",
		"complexity":  speechScript.Complexity,
		"grade":       speechScript.Grade,
		"audio_file":  speechScript.AudioFile,
	}

	_, err := coll.InsertOne(db.CTX, insertData)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	return nil
}

/*--------------ADDING READING SCRIPTS TO MYSQL--------------*/
func InsertReadScript(readingScripts PassageScripts, CreatedBy string) error {

	if err := TextSize(readingScripts.PassageText); err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	readingDB := PassageDB{}

	err := db.MySqlDB.QueryRow(DUPLICATE, readingScripts.PassageTitle, CreatedBy).Scan(&readingDB)
	if err == sql.ErrNoRows {

		stmt, err := db.MySqlDB.Prepare(INSERT_READ_SCRIPT)
		if err != nil {
			log.Println(err)
			return err
		}
		defer stmt.Close()

		result, err := stmt.Exec(
			readingScripts.PassageTitle,
			readingScripts.PassageText,
			readingScripts.Grade,
			readingScripts.Complexity,
			readingScripts.MaxTime,
			time.Now(),
			CreatedBy)

		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}

		passageId, err := result.LastInsertId()
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}

		stmt, err = db.MySqlDB.Prepare(INSERT_PASSAGE_QUEST)
		if err != nil {
			log.Println(err)
			return err
		}
		defer stmt.Close()

		quesFor := "reading"
		for _, quest := range readingScripts.Questions {

			if err := TextSize(quest.QuestionText); err != nil {
				log.Println("ERROR: ", err)
				return err
			}

			answer := quest.Answer
			quesText := quest.QuestionText
			quesType := quest.QuestionType
			var opt []string
			for _, val := range quest.Options {
				opt = append(opt, val.OptionText)
			}

			if len(opt) == 4 {
				_, err := stmt.Exec(
					passageId,
					quesFor,
					quesType,
					quesText,
					opt[0],
					opt[1],
					opt[2],
					opt[3],
					"",
					answer,
				)
				if err != nil {
					log.Println("ERROR: ", err)
					return err
				}

				rows, err := result.RowsAffected()
				if err != nil {
					log.Println("ERROR: ", err)
					return err
				}

				if rows == 0 {
					err := errors.New(ErrCreateFailure)
					return err
				}
			} else if len(opt) == 5 {
				_, err := stmt.Exec(
					passageId,
					quesFor,
					quesType,
					quesText,
					opt[0],
					opt[1],
					opt[2],
					opt[3],
					opt[4],
					answer,
				)
				if err != nil {
					log.Println("ERROR: ", err)
					return err
				}

				rows, err := result.RowsAffected()
				if err != nil {
					log.Println("ERROR: ", err)
					return err
				}

				if rows == 0 {
					err := errors.New(ErrCreateFailure)
					return err
				}

			} else if len(opt) == 2 {
				_, err := stmt.Exec(
					passageId,
					quesFor,
					quesType,
					quesText,
					opt[0],
					opt[1],
					"", "", "",
					answer,
				)
				if err != nil {
					log.Println("ERROR: ", err)
					return err
				}

				rows, err := result.RowsAffected()
				if err != nil {
					log.Println("ERROR: ", err)
					return err
				}

				if rows == 0 {
					err := errors.New(ErrCreateFailure)
					return err
				}

			} else {
				return nil
			}
		}
	} else if err != nil {
		err := errors.New(ErrPassageExists)
		return err
	}

	return nil

}

func InsertListeningScript(listening PassageScripts, CreatedBy string) error {

	if err := TextSize(listening.PassageText); err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	listeningDB := PassageDB{}

	err := db.MySqlDB.QueryRow(DUPLICATE, listening.PassageTitle, CreatedBy).Scan(&listeningDB)
	if err == sql.ErrNoRows {

		stmt, err := db.MySqlDB.Prepare(INSERT_LISTEN_SCRIPT)
		if err != nil {
			log.Println(err)
			return err
		}
		defer stmt.Close()

		result, err := stmt.Exec(
			listening.PassageTitle,
			listening.AudioFile,
			listening.Grade,
			listening.Complexity,
			listening.MaxTime,
			time.Now(),
			CreatedBy)

		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}

		passageId, err := result.LastInsertId()
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}

		stmt, err = db.MySqlDB.Prepare(INSERT_PASSAGE_QUEST)
		if err != nil {
			log.Println(err)
			return err
		}
		defer stmt.Close()

		quesFor := "listening"
		for _, quest := range listening.Questions {

			if err := TextSize(quest.QuestionText); err != nil {
				log.Println("ERROR: ", err)
				return err
			}

			answer := quest.Answer
			quesText := quest.QuestionText
			quesType := quest.QuestionType
			var opt []string
			for _, val := range quest.Options {
				opt = append(opt, val.OptionText)
			}

			if len(opt) == 4 {
				_, err := stmt.Exec(
					passageId,
					quesFor,
					quesType,
					quesText,
					opt[0],
					opt[1],
					opt[2],
					opt[3],
					"",
					answer,
				)
				if err != nil {
					log.Println("ERROR: ", err)
					return err
				}

				rows, err := result.RowsAffected()
				if err != nil {
					log.Println("ERROR: ", err)
					return err
				}

				if rows == 0 {
					err := errors.New(ErrCreateFailure)
					return err
				}

			} else if len(opt) == 5 {
				_, err := stmt.Exec(
					passageId,
					quesFor,
					quesType,
					quesText,
					opt[0],
					opt[1],
					opt[2],
					opt[3],
					opt[4],
					answer,
				)
				if err != nil {
					log.Println("ERROR: ", err)
					return err
				}

				rows, err := result.RowsAffected()
				if err != nil {
					log.Println("ERROR: ", err)
					return err
				}

				if rows == 0 {
					err := errors.New(ErrCreateFailure)
					return err
				}

			} else if len(opt) == 2 {
				_, err := stmt.Exec(
					passageId,
					quesFor,
					quesType,
					quesText,
					opt[0],
					opt[1],
					"", "", "",
					answer,
				)
				if err != nil {
					log.Println("ERROR: ", err)
					return err
				}
				rows, err := result.RowsAffected()
				if err != nil {
					log.Println("ERROR: ", err)
					return err
				}

				if rows == 0 {
					err := errors.New(ErrCreateFailure)
					return err
				}
			} else if len(opt) == 0 {
				_, err := stmt.Exec(
					passageId,
					quesFor,
					quesType,
					quesText,
					"", "", "", "", "",
					answer,
				)
				if err != nil {
					log.Println("ERROR: ", err)
					return err
				}

				rows, err := result.RowsAffected()
				if err != nil {
					log.Println("ERROR: ", err)
					return err
				}

				if rows == 0 {
					err := errors.New(ErrCreateFailure)
					return err
				}

			} else {
				return nil
			}
		}
	} else if err != nil {
		err := errors.New(ErrPassageExists)
		return err
	}

	return nil

}

/*-----------------------GET THE SPEECH PASSAGES---------------------------------*/
func GetSpeechData() ([]SpeechResp, error) {
	coll := db.MongoClient.Database("SLA_Content").Collection("speech_content")

	speechResp := []SpeechResp{}
	filter := bson.M{}
	opts := options.Find().
		SetProjection(bson.M{
			"_id":         1,
			"title":       1,
			"script_text": 1,
			"complexity":  1,
			"grade":       1,
			"audio_file":  1,
		})

	cursor, err := coll.Find(db.CTX, filter, opts)
	if err != nil {
		log.Println("ERROR: ", err)
		return speechResp, err
	}

	if err = cursor.All(db.CTX, &speechResp); err != nil {
		log.Println("ERROR: ", err)
		return speechResp, err
	}

	return speechResp, nil
}

func GetReadScripts() ([]PassageInfo, error) {
	reading := []PassageInfo{}

	stmtStr := SELECT_PASSAGE_READ

	stmt, err := db.MySqlDB.Prepare(stmtStr)
	if err != nil {
		log.Println("ERROR: ", err)
		return reading, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println("ERROR: ", err)
		return reading, err
	}
	defer rows.Close()

	readingDBs := []PassageDB{}

	for rows.Next() {
		var readingDB PassageDB
		if err := rows.Scan(&readingDB.ID,
			&readingDB.PassageTitle,
			&readingDB.PassageText,
			&readingDB.Grade,
			&readingDB.Complexity,
			&readingDB.MaxTime); err != nil {
			log.Println("ERROR: ", err)
			return reading, err
		}
		readingDBs = append(readingDBs, readingDB)
	}

	for _, readingDB := range readingDBs {

		questions := []QuestionInfo{}

		var eachRead PassageInfo
		eachRead.ID = readingDB.ID
		eachRead.Type = "reading"
		eachRead.PassageTitle = readingDB.PassageTitle
		eachRead.PassageText = readingDB.PassageText
		eachRead.Grade = readingDB.Grade
		eachRead.Complexity = readingDB.Complexity
		eachRead.MaxTime = readingDB.MaxTime

		stmtStr := SELECT_PASSAGE_QUEST

		stmt, err := db.MySqlDB.Prepare(stmtStr)
		if err != nil {
			log.Println("ERROR: ", err)
			return reading, err
		}
		defer stmt.Close()

		rows, err = stmt.Query(readingDB.ID, eachRead.Type)
		if err != nil {
			log.Println("ERROR: ", err)
			return reading, err
		}
		defer rows.Close()

		for rows.Next() {
			var eachQuestion QuestionInfo
			if err := rows.Scan(
				&eachQuestion.QuestID,
				&eachQuestion.QuestType,
				&eachQuestion.QuestText,
				&eachQuestion.Option1,
				&eachQuestion.Option2,
				&eachQuestion.Option3,
				&eachQuestion.Option4,
				&eachQuestion.Option5,
				&eachQuestion.Answer); err != nil {
				log.Println("ERROR: ", err)
				return reading, err
			}

			questions = append(questions, eachQuestion)
		}
		eachRead.QuestionDetails = questions
		reading = append(reading, eachRead)
	}
	return reading, nil
}

func GetListenScripts() ([]PassageInfo, error) {
	listening := []PassageInfo{}

	stmtStr := SELECT_PASSAGE_LISTEN

	stmt, err := db.MySqlDB.Prepare(stmtStr)
	if err != nil {
		log.Println("ERROR: ", err)
		return listening, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println("ERROR: ", err)
		return listening, err
	}
	defer rows.Close()

	listeningDBs := []PassageDB{}

	for rows.Next() {
		var listeningDB PassageDB
		if err := rows.Scan(&listeningDB.ID,
			&listeningDB.PassageTitle,
			&listeningDB.AudioPath,
			&listeningDB.Grade,
			&listeningDB.Complexity,
			&listeningDB.MaxTime); err != nil {
			log.Println("ERROR: ", err)
			return listening, err
		}
		listeningDBs = append(listeningDBs, listeningDB)
	}

	for _, listeningDB := range listeningDBs {

		questions := []QuestionInfo{}

		var eachListen PassageInfo
		eachListen.ID = listeningDB.ID
		eachListen.Type = "listening"
		eachListen.PassageTitle = listeningDB.PassageTitle
		eachListen.AudioFile = listeningDB.AudioPath
		eachListen.Grade = listeningDB.Grade
		eachListen.Complexity = listeningDB.Complexity
		eachListen.MaxTime = listeningDB.MaxTime

		stmtStr := SELECT_PASSAGE_QUEST

		stmt, err := db.MySqlDB.Prepare(stmtStr)
		if err != nil {
			log.Println("ERROR: ", err)
			return listening, err
		}
		defer stmt.Close()

		rows, err = stmt.Query(listeningDB.ID, eachListen.Type)
		if err != nil {
			log.Println("ERROR: ", err)
			return listening, err
		}
		defer rows.Close()

		for rows.Next() {
			var eachQuestion QuestionInfo
			if err := rows.Scan(
				&eachQuestion.QuestID,
				&eachQuestion.QuestType,
				&eachQuestion.QuestText,
				&eachQuestion.Option1,
				&eachQuestion.Option2,
				&eachQuestion.Option3,
				&eachQuestion.Option4,
				&eachQuestion.Option5,
				&eachQuestion.Answer); err != nil {
				log.Println("ERROR: ", err)
				return listening, err
			}
			questions = append(questions, eachQuestion)
		}
		eachListen.QuestionDetails = questions
		listening = append(listening, eachListen)
	}
	return listening, nil

}

func UpdateSpeech(speechReq SpeechResp) error {
	coll := db.MongoClient.Database("SLA_Content").Collection("speech_content")

	if err := TextSize(speechReq.ScriptText); err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	Getaudio := SpeechResp{}

	cursor, err := coll.Find(db.CTX, bson.M{"_id": speechReq.ID}, options.Find().
		SetProjection(bson.M{
			"audio_file": 1,
		}))

	if err != nil {
		log.Println("ERROR: ", err)
		err := errors.New(ErrNoDocumentsExits)
		return err
	}

	for cursor.Next(db.CTX) {
		if err = cursor.Decode(&Getaudio); err != nil {
			log.Println("ERROR: ", err)
		}
	}
	if err := cursor.Err(); err != nil {
		log.Println(err)
	}

	filter := bson.D{{Key: "_id", Value: speechReq.ID}}
	updates := bson.M{}
	update := bson.M{"$set": updates}

	if speechReq.Title != "" {
		updates["title"] = speechReq.Title
	}

	if speechReq.ScriptText != "" {
		updates["words"] = Get_words(speechReq.ScriptText)
		updates["script_text"] = speechReq.ScriptText
	}

	if speechReq.Grade != "" {
		updates["grade"] = speechReq.Grade
	}

	if speechReq.Complexity != "" {
		updates["complexity"] = speechReq.Complexity
	}

	if speechReq.AudioFile != "" {
		updates["audio_file"] = speechReq.AudioFile
		if err := RemoveSpeechAudio(Getaudio.AudioFile); err != nil {
			log.Println("ERROR: ", err)
			return err
		}
	}

	result, err := coll.UpdateOne(db.CTX, filter, update)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	if result.ModifiedCount == 1 {
		return nil
	} else {
		return errors.New(ErrUpdateFailure)
	}

}

func (readReq PassageInfo) UpdateRead() error {

	if err := TextSize(readReq.PassageText); err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	query_reading := "UPDATE sententia.reading_comprehension_info SET"
	params := []interface{}{}
	updates := []string{}

	if readReq.PassageTitle != "" {
		updates = append(updates, "passage_title= ?")
		params = append(params, readReq.PassageTitle)
	}

	if readReq.PassageText != "" {
		updates = append(updates, "passage_text= ?")
		params = append(params, readReq.PassageText)
	}

	if readReq.Grade != "" {
		updates = append(updates, "grade= ?")
		params = append(params, readReq.Grade)
	}

	if readReq.Complexity != "" {
		updates = append(updates, "complexity= ?")
		params = append(params, readReq.Complexity)
	}

	if readReq.MaxTime != 0 {
		updates = append(updates, "max_time= ?")
		params = append(params, readReq.MaxTime)
	}

	query_reading += " " + joinStrings(", ", updates)
	query_reading += " WHERE passage_id = ?"
	params = append(params, readReq.ID)

	_, err := db.MySqlDB.Exec(query_reading, params...)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	for _, val := range readReq.QuestionDetails {
		query_passage := "UPDATE sententia.passage_questions SET"
		params := []interface{}{}
		updates := []string{}

		if val.Answer != "" {
			updates = append(updates, "answare= ?")
			params = append(params, val.Answer)
		}

		if val.Option1 != "" {
			updates = append(updates, "option_1= ?")
			params = append(params, val.Option1)
		}

		if val.Option2 != "" {
			updates = append(updates, "option_2= ?")
			params = append(params, val.Option2)
		}

		if val.Option3 != "" {
			updates = append(updates, "option_3= ?")
			params = append(params, val.Option3)
		}

		if val.Option4 != "" {
			updates = append(updates, "option_4= ?")
			params = append(params, val.Option4)
		}

		if val.Option5 != "" {
			updates = append(updates, "option_5= ?")
			params = append(params, val.Option5)
		}

		if val.QuestType != "" {
			updates = append(updates, "qus_type= ?")
			params = append(params, val.QuestType)
		}

		if val.QuestText != "" {
			updates = append(updates, "qus_text= ?")
			params = append(params, val.QuestText)
		}
		query_passage += " " + joinStrings(", ", updates)
		query_passage += " WHERE passage_id = ? AND question_id = ?"
		params = append(params, readReq.ID)
		params = append(params, val.QuestID)

		_, err := db.MySqlDB.Exec(query_passage, params...)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}
	}

	return nil
}

func UpdateListen(listenReq PassageInfo) error {

	if err := TextSize(listenReq.PassageText); err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	query_listen := "UPDATE sententia.listening_comprehension_info SET"
	params := []interface{}{}
	updates := []string{}

	if listenReq.PassageTitle != "" {
		updates = append(updates, "passage_title= ?")
		params = append(params, listenReq.PassageTitle)
	}

	if listenReq.AudioFile != "" {
		updates = append(updates, "audio_path= ?")
		params = append(params, listenReq.AudioFile)
		err := RemoveAudioFile(listenReq.ID)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}
	}

	if listenReq.Grade != "" {
		updates = append(updates, "grade= ?")
		params = append(params, listenReq.Grade)
	}

	if listenReq.Complexity != "" {
		updates = append(updates, "complexity= ?")
		params = append(params, listenReq.Complexity)
	}

	if listenReq.MaxTime != 0 {
		updates = append(updates, "max_time= ?")
		params = append(params, listenReq.MaxTime)
	}

	query_listen += " " + joinStrings(", ", updates)
	query_listen += " WHERE passage_id = ?"
	params = append(params, listenReq.ID)

	result, err := db.MySqlDB.Exec(query_listen, params...)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	for _, val := range listenReq.QuestionDetails {
		query_passage := "UPDATE sententia.passage_questions SET"
		params := []interface{}{}
		updates := []string{}

		if val.Answer != "" {
			updates = append(updates, "answare= ?")
			params = append(params, val.Answer)
		}

		if val.Option1 != "" {
			updates = append(updates, "option_1= ?")
			params = append(params, val.Option1)
		}

		if val.Option2 != "" {
			updates = append(updates, "option_2= ?")
			params = append(params, val.Option2)
		}

		if val.Option3 != "" {
			updates = append(updates, "option_3= ?")
			params = append(params, val.Option3)
		}

		if val.Option4 != "" {
			updates = append(updates, "option_4= ?")
			params = append(params, val.Option4)
		}

		if val.Option5 != "" {
			updates = append(updates, "option_5= ?")
			params = append(params, val.Option5)
		}

		if val.QuestType != "" {
			updates = append(updates, "qus_type= ?")
			params = append(params, val.QuestType)
		}

		if val.QuestText != "" {
			updates = append(updates, "qus_text= ?")
			params = append(params, val.QuestText)
		}

		query_passage += " " + joinStrings(", ", updates)
		query_passage += " WHERE passage_id = ? AND question_id = ?"
		params = append(params, listenReq.ID)
		params = append(params, val.QuestID)

		result, err := db.MySqlDB.Exec(query_passage, params...)
		if err != nil {
			log.Println("ERROR:", err)
			return err
		}

		_, err = result.RowsAffected()
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}

	}

	return nil
}

func DeleteSpeech(id primitive.ObjectID) error {
	coll := db.MongoClient.Database("SLA_Content").Collection("speech_content")

	Getaudio := SpeechResp{}
	cursor, err := coll.Find(db.CTX, bson.M{"_id": id}, options.Find().
		SetProjection(bson.M{
			"audio_file	": 1,
		}))
	if err != nil {
		log.Println("ERROR: ", err)
		err := errors.New(ErrNoDocumentsExits)
		return err
	}

	for cursor.Next(db.CTX) {
		if err = cursor.Decode(&Getaudio); err != nil {
			log.Println("ERROR: ", err)
		}
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	if Getaudio.AudioFile != "" {
		if err := RemoveSpeechAudio(Getaudio.AudioFile); err != nil {
			log.Println("ERROR: ", err)
			return err
		}
	}

	filter := bson.D{{Key: "_id", Value: id}}

	_, err = coll.DeleteOne(db.CTX, filter)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}
	return nil
}

func DeleteReading(id int64) error {
	stmtStr := DELETE_READ

	stmt, err := db.MySqlDB.Prepare(stmtStr)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	if rows == 0 {
		err := errors.New(ErrDeleteFailure)
		return err
	}

	return nil
}

func DeleteListening(id int64) error {

	stmtStr := DELETE_LISTEN

	err := RemoveAudioFile(id)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	stmt, err := db.MySqlDB.Prepare(stmtStr)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}
	if rows == 0 {
		err := errors.New(ErrDeleteFailure)
		return err
	}

	return nil
}

func GetGrammar() ([]Grammar, error) {
	coll := db.MongoClient.Database("eklaas").Collection("tbl_question")

	grammar := []Grammar{}
	filter := bson.M{"options": bson.M{"$ne": nil}}
	opts := options.Find().
		SetProjection(bson.M{
			"serial_number":       1,
			"topic_id":            1,
			"question_id":         1,
			"question_type":       1,
			"complexity_smc":      1,
			"grade":               1,
			"question_text":       1,
			"question_created_by": 1,
			"question_created_on": 1,
			"answer_id":           1,
			"options":             1,
			"_id":                 1,
		})

	cursor, err := coll.Find(db.CTX, filter, opts)
	if err != nil {
		log.Println("ERROR: ", err)
		return grammar, err
	}

	if err = cursor.All(db.CTX, &grammar); err != nil {
		log.Println("ERROR: ", err)
		return grammar, err
	}

	return grammar, nil
}

func AddGrammar(grammar Grammar, tknData *auth.JWTData) error {
	coll := db.MongoClient.Database("eklaas").Collection("tbl_question")

	if err := TextSize(grammar.QuestionText); err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	serialNumber, err := GetSerialNumber(coll)
	if err != nil {
		log.Println(err)
	}
	grammar.SerialNumber = serialNumber + 1
	grammar.QuestionId = "QUEST-" + GenerateRandomString(15)

	for i := range grammar.Options {
		grammar.Options[i].OptionId = "OPTION-" + GenerateRandomString(15)
		if grammar.AnswerText == grammar.Options[i].OptionText {
			grammar.AnswerId = grammar.Options[i].OptionId
		}

	}
	grammar.CreatedBy = tknData.UserName
	grammar.CreatedOn = time.Now()

	Insert := bson.M{"serial_number": grammar.SerialNumber,
		"topic_id":                 grammar.TopicId,
		"question_id":              grammar.QuestionId,
		"question_type":            grammar.QuestionType,
		"complexity_smc":           grammar.Complexity,
		"grade":                    grammar.Grade,
		"question_text":            grammar.QuestionText,
		"question_created_on":      grammar.CreatedOn,
		"question_created_by_name": grammar.CreatedBy,
		"answer_id":                grammar.AnswerId,
		"options":                  grammar.Options,
	}

	result, err := coll.InsertOne(db.CTX, Insert)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}
	if result.InsertedID == nil {
		return errors.New(ErrCreateFailure)
	}

	return nil

}

func UpdateGrammar(grammar Grammar) error {
	coll := db.MongoClient.Database("eklaas").Collection("tbl_question")

	if err := TextSize(grammar.QuestionText); err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	filter := bson.D{{Key: "_id", Value: grammar.ID}}
	updates := bson.M{}
	update := bson.M{"$set": updates}

	if grammar.QuestionText != "" {
		updates["question_text"] = grammar.QuestionText
	}

	if grammar.QuestionType != 0 {
		updates["question_type"] = grammar.QuestionType
	}

	if grammar.Grade != "" {
		updates["grade"] = grammar.Grade
	}

	if grammar.Complexity != "" {
		updates["complexity_smc"] = grammar.Complexity
	}

	if grammar.AnswerId != "" {
		updates["answer_id"] = grammar.AnswerId
	}

	if grammar.Options != nil {
		filter := bson.M{
			"options": bson.M{
				"$elemMatch": bson.M{
					"option_id": grammar.AnswerId,
				},
			},
		}

		count, err := coll.CountDocuments(db.CTX, filter)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}
		if count != 1 {
			return errors.New(ErrUpdateFailure)
		}

		updates["options"] = grammar.Options
	}

	updateResult, err := coll.UpdateOne(db.CTX, filter, update)
	if err != nil {
		log.Fatal(err)
	}

	if updateResult.ModifiedCount == 1 {
		return nil
	} else {
		return errors.New(ErrUpdateFailure)
	}

}

func DeleteGrammar(id primitive.ObjectID) error {
	coll := db.MongoClient.Database("eklaas").Collection("tbl_question")

	filter := bson.D{{Key: "_id", Value: id}}

	result, err := coll.DeleteOne(db.CTX, filter)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	if result.DeletedCount == 0 {
		err := errors.New(ErrDeleteFailure)
		return err
	}

	return nil
}

func AddVocabulary(vocabulary Vocalbulary) error {

	if err := TextSize(vocabulary.QuestionText); err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	err := db.MySqlDB.QueryRow(VOCABULARY_DUPLICATE, vocabulary.QuestionText).Scan()
	if err == sql.ErrNoRows {

		stmt, err := db.MySqlDB.Prepare(INSERT_VOCABULARY_QUEST)
		if err != nil {
			log.Println(err)
			return err
		}
		defer stmt.Close()

		result, err := stmt.Exec(
			vocabulary.Grade,
			vocabulary.Complexity,
			vocabulary.QuestionType,
			vocabulary.QuestionText,
			vocabulary.Option1,
			vocabulary.Option2,
			vocabulary.Option3,
			vocabulary.Option4,
			vocabulary.Answer)
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}

		rows, err := result.RowsAffected()
		if err != nil {
			log.Println("ERROR: ", err)
			return err
		}

		if rows == 0 {
			err := errors.New(ErrCreateFailure)
			return err
		}

	} else if err != nil {
		err := errors.New(ErrPassageExists)
		return err
	}

	return nil
}

func GetVocalbulary() ([]Vocalbulary, error) {
	vocabulary := []Vocalbulary{}

	stmt, err := db.MySqlDB.Prepare(SELECT_VOCABULARY_QUEST)
	if err != nil {
		log.Println("ERROR: ", err)
		return vocabulary, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Println("ERROR: ", err)
		return vocabulary, err
	}
	defer rows.Close()

	for rows.Next() {
		var vocabularydb Vocalbulary
		if err := rows.Scan(&vocabularydb.QuestionID,
			&vocabularydb.Grade,
			&vocabularydb.Complexity,
			&vocabularydb.QuestionType,
			&vocabularydb.QuestionText,
			&vocabularydb.Option1,
			&vocabularydb.Option2,
			&vocabularydb.Option3,
			&vocabularydb.Option4,
			&vocabularydb.Answer); err != nil {
			log.Println(err)
		}
		vocabulary = append(vocabulary, vocabularydb)
	}

	return vocabulary, nil

}

func UpdateVocalbulary(vocabulary Vocalbulary) error {

	if err := TextSize(vocabulary.QuestionText); err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	query_reading := "UPDATE sententia.vocabulary_questions SET"
	params := []interface{}{}
	updates := []string{}

	if vocabulary.QuestionType != "" {
		updates = append(updates, "qus_type= ?")
		params = append(params, vocabulary.QuestionType)
	}

	if vocabulary.QuestionText != "" {
		updates = append(updates, "qus_text= ?")
		params = append(params, vocabulary.QuestionText)
	}

	if vocabulary.Grade != "" {
		updates = append(updates, "grade= ?")
		params = append(params, vocabulary.Grade)
	}

	if vocabulary.Complexity != "" {
		updates = append(updates, "complexity= ?")
		params = append(params, vocabulary.Complexity)
	}

	if vocabulary.Option1 != "" {
		updates = append(updates, "opt_1= ?")
		params = append(params, vocabulary.Option1)
	}

	if vocabulary.Option2 != "" {
		updates = append(updates, "opt_2= ?")
		params = append(params, vocabulary.Option2)
	}

	if vocabulary.Option3 != "" {
		updates = append(updates, "opt_3= ?")
		params = append(params, vocabulary.Option3)
	}

	if vocabulary.Option4 != "" {
		updates = append(updates, "opt_4= ?")
		params = append(params, vocabulary.Option4)
	}

	if vocabulary.Answer != "" {
		updates = append(updates, "answare= ?")
		params = append(params, vocabulary.Answer)
	}

	query_reading += " " + joinStrings(", ", updates)
	query_reading += " WHERE vocabulary_ques_id	= ?"
	params = append(params, vocabulary.QuestionID)

	result, err := db.MySqlDB.Exec(query_reading, params...)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	if rows == 0 {
		err := errors.New(ErrUpdateFailure)
		return err
	}

	return nil
}

func DeleteVocabulary(QuestionId int64) error {

	stmt, err := db.MySqlDB.Prepare(DELETE_VOCABULARY_QUEST)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(QuestionId)
	if err != nil {
		log.Println(err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	if rows == 0 {
		err := errors.New(ErrDeleteFailure)
		return err
	}

	return nil
}
