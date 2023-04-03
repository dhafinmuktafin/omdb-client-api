package util

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dhafinmuktafin/omdb-client-api/internal/model/types"
	"github.com/eapache/go-resiliency/breaker"
	"github.com/gorilla/mux"
)

type HttpMiddleware struct {
	Router         *mux.Router
	CircuitBreaker *breaker.Breaker
	Timeout        time.Duration
	FileBufferSize int
}

type MuxHandler func(http.ResponseWriter,
	*http.Request) (*types.JSONResponse, error)

func NewMiddleware(router *mux.Router, breakerTimeout time.Duration) *HttpMiddleware {
	return &HttpMiddleware{
		Router:  router,
		Timeout: breakerTimeout,
	}
}

func (m *HttpMiddleware) GET(path string, handler MuxHandler) {
	m.Router.Handle(path, m.muxHandler(handler)).Methods("GET")
}

func (m *HttpMiddleware) muxHandler(reqHandler MuxHandler) (handler http.Handler) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.executeRequestResponse(reqHandler, w, r)
	})
}

func (m *HttpMiddleware) executeRequestResponse(reqHandler MuxHandler, w http.ResponseWriter, r *http.Request) {
	var err error

	response := &types.JSONResponse{
		Header: types.JSONRespHeader{
			ServerProcessTime: "",
			Message:           "",
			Reason:            "",
			Code:              200,
		},
		Data:        nil,
		ErrorString: "",
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS,HEAD")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin")

	ctx := r.Context()
	ctx, cancelFunc := context.WithTimeout(ctx, m.Timeout)
	defer cancelFunc()

	r = r.WithContext(ctx)

	start := time.Now()
	response, err = reqHandler(w, r)
	if err != nil {
		response.Header.Message = err.Error()
	}

	if response.Header.Code == 0 {
		if err != nil {
			response.Header.Code = http.StatusInternalServerError
		} else {
			response.Header.Code = http.StatusOK
		}
	}

	response.Header.ServerProcessTime = time.Since(start).String()
	byteResult, err := json.Marshal(response)
	if err != nil {
		response.Header = types.JSONRespHeader{
			ServerProcessTime: time.Since(start).String(),
			Message:           "",
			Reason:            fmt.Sprintf("HTTP%d", 500),
			Code:              http.StatusInternalServerError,
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Header.Code)
	_, err = w.Write(byteResult)
	if err != nil {
		response.Header.Code = http.StatusInternalServerError
		response.Header.Reason = err.Error()
	}
}
