package fixtures

import (
	"homework-6/internal/serv/server"
	"homework-6/test/states"
)

type PostBuilder struct {
	instance *server.PostRequest
}

func Post() *PostBuilder {
	return &PostBuilder{instance: &server.PostRequest{}}
}

func (b *PostBuilder) ID(v int64) *PostBuilder {
	b.instance.ID = v
	return b
}
func (b *PostBuilder) Name(v string) *PostBuilder {
	b.instance.Name = v
	return b
}

func (b *PostBuilder) IdAuthor(v int64) *PostBuilder {
	b.instance.IdAuthor = v
	return b
}

func (b *PostBuilder) Sales(v int64) *PostBuilder {
	b.instance.Sales = v
	return b
}

func (b *PostBuilder) P() *server.PostRequest {
	return b.instance
}

func (b *PostBuilder) V() server.PostRequest {
	return *b.instance
}

func (b *PostBuilder) Valid() *PostBuilder {
	return Post().ID(states.Article1ID).Name(states.ArticleName1).IdAuthor(1).Sales(1)
}
