package main

import (
	"back/models"
	"back/servers"
	"back/utils"
	"fmt"
)

func main() {
	utils.LoggingSettings()
	fmt.Println(models.Db)

	servers.StartMainServer()

	// wid := models.GetWorkspaceIdByName("test_workspace")
	// fmt.Println(wid)
	// user := models.GetUserByWorkspaceIdEmail(wid, "test@test.com")
	// fmt.Println(user)

}
