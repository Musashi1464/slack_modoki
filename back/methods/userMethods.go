package methods

import (
	"back/models"
	"back/types"
	"log"

	"github.com/labstack/echo"
)

func ServerPostLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		form := new(types.LoginForm)
		loginreturn := types.LoginReturn{Success: false, Message: ""}
		err := c.Bind(form)
		if err != nil {
			log.Println(err)
			loginreturn.Message = "server error"
			return c.JSON(500, loginreturn)
		}
		if form.WorkspaceName == "" || form.Email == "" || form.Password == "" {
			loginreturn.Message = "wrong input"
			return c.JSON(400, loginreturn)
		}
		wid, err := models.GetWorkspaceIdByName(form.WorkspaceName)
		if err != nil {
			log.Println(err)
			loginreturn.Message = "server error"
			return c.JSON(500, loginreturn)
		}
		if wid == 0 {
			loginreturn.Message = "wrong workspace_name, email, password"
			return c.JSON(400, loginreturn)
		}
		user, err := models.GetUserByWorkspaceIdEmail(wid, form.Email)
		if err != nil {
			log.Println(err)
			loginreturn.Message = "server error"
			return c.JSON(500, loginreturn)
		}
		if user.Password == models.Encrypt(form.Password) {
			loginreturn.Success = true
		}
		return c.JSON(200, loginreturn)
	}
}
