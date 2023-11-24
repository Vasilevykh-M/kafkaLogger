package server

import (
	"github.com/golang/mock/gomock"
	mock_repository "homework-6/internal/serv/repository/mocks"
	"testing"
)

type articlesRepoFixture struct {
	ctrl   *gomock.Controller
	repo   *Server
	mockDb *mock_repository.MockIArticleRepo
}

func setUp(t *testing.T) articlesRepoFixture {
	ctrl := gomock.NewController(t)
	mockDb := mock_repository.NewMockIArticleRepo(ctrl)
	repo := NewServer(mockDb)
	return articlesRepoFixture{
		ctrl:   ctrl,
		repo:   repo,
		mockDb: mockDb,
	}
}

func (a *articlesRepoFixture) tearDown() {
	a.ctrl.Finish()
}
