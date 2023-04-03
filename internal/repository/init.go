package repository

import (
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/core/config"
	externalApi "github.com/dhafinmuktafin/omdb-client-api/internal/repository/external_api"
	"github.com/dhafinmuktafin/omdb-client-api/internal/repository/mysql"
	"github.com/prometheus/common/log"
	"net/http"
	"time"
)

type RepositoryTypes struct {
	MySQL  MySQLRepository
	ExtAPI ExternalAPIRepo
}

func Init(cfg *config.Config) (repoTypes *RepositoryTypes, err error) {
	dbOmdbClientConnection, err := mysql.New(&cfg.Database)
	if err != nil {
		log.Errorf("[Repository][Init] error initialize mysql %+v", err)
		return
	}

	httpPoolConnection := &http.Client{
		Timeout: cfg.HTTP.Timeout * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        cfg.HTTP.MaxIdleConns,
			MaxIdleConnsPerHost: cfg.HTTP.MaxIdleConnsPerHost,
			IdleConnTimeout:     cfg.HTTP.IdleConnTimeout * time.Second,
		},
	}

	externalAPI := externalApi.New(cfg, httpPoolConnection)

	repoTypes = &RepositoryTypes{
		MySQL:  dbOmdbClientConnection,
		ExtAPI: externalAPI,
	}
	return
}
