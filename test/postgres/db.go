package postgres

import (
	"context"
	"fmt"
	"homework-6/config"
	"homework-6/internal/serv/db"
	"strings"
	"sync"
	"testing"
)

type TDB struct {
	DB db.DBops
	sync.Mutex
}

func NewFromEnv(conStrDB *config.ConnStructDB) *TDB {
	db, err := db.NewDB(context.Background(), conStrDB)
	if err != nil {
		panic(err)
	}
	return &TDB{DB: db}
}

func (d *TDB) SetUp(t *testing.T, args ...interface{}) {
	t.Helper()
	d.Lock()
	d.Truncate(context.Background())

}

func (d *TDB) TearDown() {
	defer d.Unlock()
	d.Truncate(context.Background())
}

func (d *TDB) Truncate(ctx context.Context) {
	var tables []string
	err := d.DB.Select(ctx, &tables, "SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE' AND table_name != 'goose_db_version'")
	if err != nil {
		panic(err)
	}
	if len(tables) == 0 {
		panic("run migration plz")
	}
	q := fmt.Sprintf("Truncate table %s RESTART IDENTITY", strings.Join(tables, ","))
	if err := d.DB.Exec(ctx, q); err != nil {
		panic(err)
	}
}
