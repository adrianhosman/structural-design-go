package usecase

import (
	"testing"

	"github.com/adrianhosman/structural-design-go/config"
	marveldal "github.com/adrianhosman/structural-design-go/dal/api/marvel"
	cachedal "github.com/adrianhosman/structural-design-go/dal/cache"
	repodal "github.com/adrianhosman/structural-design-go/dal/repo"
)

func TestNew(t *testing.T) {
	type args struct {
		cfg    *config.Config
		marvel marveldal.MarvelDAL
		cache  cachedal.CacheDAL
		repo   repodal.InvoiceDAL
	}
	tests := []struct {
		name string
		args args
		want Usecase
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			New(tt.args.cfg, tt.args.marvel, tt.args.cache, tt.args.repo)
		})
	}
}
