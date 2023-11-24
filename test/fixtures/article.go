package fixtures

import (
	"homework-6/internal/serv/server"
	"homework-6/test/states"
	"time"
)

type ArticleBuilder struct {
	instance *server.ArticleRequest
}

func Article() *ArticleBuilder {
	return &ArticleBuilder{instance: &server.ArticleRequest{}}
}

func (b *ArticleBuilder) ID(v int64) *ArticleBuilder {
	b.instance.ID = v
	return b
}
func (b *ArticleBuilder) Name(v string) *ArticleBuilder {
	b.instance.Name = v
	return b
}

func (b *ArticleBuilder) Rating(v int64) *ArticleBuilder {
	b.instance.Rating = v
	return b
}

func (b *ArticleBuilder) CreatedAt(v time.Time) *ArticleBuilder {
	b.instance.CreatedAt = v
	return b
}

func (b *ArticleBuilder) P() *server.ArticleRequest {
	return b.instance
}

func (b *ArticleBuilder) V() server.ArticleRequest {
	return *b.instance
}

func (b *ArticleBuilder) Valid() *ArticleBuilder {
	return Article().ID(states.Article1ID).Name(states.ArticleName1).Rating(1).CreatedAt(time.Time{})
}
