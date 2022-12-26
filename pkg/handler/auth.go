package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager"
	"net/http"
	"os"
	"strings"
)

func (h *Handler) InitAdmin(c *gin.Context) {
	input := app.User{
		Login:    os.Getenv("ADMIN_LOGIN"),
		Password: os.Getenv("ADMIN_PASSWORD"),
		Age:      os.Getenv("ADMIN_AGE"),
	}

	id, err := h.services.Authorization.CreateAdmin(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signUp(c *gin.Context) {
	var input app.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type userDataInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password_hash" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input userDataInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Login, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) logout(c *gin.Context) {
	userID, userAge, err := h.parseAuthHeader(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	favouriteFilms, _ := h.services.FavouriteFilms.GetAllFavouriteFilms(userID)
	wishFilms, _ := h.services.WishFilms.GetAllWishFilms(userID)

	var input userDataInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	if err := h.services.Authorization.GetUser(input.Login, input.Password); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	if err := h.services.Authorization.DeleteUser(userID); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	user := app.User{
		Login:    input.Login,
		Password: input.Password,
		Age:      userAge,
	}

	var id int
	if strings.Compare(input.Login, os.Getenv("ADMIN_LOGIN")) == 0 {
		id, err = h.services.Authorization.CreateAdmin(user)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())

			return
		}
	} else {
		id, err = h.services.Authorization.CreateUser(user)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())

			return
		}

		for _, v := range favouriteFilms {
			_, err := h.services.FavouriteFilms.AddFavouriteFilm(id, v.ID)
			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())

				return
			}
		}

		for _, vv := range wishFilms {
			_, err := h.services.WishFilms.AddWishFilm(id, vv.ID)
			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())

				return
			}
		}
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
