package controller

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"
)

const (
	NumberOfResults = 10
)

type (
	ArticleTime time.Time
	Article     struct {
		Title        string    `bson:"title,omitempty" json:"title,omitempty"`
		Description  string    `bson:"description,omitempty" json:"description,omitempty"`
		MainText     string    `bson:"maintext,omitempty" json:"maintext,omitempty"`
		Authors      []string  `bson:"authors,omitempty" json:"authors,omitempty"`
		Category     string    `bson:"category,omitempty" json:"category,omitempty"`
		DatePublish  time.Time `bson:"date_publish,omitempty" json:"date_publish,omitempty"`
		SourceDomain string    `bson:"source_domain,omitempty" json:"source_domain,omitempty"`
		URL          string    `bson:"url,omitempty" json:"url,omitempty"`
	}
	SearchFilter struct {
		StartDate     *time.Time `json:"start-date,omitempty"`
		EndDate       *time.Time `json:"end-date,omitempty"`
		Sources       []string   `json:"sources,omitempty"`
		MainTextWords []string   `json:"maintext-words,omitempty"`
		TitleWords    []string   `json:"title-words,omitempty"`
		OldFirst      bool       `json:"old-first,omitempty"`
	}
)

func (c *Controller) LatestNews(page int64) ([]Article, error) {
	skip := NumberOfResults * page
	opts := options.Find()
	opts.Skip = &skip
	opts.SetSort(bson.D{{"date_publish", -1}})
	cursor, findError := c.Mongo.Find(context.Background(),
		bson.D{},
		opts,
	)
	if findError != nil {
		return nil, findError
	}
	var (
		article  Article
		articles = make([]Article, 0, NumberOfResults)
	)
	for index := 0; index < NumberOfResults && cursor.Next(context.Background()); index++ {
		decodeError := cursor.Decode(&article)
		if decodeError != nil {
			return nil, decodeError
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (c *Controller) SearchNews(page int64, searchFilter SearchFilter) ([]Article, error) {
	skip := NumberOfResults * page
	// Filter
	filter := bson.M{}
	// Date Filter
	datePublishFilter := bson.M{}
	if searchFilter.StartDate != nil {
		datePublishFilter["$gt"] = *searchFilter.StartDate

	}
	if searchFilter.EndDate != nil {
		datePublishFilter["$lt"] = *searchFilter.EndDate
	}
	if len(datePublishFilter) > 0 {
		filter["date_publish"] = datePublishFilter
	}
	// Main text filter
	if len(searchFilter.MainTextWords) > 0 {
		filter["maintext"] = bson.M{
			"$regex":   fmt.Sprintf(".*%s.*", strings.Join(searchFilter.MainTextWords, ".*")),
			"$options": "i",
		}
	}
	// Title filter
	if len(searchFilter.TitleWords) > 0 {
		filter["title"] = bson.M{
			"$regex":   fmt.Sprintf(".*%s.*", strings.Join(searchFilter.TitleWords, ".*")),
			"$options": "i",
		}
	}
	// Sources filter
	if len(searchFilter.Sources) > 0 {
		filter["source_domain"] = bson.M{
			"$in": searchFilter.Sources,
		}
	}
	// Options
	opts := options.Find()
	opts.Skip = &skip
	order := -1
	if searchFilter.OldFirst {
		order = 1
	}
	opts.SetSort(bson.D{{"date_publish", order}})

	// Request
	cursor, findError := c.Mongo.Find(context.Background(),
		filter,
		opts,
	)
	if findError != nil {
		return nil, findError
	}
	var (
		article  Article
		articles = make([]Article, 0, NumberOfResults)
	)
	for index := 0; index < NumberOfResults && cursor.Next(context.Background()); index++ {
		decodeError := cursor.Decode(&article)
		if decodeError != nil {
			return nil, decodeError
		}
		articles = append(articles, article)
	}
	return articles, nil
}
