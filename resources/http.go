package resources

import (
	"net/http"
	"time"

	"github.com/adrianhosman/structural-design-go/config"
)

func InitMarvelClient(cfg *config.Config) *http.Client {
	return &http.Client{
		Timeout: time.Duration(cfg.Marvel.TimeoutInMilliSecond) * time.Millisecond,
	}
}
