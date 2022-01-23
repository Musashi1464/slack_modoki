package models

import (
	"back/types"
	"database/sql"
	"log"
	"time"
)

func CraeteUser(workspaceid int, name string, email string, password string) (err error) {
	cmd := `INSERT INTO users (
		workspace_id,
		name,
		email,
		password,
		created_at,
		updated_at) VALUE (?, ?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd,
		workspaceid,
		name,
		email,
		Encrypt(password),
		time.Now(),
		time.Now())
	if err != nil {
		log.Print(err)
	}
	return err
}

func GetUserByWorkspaceIdEmail(workspaceid int, email string) (user types.User, err error) {
	user = types.User{}
	cmd := `SELECT id, workspace_id, name, email, password, created_at, updated_at
						FROM users WHERE workspace_id = ? AND email = ?`

	err = Db.QueryRow(cmd, workspaceid, email).Scan(
		&user.ID,
		&user.WorkspaceID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		} else {
			log.Println(err)
		}
	}

	return user, err
}
