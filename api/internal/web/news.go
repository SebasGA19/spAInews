package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

type (
	ArticleTime time.Time
	Article     struct {
		Authors      []string    `bson:"authors"`
		DateDownload ArticleTime `bson:"date_download"`
		DateModify   ArticleTime `bson:"date_modify"`
		DatePublish  ArticleTime `bson:"date_publish"`
		Description  string      `bson:"description"`
		Filename     string      `bson:"filename"`
		ImageURL     string      `bson:"image_url"`
		Language     string      `bson:"language"`
		SourceDomain string      `bson:"source_domain"`
		MainText     string      `bson:"maintext"`
		Title        string      `bson:"title"`
		URL          string      `bson:"url"`
	}
	NewsResponse struct {
		CurrentPage     int64     `json:"current-page"`
		NumberOfResults int       `json:"number-of-results"`
		News            []Article `json:"news"`
	}
)

func (at ArticleTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(at).Format("01/02/2006"))), nil
}

func (backend *Backend) LatestNews(ctx *gin.Context) {
	rawPage := ctx.Param(NewsPageVariable)
	page, parsePageError := strconv.ParseInt(rawPage, 10, 64)
	if parsePageError != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewError(InternalServerErrorCode))
		return
	}
	mongoArticles, getArticlesError := backend.Controller.LatestNews(page - 1)
	if getArticlesError != nil {
		log.Print(getArticlesError)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewError(InternalServerErrorCode))
		return
	}
	articles := make([]Article, 0, len(mongoArticles))
	for _, article := range mongoArticles {
		articles = append(articles, Article{
			Authors:      article.Authors,
			DateDownload: ArticleTime(article.DateDownload),
			DateModify:   ArticleTime(article.DateModify),
			DatePublish:  ArticleTime(article.DatePublish),
			Description:  article.Description,
			Filename:     article.Filename,
			ImageURL:     article.ImageURL,
			Language:     article.Language,
			SourceDomain: article.SourceDomain,
			MainText:     article.MainText,
			Title:        article.Title,
			URL:          article.URL,
		})
	}
	ctx.JSON(http.StatusOK, NewsResponse{
		CurrentPage:     page,
		NumberOfResults: len(articles),
		News:            articles,
	})
}

func (backend *Backend) SearchNews(ctx *gin.Context) {

}
