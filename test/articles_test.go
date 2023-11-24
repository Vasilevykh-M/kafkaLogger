//go:build integration
// +build integration

package tests

import (
	"context"
	"github.com/stretchr/testify/assert"
	"homework-6/internal/serv/core"
	"homework-6/internal/serv/repository/postgres"
	"homework-6/internal/serv/server"
	"homework-6/test/fixtures"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateArticle(t *testing.T) {
	var (
		ctx = context.Background()
	)
	t.Run("success", func(t *testing.T) {

		db.SetUp(t)
		defer db.TearDown()
		//arrange
		repo := postgres.NewArticles(db.DB)
		serv := core.FacadeServer{&server.Server{Repo: repo}, answerService}

		//act
		err, resp := serv.Server.CreateArticle(ctx, fixtures.Article().Valid().P())

		//assert

		require.Equal(t, http.StatusOK, err)
		assert.NotZero(t, resp)

	})
}

func TestCreatePost(t *testing.T) {
	var (
		ctx = context.Background()
	)
	t.Run("success", func(t *testing.T) {
		db.SetUp(t)
		defer db.TearDown()
		//arrange
		repo := postgres.NewArticles(db.DB)
		serv := core.FacadeServer{&server.Server{Repo: repo}, answerService}

		//act

		err, respA := serv.Server.CreateArticle(ctx, fixtures.Article().Valid().P())

		require.Equal(t, http.StatusOK, err)
		assert.NotZero(t, respA)

		err, resp := serv.Server.CreatePost(ctx, 1, fixtures.Post().Valid().P())

		//assert

		require.Equal(t, http.StatusOK, err)
		assert.NotZero(t, resp)

	})
}

func TestGetByID(t *testing.T) {
	var (
		ctx = context.Background()
	)
	t.Run("success", func(t *testing.T) {

		db.SetUp(t)
		defer db.TearDown()
		//arrange
		repo := postgres.NewArticles(db.DB)
		serv := core.FacadeServer{&server.Server{Repo: repo}, answerService}

		//act

		err, respA := serv.Server.CreateArticle(ctx, fixtures.Article().Valid().P())

		require.Equal(t, http.StatusOK, err)
		assert.NotZero(t, respA)

		err, respJAP := serv.Server.Get(ctx, 1)

		//assert
		require.Equal(t, http.StatusOK, err)
		assert.NotZero(t, respJAP)
	})
}

func TestDeleteArticle(t *testing.T) {
	var (
		ctx = context.Background()
	)
	t.Run("success", func(t *testing.T) {

		db.SetUp(t)
		defer db.TearDown()
		//arrange
		repo := postgres.NewArticles(db.DB)
		serv := core.FacadeServer{&server.Server{Repo: repo}, answerService}

		//act

		err, respA := serv.Server.CreateArticle(ctx, fixtures.Article().Valid().P())

		require.Equal(t, http.StatusOK, err)
		assert.NotZero(t, respA)

		err = serv.Server.DeleteArticle(ctx, 1)

		//assert
		require.Equal(t, http.StatusOK, err)
	})
}

func TestUpdateArticle(t *testing.T) {
	var (
		ctx = context.Background()
	)
	t.Run("success", func(t *testing.T) {
		db.SetUp(t)
		defer db.TearDown()
		//arrange
		repo := postgres.NewArticles(db.DB)
		serv := core.FacadeServer{&server.Server{Repo: repo}, answerService}

		//act
		err, respA := serv.Server.CreateArticle(ctx, fixtures.Article().Valid().P())

		require.Equal(t, http.StatusOK, err)
		assert.NotZero(t, respA)

		err, respA = serv.Server.UpdateArticle(ctx, fixtures.Article().Valid().P())

		//assert
		require.Equal(t, http.StatusOK, err)
	})
}
