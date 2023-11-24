package server

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"homework-6/internal/serv/repository"
	"net/http"
	"testing"
)

/*хотел написать тест для парсинга запроса, но не успел разобраться с тем как парсит горилла все это добро*/

func TestArticles_Get(t *testing.T) {
	t.Parallel()
	var (
		ctx   = context.Background()
		id    = 1
		posts = []*repository.JoinArticlePost{
			&repository.JoinArticlePost{},
		}
	)
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		// arrange

		s := setUp(t)
		defer s.tearDown()

		s.mockDb.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(posts, nil)

		// act
		id, res := s.repo.Get(ctx, int64(id))
		// assert
		assert.Equal(t, http.StatusOK, id)
		assert.NotNil(t, res)
	})
	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Run("not found", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockDb.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(posts, repository.ErrObjectNotFound)

			// act
			id, res := s.repo.Get(ctx, int64(id))
			// assert
			assert.Equal(t, http.StatusNotFound, id)
			assert.Nil(t, res)
		})
		t.Run("internal error", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()
			s.mockDb.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(posts, assert.AnError)

			// act
			id, res := s.repo.Get(ctx, int64(id))
			// assert
			assert.Equal(t, http.StatusInternalServerError, id)
			assert.Nil(t, res)
		})
	})
}

func TestArticles_CreateArticle(t *testing.T) {
	t.Parallel()
	var (
		ctx     = context.Background()
		id      = 1
		article = &ArticleRequest{}
	)
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		// arrange

		s := setUp(t)
		defer s.tearDown()

		s.mockDb.EXPECT().AddArticles(gomock.Any(), gomock.Any()).Return(int64(id), nil)

		// act
		codeState, res := s.repo.CreateArticle(ctx, article)
		// assert
		assert.Equal(t, http.StatusOK, codeState)
		assert.NotNil(t, res)
		assert.Equal(t, int64(id), res.ID)
	})
	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Run("internal server error", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockDb.EXPECT().AddArticles(gomock.Any(), gomock.Any()).Return(int64(id), assert.AnError)

			// act
			codeState, res := s.repo.CreateArticle(ctx, article)
			// assert
			assert.Equal(t, http.StatusInternalServerError, codeState)
			assert.Nil(t, res)
		})
	})
}

func TestArticles_CreatePost(t *testing.T) {
	t.Parallel()
	var (
		ctx   = context.Background()
		id    = 1
		posts = &PostRequest{}
	)
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		// arrange

		s := setUp(t)
		defer s.tearDown()

		s.mockDb.EXPECT().AddPost(gomock.Any(), gomock.Any()).Return(int64(id), nil)

		// act
		codeState, res := s.repo.CreatePost(ctx, int64(id), posts)
		// assert
		assert.Equal(t, http.StatusOK, codeState)
		assert.NotNil(t, res)
		assert.Equal(t, int64(id), res.ID)
	})
	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Run("not found", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockDb.EXPECT().AddPost(gomock.Any(), gomock.Any()).Return(int64(0), repository.ErrObjectNotFound)

			// act
			codeState, res := s.repo.CreatePost(ctx, int64(id), posts)
			// assert
			assert.Equal(t, http.StatusNotFound, codeState)
			assert.Nil(t, res)
		})
		t.Run("internal error", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockDb.EXPECT().AddPost(gomock.Any(), gomock.Any()).Return(int64(0), assert.AnError)

			// act
			codeState, res := s.repo.CreatePost(ctx, int64(id), posts)
			// assert
			assert.Equal(t, http.StatusInternalServerError, codeState)
			assert.Nil(t, res)
		})
	})
}

func TestArticles_DeleteArticle(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
		id  = 1
	)
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		// arrange

		s := setUp(t)
		defer s.tearDown()

		s.mockDb.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)

		// act
		codeState := s.repo.DeleteArticle(ctx, int64(id))
		// assert
		assert.Equal(t, http.StatusOK, codeState)
	})
	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Run("not found", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockDb.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(repository.ErrObjectNotFound)

			// act
			codeState := s.repo.DeleteArticle(ctx, int64(id))
			// assert
			assert.Equal(t, http.StatusNotFound, codeState)
		})
		t.Run("internal error", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockDb.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(assert.AnError)

			// act
			codeState := s.repo.DeleteArticle(ctx, int64(id))
			// assert
			assert.Equal(t, http.StatusInternalServerError, codeState)
		})
	})
}

func TestArticles_UpdateArticle(t *testing.T) {
	t.Parallel()
	var (
		ctx   = context.Background()
		id    = 1
		posts = ArticleRequest{}
	)
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		// arrange

		s := setUp(t)
		defer s.tearDown()

		s.mockDb.EXPECT().Update(gomock.Any(), gomock.Any()).Return(int64(id), nil)

		// act
		codeState, res := s.repo.UpdateArticle(ctx, &posts)
		// assert
		assert.Equal(t, http.StatusOK, codeState)
		assert.NotNil(t, res)
		assert.Equal(t, int64(id), res.ID)
	})
	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Run("not found", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockDb.EXPECT().Update(gomock.Any(), gomock.Any()).Return(int64(0), repository.ErrObjectNotFound)

			// act
			codeState, res := s.repo.UpdateArticle(ctx, &posts)
			// assert
			assert.Equal(t, http.StatusNotFound, codeState)
			assert.Nil(t, res)
		})
		t.Run("internal error", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockDb.EXPECT().Update(gomock.Any(), gomock.Any()).Return(int64(0), assert.AnError)

			// act
			codeState, res := s.repo.UpdateArticle(ctx, &posts)
			// assert
			assert.Equal(t, http.StatusInternalServerError, codeState)
			assert.Nil(t, res)
		})
	})
}
