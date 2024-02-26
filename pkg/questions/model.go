package questions

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SpeechScripts struct {
	Title      string `json:"title" bson:"title"`
	ScriptText string `json:"scriptText" bson:"script_text"`
	AddedBy    string `json:"addedBy" bson:"added_by"`
	Complexity string `json:"complexity" bson:"complexity"`
	Grade      string `json:"grade" bson:"grade"`
	AudioFile  string `json:"audioFile" bson:"audio_file"`
}
type SpeechResp struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Title      string             `json:"title" bson:"title"`
	ScriptText string             `json:"scriptText" bson:"script_text"`
	Grade      string             `json:"grade" bson:"grade"`
	Complexity string             `json:"complexity" bson:"complexity"`
	AudioFile  string             `json:"audioFile" bson:"audio_file"`
}

type PassageScripts struct {
	PassageTitle string        `json:"passageTitle"`
	Complexity   string        `json:"complexity"`
	CreatedBy    string        `json:"createdBy"`
	Grade        string        `json:"grade"`
	MaxTime      int64         `json:"maxTime"`
	PassageText  string        `json:"passageText"`
	AudioFile    string        `json:"audioFile"`
	Questions    []QuestionReq `json:"questions"`
}

type QuestionReq struct {
	Answer       string       `json:"answer"`
	Options      []OptionsReq `json:"options"`
	QuestionText string       `json:"questionText"`
	QuestionType string       `json:"questionType" `
}

type OptionsReq struct {
	OptionText string `json:"optionText" bson:"option_text"`
}

type PassageDB struct {
	ID           int64  `db:"passage_id"`
	PassageTitle string `db:"passage_title"`
	PassageText  string `db:"passage_text"`
	Grade        string `db:"grade"`
	Complexity   string `db:"complexity"`
	MaxTime      int64  `db:"max_time"`
	CreatedDate  string `db:"created_date"`
	CreatedBy    string `db:"created_by"`
	AudioPath    string `db:"audio_path"`
}

type PassageInfo struct {
	ID              int64          `json:"passageId"`
	Type            string         `json:"type"`
	PassageTitle    string         `json:"passageTitle"`
	PassageText     string         `json:"passageText"`
	Grade           string         `json:"grade"`
	Complexity      string         `json:"complexity"`
	MaxTime         int64          `json:"maxTime"`
	AudioFile       string         `json:"audioFile"`
	QuestionDetails []QuestionInfo `json:"questionDetails"`
}

type QuestionInfo struct {
	QuestID   int64  `json:"questId"`
	QuestType string `json:"questType"`
	QuestText string `json:"questText"`
	Option1   string `json:"option_1"`
	Option2   string `json:"option_2"`
	Option3   string `json:"option_3"`
	Option4   string `json:"option_4"`
	Option5   string `json:"option_5"`
	Answer    string `json:"answer"`
}

type Grammar struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	SerialNumber int32              `json:"serial_number" bson:"serial_number"`
	TopicId      string             `json:"topicId,omitempty" bson:"topic_id"`
	QuestionId   string             `json:"questionId" bson:"question_id"`
	QuestionText string             `json:"questionText,omitempty" bson:"question_text"`
	QuestionType int64              `json:"questionType,omitempty" bson:"question_type"`
	Grade        string             `json:"grade,omitempty" bson:"grade"`
	Complexity   string             `json:"complexity,omitempty" bson:"complexity_smc"`
	CreatedBy    string             `json:"created_by,omitempty" bson:"question_created_by"`
	CreatedOn    time.Time          `json:"created_on,omitempty" bson:"question_created_on"`
	AnswerId     string             `json:"answerId" bson:"answer_id"`
	AnswerText   string             `json:"answerText,omitempty"`
	Options      []GrammarOpts      `json:"options"`
}

type GrammarOpts struct {
	OptionId   string      `json:"optionId" bson:"option_id"`
	OptionText interface{} `json:"optionText" bson:"option_text"`
}

type Vocalbulary struct {
	QuestionID   int64  `json:"questionId"`
	Grade        string `json:"grade"`
	Complexity   string `json:"complexity"`
	QuestionType string `json:"questionType"`
	QuestionText string `json:"questionText"`
	Option1      string `json:"option1"`
	Option2      string `json:"option2"`
	Option3      string `json:"option3"`
	Option4      string `json:"option4"`
	Answer       string `json:"answer"`
}
