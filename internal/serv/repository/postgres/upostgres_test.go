package postgres

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArticles_GetByID(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
		id  = 1
	)
	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Run("not found", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockDb.EXPECT().Select(gomock.Any(), gomock.Any(), `SELECT articles.name_article, articles.rating, articles.created_at, post.id, post.name_post, post.sales FROM articles LEFT JOIN post ON articles.id = post.id_author WHERE articles.id=$1;`, gomock.Any()).Return(pgx.ErrNoRows)

			// act
			user, err := s.repo.GetByID(ctx, int64(id))
			// assert
			require.EqualError(t, err, "object not found")

			assert.Nil(t, user)
		})
		t.Run("internal error", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockDb.EXPECT().Select(gomock.Any(), gomock.Any(), `SELECT articles.name_article, articles.rating, articles.created_at, post.id, post.name_post, post.sales FROM articles LEFT JOIN post ON articles.id = post.id_author WHERE articles.id=$1;`, gomock.Any()).Return(assert.AnError)

			// act
			user, err := s.repo.GetByID(ctx, int64(id))
			// assert
			require.EqualError(t, err, "err network")

			assert.Nil(t, user)
		})
	})
}

func TestArticles_Delete(t *testing.T) {
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

		s.mockDb.EXPECT().Exec(gomock.Any(), `DELETE FROM articles WHERE id = $1;`, gomock.Any()).Return(nil)

		// act
		err := s.repo.Delete(ctx, int64(id))
		// assert

		require.NoError(t, err)
	})
	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Run("not found", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockDb.EXPECT().Exec(gomock.Any(), `DELETE FROM articles WHERE id = $1;`, gomock.Any()).Return(pgx.ErrNoRows)

			// act
			err := s.repo.Delete(ctx, int64(id))
			// assert
			require.EqualError(t, err, "object not found")
		})
		t.Run("internal error", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockDb.EXPECT().Exec(gomock.Any(), `DELETE FROM articles WHERE id = $1;`, gomock.Any()).Return(assert.AnError)

			// act
			err := s.repo.Delete(ctx, int64(id))
			// assert
			require.EqualError(t, err, "err network")
		})
	})
}
