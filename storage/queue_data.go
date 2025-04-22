package storage

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var queues_db *sql.DB

func Init() {
	var err error
	queues_db, err = sql.Open("sqlite3", "storage/queues.db")
	if err != nil{
		panic(err)
	}

	err = queues_db.Ping()
	if err != nil{
		panic(err)
	}

	q := `CREATE TABLE IF NOT EXISTS queue_members(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		chat_id INTEGER,
		topic_id INTEGER,
		user_id INTEGER,
		flname TEXT,
		username TEXT,
		priority INTEGER DEFAULT 0)`
	
	_, err = queues_db.Exec(q) 

	if err != nil{
		panic(err)
	}
}

func Add(chatID, topicID, userID int64, flname, username string, priority int) error{
	q := `
	INSERT INTO queue_members (chat_id, topic_id, user_id, flname, username, priority)
	VALUES(?, ?, ?, ?, ?, ?)
	`
	_, err := queues_db.Exec(q, chatID, topicID, userID, flname, username, priority)

	return err
}

func Pop(chatID , topicID int64) error{
	var q string
	/*if STATE == "normal" {
		q = `DELETE FROM queue_members
		WHERE id = (
		SELECT id FROM queue_members
		WHERE chat_id = ? AND topic_id = ?
		ORDER BY joined_at ASC
		LIMIT 1)`
	}else if STATE == "sorted"{
		q = `DELETE FROM queue_members
		WHERE id = (
		SELECT id FROM queue_members
		WHERE chat_id = ? AND topic_id = ? AND priopity >= 1
		ORDER BY prioriy ASC, joined_at ASC
		LIMIT 1)`
	}*/
	q = `DELETE FROM queue_members
		WHERE id = (
		SELECT id FROM queue_members
		WHERE chat_id = ? AND topic_id = ? AND priority = 0
		ORDER BY id ASC
		LIMIT 1)`


	_, err := queues_db.Exec(q,chatID, topicID)

	return err
}

func Remove(chatID, topicID int64) error{
	q := `DELETE FROM queue_members
	WHERE chat_id = ? AND topic_id = ?`

	_, err := queues_db.Exec(q,chatID, topicID)

	return err
}

func RemovePersone(chatID, topicID, userID int64) error{
	q := `
	DELETE FROM	queue_members
	WHERE chat_id = ? AND topic_id = ? AND user_id = ?
	`
	_, err := queues_db.Exec(q,chatID,topicID,userID)
	return err
}

func Get(chatID, topicID int64)([]string, error){
	q := `
		SELECT flname
		FROM queue_members
		WHERE chat_id = ? AND topic_id = ?
		ORDER BY id ASC
		`

	rows, err := queues_db.Query(q, chatID, topicID)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var queue []string
	for rows.Next(){
		var member string
		err := rows.Scan(&member)
		if err != nil{
			return nil, err
		}
		queue = append(queue, member)
	}
	return queue, nil
}

