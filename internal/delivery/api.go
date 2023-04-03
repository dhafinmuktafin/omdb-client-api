package delivery

import (
	omdbClient "github.com/dhafinmuktafin/omdb-client-api/internal/delivery/api"
)

func (c *OmdbClientDelivery) InitRoutes() {
	c.RouterWrapper.OmdbClientRouter = omdbClient.New(c.Middleware, c.OmdbUseCase, c.Config)
}
