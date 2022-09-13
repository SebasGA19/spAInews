package controller

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	NumberOfResults = 9
)

type (
	Article struct {
		Title        string    `bson:"title"`
		Description  string    `bson:"description"`
		MainText     string    `bson:"maintext"`
		Authors      []string  `bson:"authors"`
		Category     string    `bson:"category"`
		DatePublish  time.Time `bson:"date_publish"`
		SourceDomain string    `bson:"source_domain"`
		URL          string    `bson:"url"`
	}
	SearchFilter struct {
		StartDate     *time.Time
		EndDate       *time.Time
		Sources       []string
		MainTextWords []string
		TitleWords    []string
		OldFirst      bool
	}
)

func (c *Controller) LatestNews(page int64) ([]Article, error) {
	skip := NumberOfResults * page
	opts := options.Find()
	opts.Skip = &skip
	opts.SetSort(bson.D{{"date_publish", -1}})
	cursor, findError := c.Mongo.Find(c.ctx,
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
	for index := 0; index < NumberOfResults && cursor.Next(c.ctx); index++ {
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
	var andConditions, orConditions []bson.M
	if searchFilter.StartDate != nil {
		andConditions = append(andConditions, bson.M{
			"date_publish": bson.M{"$gt": *searchFilter.StartDate},
		})
	}
	if searchFilter.EndDate != nil {
		andConditions = append(andConditions, bson.M{
			"date_publish": bson.M{"$lt": *searchFilter.EndDate},
		})
	}
	for _, mainTextWord := range searchFilter.MainTextWords {
		andConditions = append(andConditions,
			bson.M{"maintext": bson.M{"$regex": fmt.Sprintf(".*%s.*", mainTextWord)}})
	}
	for _, titleWord := range searchFilter.TitleWords {
		andConditions = append(andConditions,
			bson.M{"title": bson.M{"$regex": fmt.Sprintf(".*%s.*", titleWord)}})
	}
	for _, source := range searchFilter.Sources {
		orConditions = append(orConditions,
			bson.M{"source_domain": bson.M{"$regex": source}},
		)
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
	cursor, findError := c.Mongo.Find(c.ctx,
		bson.M{
			"$and": andConditions,
			"$or":  orConditions,
		},
		opts,
	)
	if findError != nil {
		return nil, findError
	}
	var (
		article  Article
		articles = make([]Article, 0, NumberOfResults)
	)
	for index := 0; index < NumberOfResults && cursor.Next(c.ctx); index++ {
		decodeError := cursor.Decode(&article)
		if decodeError != nil {
			return nil, decodeError
		}
		articles = append(articles, article)
	}
	return articles, nil
}
