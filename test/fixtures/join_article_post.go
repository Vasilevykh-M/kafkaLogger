package fixtures

import (
	"homework-6/internal/serv/server"
	"homework-6/test/states"
	"time"
)

type JoinArticleBuilder struct {
	instance *server.JoinArticlePost
}

func JoinArticle() *JoinArticleBuilder {
	return &JoinArticleBuilder{instance: &server.JoinArticlePost{}}
}

func (b *JoinArticleBuilder) NameArticle(v string) *JoinArticleBuilder {
	b.instance.NameArticle = v
	return b
}

func (b *JoinArticleBuilder) Rating(v int64) *JoinArticleBuilder {
	b.instance.Rating = v
	return b
}

func (b *JoinArticleBuilder) CreatedAt(v time.Time) *JoinArticleBuilder {
	b.instance.CreatedAt = v
	return b
}

func (b *JoinArticleBuilder) IDPost(v int64) *JoinArticleBuilder {
	b.instance.IDPost = v
	return b
}

func (b *JoinArticleBuilder) NamePost(v string) *JoinArticleBuilder {
	b.instance.NamePost = v
	return b
}

func (b *JoinArticleBuilder) Sales(v int64) *JoinArticleBuilder {
	b.instance.Sales = v
	return b
}

func (b *JoinArticleBuilder) P() *server.JoinArticlePost {
	return b.instance
}

func (b *JoinArticleBuilder) V() server.JoinArticlePost {
	return *b.instance
}

func (b *JoinArticleBuilder) Valid() *JoinArticleBuilder {
	return JoinArticle().NameArticle(states.ArticleName1).Rating(1).CreatedAt(time.Now()).IDPost(1).NamePost(states.ArticleName1).Sales(1)
}
