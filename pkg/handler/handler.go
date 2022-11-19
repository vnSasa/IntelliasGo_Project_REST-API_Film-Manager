package handler

import (
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/pkg/service"
	"github.com/gin-gonic/gin"
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
		auth.POST("/logout", h.logout)
	}
	
	apiAdmin := router.Group("/api-admin", h.adminIdentity)
	{
		directors := apiAdmin.Group("/director")
		{
			directors.POST("/create", h.createDiretor)
			directors.GET("/all", h.getAllDiretors)
			directors.GET("/:id", h.getDiretorById)
			directors.PUT("/:id", h.updateDiretor)
			directors.DELETE("/:id", h.deleteDiretor)
		}
		films := apiAdmin.Group("/films")
		{
			films.POST("/create", h.createFilm)
			films.GET("/all", h.getAllFilms)
			films.GET("/:id", h.getFilmById)
			films.PUT("/:id", h.updateFilm)
			films.DELETE("/:id", h.deleteFilm)
		}
	}
	
	apiUser := router.Group("/api-user", h.userIdentity)
	{
		apiUser.GET("/directors", h.getDirectors)
		apiUser.GET("/films", h.getFilms)
		
		favourite := apiUser.Group("/favourite")
		{
			favourite.POST("/create", h.createFavourite)
			favourite.GET("/all", h.getFavourite)
			favourite.GET("/:id", h.getFavouriteById)
			favourite.PUT("/:id", h.updateFavourite)
			favourite.DELETE("/:id", h.deleteFavourite)
		}
		
		wish := apiUser.Group("/wish")
		{
			wish.POST("/create", h.createWish)
			wish.GET("/all", h.getWish)
			wish.GET("/:id", h.getWishById)
			wish.PUT("/:id", h.updateWish)
			wish.DELETE("/:id", h.deleteWish)
		}
	}

	return router
}