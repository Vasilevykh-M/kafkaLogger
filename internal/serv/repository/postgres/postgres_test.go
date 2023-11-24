package postgres

import (
	"github.com/golang/mock/gomock"
	mock_database "homework-6/internal/serv/db/mocks"
	"homework-6/internal/serv/repository"
	"testing"
)

type articlesRepoFixture struct {
	ctrl   *gomock.Controller
	repo   repository.IArticleRepo
	mockDb *mock_database.MockDBops
}

func setUp(t *testing.T) articlesRepoFixture {
	ctrl := gomock.NewController(t)
	mockDb := mock_database.NewMockDBops(ctrl)
	repo := NewArticles(mockDb)
	return articlesRepoFixture{
		ctrl:   ctrl,
		repo:   repo,
		mockDb: mockDb,
	}
}

func (a *articlesRepoFixture) tearDown() {
	a.ctrl.Finish()
}
