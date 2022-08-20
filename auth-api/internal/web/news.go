package web

import (
	"github.com/SebasGA19/spAInews/auth-api/internal/controller"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type (
	NewsResponse struct {
		CurrentPage     int64                `json:"current-page"`
		NumberOfResults int                  `json:"number-of-results"`
		News            []controller.Article `json:"news"`
	}
)

func (backend *Backend) LatestNews(ctx *gin.Context) {
	rawPage := ctx.Param(NewsPageVariable)
	page, parsePageError := strconv.ParseInt(rawPage, 10, 64)
	if parsePageError != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewError(InternalServerErrorCode))
		return
	}
	articles, getArticlesError := backend.Controller.LatestNews(page)
	if getArticlesError != nil {
		log.Print(getArticlesError)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewError(InternalServerErrorCode))
		return
	}
	ctx.JSON(http.StatusOK, NewsResponse{
		CurrentPage:     page,
		NumberOfResults: len(articles),
		News:            articles,
	})
}

func (backend *Backend) SearchNews(ctx *gin.Context) {

}
