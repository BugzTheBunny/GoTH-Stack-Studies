package handler

import (
	"github.com/BugzTheBuinny/got/model"
	"github.com/BugzTheBuinny/got/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (uh UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Email: "google-account@gmail.com",
	}
	return renderTemplate(c, user.Show(u))
}
