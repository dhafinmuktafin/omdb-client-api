package api

import (
	"context"
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/types"
	"github.com/gorilla/mux"
	"github.com/prometheus/common/log"
	"net/http"
	"strconv"
)

//HealthCheck used for checking service status
func (c *OmdbClientAPI) HealthCheck(w http.ResponseWriter, r *http.Request) (resp *types.JSONResponse, err error) {

	resp = &types.JSONResponse{}
	resp.Header.Code = http.StatusOK
	resp.Data = []string{"OmdbClientAPI service is up and running."}

	return
}

//GetSearchMovieByKeyword used for checking service status
func (c *OmdbClientAPI) GetSearchMovieByKeyword(w http.ResponseWriter, r *http.Request) (resp *types.JSONResponse, err error) {
	ctx := context.Background()
	resp = &types.JSONResponse{}

	page := r.FormValue("page")
	keyword := r.FormValue("keyword")
	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		return
	}

	result, err := c.UseCase.GetMovieListByKeyword(ctx, pageInt, keyword)
	if err != nil {
		log.Errorf("[GetSearchMovieByKeyword] handler error while calling Usecase func GetMovieListByKeyword , err : %+v\n", err)
		resp.Data = err
		resp.Header.Code = http.StatusInternalServerError
		return
	}

	resp.Header.Code = http.StatusOK
	resp.Data = result

	return
}

//GetMovieDetailByMovieId used for checking service status
func (c *OmdbClientAPI) GetMovieDetailByMovieId(w http.ResponseWriter, r *http.Request) (resp *types.JSONResponse, err error) {
	ctx := context.Background()

	resp = &types.JSONResponse{}
	id := mux.Vars(r)["id"]
	result, err := c.UseCase.GetMovieByMovieId(ctx, id)
	if err != nil {
		log.Errorf("[GetMovieDetailByMovieId] handler error while calling Usecase func GetMovieByMovieId , err : %+v\n", err)
		resp.Data = err
		resp.Header.Code = http.StatusInternalServerError
		return
	}

	resp.Header.Code = http.StatusOK
	resp.Data = result

	return
}
