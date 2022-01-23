package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
var err error

const (
	tableNameWorkspace = "workspaces"
	tableNameUser      = "users"
)

func init() {
	Db, err = sql.Open("mysql", "developer:password@/slack_modoki?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	cmdW := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INT PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(100),
		created_at DATETIME,
		updated_at DATETIME)`, tableNameWorkspace)

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INT PRIMARY KEY AUTO_INCREMENT,
		workspace_id INT,
		name VARCHAR(100),
		email VARCHAR(100),
		password VARCHAR(100),
		created_at DATETIME,
		updated_at DATETIME)`, tableNameUser)

	Db.Exec(cmdW)
	Db.Exec(cmdU)
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
