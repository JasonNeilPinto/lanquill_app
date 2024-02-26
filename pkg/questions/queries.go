package questions

var DUPLICATE = `SELECT 
					passage_title,
					created_by 
					from sententia.reading_comprehension_info
				WHERE 
					passage_title = ? AND created_by = ?`

var INSERT_READ_SCRIPT = `INSERT INTO 
							sententia.reading_comprehension_info 
							(passage_title, passage_text, grade, complexity, max_time, created_date, created_by) 
						VALUES 
							(?, ?, ?, ?, ?, ?, ?)`

var INSERT_LISTEN_SCRIPT = `INSERT INTO 
							sententia.listening_comprehension_info 
							(passage_title, audio_path, grade, complexity, max_time, created_date, created_by) 
						VALUES 
							(?, ?, ?, ?, ?, ?, ?)`

var INSERT_PASSAGE_QUEST = `INSERT INTO 
								sententia.passage_questions 
								(passage_id, qus_for, qus_type, qus_text, option_1, option_2, option_3, option_4,option_5, answare) 
							VALUES 
								(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

var SELECT_PASSAGE_READ = `SELECT
						passage_id, passage_title, passage_text, grade, complexity, max_time
					FROM 
						sententia.reading_comprehension_info`

var SELECT_PASSAGE_QUEST = `SELECT
								question_id,qus_type, qus_text, IFNULL(option_1,''), IFNULL(option_2,''), IFNULL(option_3,''),IFNULL(option_4,''),IFNULL(option_5,''), answare
							FROM
								 sententia.passage_questions 
							WHERE 
								passage_id = ? and qus_for = ?`

var SELECT_PASSAGE_LISTEN = `SELECT
								passage_id, passage_title, audio_path, grade, complexity, max_time
							FROM 
								sententia.listening_comprehension_info`

var DELETE_READ = `DELETE a.*, b.* 
					FROM sententia.reading_comprehension_info a 
					LEFT JOIN sententia.passage_questions b 
					ON b.passage_id = a.passage_id 
					WHERE a.passage_id =  ?`

var DELETE_LISTEN = `DELETE a.*, b.* 
					FROM sententia.listening_comprehension_info a 
					LEFT JOIN sententia.passage_questions b 
					ON b.passage_id = a.passage_id 
					WHERE a.passage_id =  ?`

var SELECT_AUDIOFILE_PATH = `SELECT
								audio_path
							FROM 
								sententia.listening_comprehension_info
							WHERE passage_id = ?`

var VOCABULARY_DUPLICATE = `SELECT 
							qus_text 
						FROM 
							vocabulary_questions WHERE qus_text= ?`

var SELECT_VOCABULARY_QUEST = `SELECT 
									vocabulary_ques_id, grade, complexity, qus_type, qus_text, opt_1, opt_2, opt_3, opt_4, answare
								FROM 
									sententia.vocabulary_questions`

var INSERT_VOCABULARY_QUEST = `INSERT INTO 
									sententia.vocabulary_questions 
									(grade, complexity, qus_type, qus_text, opt_1, opt_2, opt_3, opt_4, answare) 
								VALUES 
									(?, ?, ?, ?, ?, ?, ?, ?, ?)`

var DELETE_VOCABULARY_QUEST = `DELETE from sententia.vocabulary_questions WHERE vocabulary_ques_id = ?`
