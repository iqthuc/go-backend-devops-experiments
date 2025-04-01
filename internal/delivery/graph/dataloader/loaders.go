package dataloader

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/iqthuc/sport-shop/internal/delivery/graph/model"
	"github.com/vikstrous/dataloadgen"
)

// nhớ add middleware vào router

type Loaders struct {
	BrandLoader    *dataloadgen.Loader[int64, *model.Brand]
	CategoryLoader *dataloadgen.Loader[int64, *model.Category]
}

func NewLoaders(db *sql.DB) *Loaders {
	br := &brandReader{db: db}
	cr := &categoryReader{db: db}
	return &Loaders{
		BrandLoader:    dataloadgen.NewLoader(br.getBrands, dataloadgen.WithWait(time.Millisecond)),
		CategoryLoader: dataloadgen.NewLoader(cr.getCategories, dataloadgen.WithWait(time.Millisecond)),
	}
}

type ctxKey string

const loadersKey = ctxKey("dataloaders")

func Middleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loaders := NewLoaders(db)
		ctx := context.WithValue(r.Context(), loadersKey, loaders)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
