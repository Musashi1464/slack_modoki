package methods

import (
	"back/models"
	"back/types"
	"log"

	"github.com/labstack/echo"
)

func ServerPostWorkspace() echo.HandlerFunc {
	return func(c echo.Context) error {
		form := new(types.WorkspaceForm)
		err := c.Bind(form)
		if err != nil {
			log.Println(err)
			return c.JSON(500, "server error")
		}
		if form.WorkspaceName == "" || form.UserName == "" || form.Email == "" || form.Password == "" {
			return c.JSON(400, "wrong input")
		}
		wid, err := models.GetWorkspaceIdByName(form.WorkspaceName)
		if err != nil {
			log.Println(err)
			return c.JSON(500, "server error")
		}
		if wid != 0 {
			return c.JSON(400, "already exists")
		} else {
			err = models.CreateWorkspace(form.WorkspaceName)
			if err != nil {
				log.Print(err)
				return c.JSON(500, "server error")
			}
			wid, err = models.GetWorkspaceIdByName(form.WorkspaceName)
			if err != nil {
				log.Println(err)
				return c.JSON(500, "server error")
			}
			err = models.CraeteUser(wid, form.UserName, form.Email, form.Password)
			if err != nil {
				log.Println(err)
				return c.JSON(500, "server error")
			}
		}
		return c.JSON(200, "OK")
	}
}
