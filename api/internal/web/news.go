package web

import (
	"fmt"
	"github.com/SebasGA19/spAInews/api/internal/controller"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

type (
	ArticleTime time.Time
	Article     struct {
		Title        string      `json:"title"`
		Description  string      `json:"description"`
		MainText     string      `json:"maintext"`
		Authors      []string    `json:"authors"`
		Category     string      `json:"category"`
		DatePublish  ArticleTime `json:"date_publish"`
		SourceDomain string      `json:"source_domain"`
		URL          string      `json:"url"`
	}
	NewsResponse struct {
		CurrentPage     int64     `json:"current-page"`
		NumberOfResults int       `json:"number-of-results"`
		News            []Article `json:"news"`
	}
	SearchRequest struct {
		StartDate     string   `json:"start-date"`
		EndDate       string   `json:"end-date"`
		Sources       []string `json:"sources"`
		MainTextWords []string `json:"maintext-words"`
		TitleWords    []string `json:"title-words"`
		OldFirst      bool     `json:"old-first"`
	}
)

func (at ArticleTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(at).Format("02/01/2006"))), nil
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
			Title:        article.Title,
			Description:  article.Description,
			MainText:     article.MainText,
			Authors:      article.Authors,
			Category:     article.Category,
			DatePublish:  ArticleTime(article.DatePublish),
			SourceDomain: article.SourceDomain,
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
	rawPage := ctx.Param(NewsPageVariable)
	page, parsePageError := strconv.ParseInt(rawPage, 10, 64)
	if parsePageError != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewError(InternalServerErrorCode))
		return
	}
	var searchRequest SearchRequest
	bindError := ctx.BindJSON(&searchRequest)
	if bindError != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewError(InternalServerErrorCode))
		return
	}
	searchFilter := controller.SearchFilter{
		Sources:       searchRequest.Sources,
		MainTextWords: searchRequest.MainTextWords,
		TitleWords:    searchRequest.TitleWords,
		OldFirst:      searchRequest.OldFirst,
	}
	var (
		parseError         error
		startDate, endDate time.Time
	)
	if len(searchRequest.StartDate) > 0 {
		startDate, parseError = time.Parse("02/01/2006", searchRequest.StartDate)
		if parseError != nil {
			log.Print(parseError)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewError(InternalServerErrorCode))
			return
		}
		searchFilter.StartDate = &startDate
	}
	if len(searchRequest.EndDate) > 0 {
		endDate, parseError = time.Parse("02/01/2006", searchRequest.EndDate)
		if parseError != nil {
			log.Print(parseError)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewError(InternalServerErrorCode))
			return
		}
		searchFilter.EndDate = &endDate
	}
	mongoArticles, getArticlesError := backend.Controller.SearchNews(page-1, searchFilter)
	if getArticlesError != nil {
		log.Print(getArticlesError)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, NewError(InternalServerErrorCode))
		return
	}
	articles := make([]Article, 0, len(mongoArticles))
	for _, article := range mongoArticles {
		articles = append(articles, Article{
			Title:        article.Title,
			Description:  article.Description,
			MainText:     article.MainText,
			Authors:      article.Authors,
			Category:     article.Category,
			DatePublish:  ArticleTime(article.DatePublish),
			SourceDomain: article.SourceDomain,
			URL:          article.URL,
		})
	}
	ctx.JSON(http.StatusOK, NewsResponse{
		CurrentPage:     page,
		NumberOfResults: len(articles),
		News:            articles,
	})
}
