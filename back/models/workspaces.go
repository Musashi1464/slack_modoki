package models

import (
	"database/sql"
	"log"
	"time"
)

func CreateWorkspace(name string) (err error) {
	cmd := `INSERT INTO workspaces (
		name,
		created_at,
		updated_at) VALUE (?, ?, ?)`

	_, err = Db.Exec(cmd, name, time.Now(), time.Now())
	if err != nil {
		log.Println(err)
	}
	return err
}

// ワークスペースがあればそのIDを、なければ0を返す
func GetWorkspaceIdByName(name string) (workspaceid int, err error) {
	cmd := `SELECT ifnull(id, 0) FROM workspaces WHERE name = ?`
	err = Db.QueryRow(cmd, name).Scan(
		&workspaceid,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			log.Println(err)
		}
	}
	return workspaceid, err
}
