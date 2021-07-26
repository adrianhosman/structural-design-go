package marvel

import (
	"net/http"

	"github.com/adrianhosman/structural-design-go/config"
	"github.com/adrianhosman/structural-design-go/model"
)

type MarvelDAL interface {
	//GetCharacterByID from marvel api and return single data of character
	GetCharacterByID(id int64) (model.MarvelGetCharactersResponseResult, error)
	//GetCharacters get all character from marvel api (paginated)
	GetCharacters(param model.MarvelGetCharacterRequest) (model.MarvelGetCharactersResponseData, error)
}

type impl struct {
	client *http.Client
	cfg    *config.Config
}

func New(cfg *config.Config, httpClient *http.Client) MarvelDAL {
	return &impl{
		cfg:    cfg,
		client: httpClient,
	}
}
