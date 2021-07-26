package usecase

import (
	"github.com/adrianhosman/structural-design-go/config"
	marveldal "github.com/adrianhosman/structural-design-go/dal/api/marvel"
	cachedal "github.com/adrianhosman/structural-design-go/dal/cache"
	"github.com/adrianhosman/structural-design-go/model"
)

type impl struct {
	cfg    *config.Config
	marvel marveldal.MarvelDAL
	cache  cachedal.CacheDAL
}

type Usecase interface {
	//GetCharacterByID get single character by id
	GetCharacterByID(id int64) (model.Character, error)
	//GetAllCharacterIDs get array of all character id from cache
	GetAllCharacterIDs() ([]int64, error)
	//SaveCharacters hit marvel character api multiple times and save array of character id to cache
	SaveCharacters() error
	CalculateInvoiceData(businessID string) (*model.CalculationInvoiceResponse, error)
}

func New(cfg *config.Config, marvel marveldal.MarvelDAL, cache cachedal.CacheDAL) Usecase {
	return &impl{
		cfg:    cfg,
		marvel: marvel,
		cache:  cache,
	}
}
