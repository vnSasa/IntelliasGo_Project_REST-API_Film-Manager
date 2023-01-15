package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoute() *gin.Engine {
	router := gin.New()

	router.POST("/init-admin", h.InitAdmin)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/refresh", h.refreshSignIn)
		auth.POST("/logout", h.logout)
	}

	apiAdmin := router.Group("/api-admin", h.adminIdentity)
	{
		directors := apiAdmin.Group("/director")
		{
			directors.POST("/create", h.createDiretor)
			directors.GET("/all", h.getAllDiretors)
			directors.GET("/:id", h.getDiretorByID)
			directors.PUT("/:id", h.updateDiretor)
			directors.DELETE("/:id", h.deleteDiretor)
		}
		films := apiAdmin.Group("/films")
		{
			films.POST("/create", h.createFilm)
			films.GET("/all", h.getAllFilms)
			films.GET("/:id", h.getFilmByID)
			films.PUT("/:id", h.updateFilm)
			films.DELETE("/:id", h.deleteFilm)
			films.POST("/import", h.exportFilmstoCSV)
		}
	}

	apiUser := router.Group("/api-user", h.userIdentity)
	{
		films := apiUser.Group("/films")
		{
			films.GET("/all", h.getFilmsFilters)
			films.GET("/:id", h.getFilmByID)
			films.POST("/export", h.exportFilmstoCSV)

			favourite := films.Group("/favourite")
			{
				favourite.POST("/:id/add", h.addFavouriteFilm)
				favourite.GET("/all", h.getAllFavouriteFilms)
				favourite.DELETE("/:id", h.deleteFavourite)
				favourite.POST("/export", h.exportFtoCSV)
			}

			wish := films.Group("/wish")
			{
				wish.POST("/:id/add", h.addWishFilm)
				wish.GET("/all", h.getAllWishFilms)
				wish.DELETE("/:id", h.deleteWish)
				wish.POST("/export", h.exportWtoCSV)
			}
		}
	}

	return router
}
