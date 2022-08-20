package controller

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type (
	Article struct {
		Authors      []string  `bson:"authors" json:"authors"`
		DateDownload time.Time `bson:"date_download" json:"date_download"`
		DateModify   time.Time `bson:"date_modify" json:"date_modify"`
		DatePublish  time.Time `bson:"date_publish" json:"date_publish"`
		Description  string    `bson:"description" json:"description"`
		Filename     string    `bson:"filename" json:"filename"`
		ImageURL     string    `bson:"image_url" json:"image_url"`
		Language     string    `bson:"language" json:"language"`
		SourceDomain string    `bson:"source_domain" json:"source_domain"`
		MainText     string    `bson:"maintext" json:"maintext"`
		Title        string    `bson:"title" json:"title"`
		URL          string    `bson:"url" json:"url"`
	}
	LatestNewsFilter struct {
	}
)

func (c *Controller) LatestNews(page int64) ([]Article, error) {
	skip := 10 * page
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
		articles = make([]Article, 0, 10)
	)
	for index := 0; index < 10 && cursor.Next(c.ctx); index++ {
		decodeError := cursor.Decode(&article)
		if decodeError != nil {
			return nil, decodeError
		}
		articles = append(articles, article)
	}
	return articles, nil
}
