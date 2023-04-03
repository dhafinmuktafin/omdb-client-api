package util

import (
	"errors"
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/types"
	"github.com/eapache/go-resiliency/breaker"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"

	"reflect"
	"testing"
	"time"
)

func TestUtil_NewMiddleware(t *testing.T) {
	router := mux.NewRouter()
	want := NewMiddleware(router, 10)

	type args struct {
		router  *mux.Router
		timeout time.Duration
	}
	tests := []struct {
		name string
		args args
		want *HttpMiddleware
	}{
		{
			name: "TestNewMiddleware_success",
			args: args{
				router:  router,
				timeout: 10,
			},
			want: want,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMiddleware(tt.args.router, tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUtil_GET(t *testing.T) {
	type fields struct {
		Router         *mux.Router
		CircuitBreaker *breaker.Breaker
		Timeout        time.Duration
	}
	type args struct {
		path    string
		handler MuxHandler
	}
	handlerSuccess := func(w http.ResponseWriter, r *http.Request) (resp *types.JSONResponse, err error) {
		resp = &types.JSONResponse{
			Header: types.JSONRespHeader{
				Code: 200,
			},
		}
		return
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "TestGET_Success",
			fields: fields{
				Router:         mux.NewRouter(),
				CircuitBreaker: nil,
				Timeout:        0,
			},
			args: args{
				path:    "",
				handler: handlerSuccess,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &HttpMiddleware{
				Router:         tt.fields.Router,
				CircuitBreaker: tt.fields.CircuitBreaker,
				Timeout:        tt.fields.Timeout,
			}
			m.GET(tt.args.path, tt.args.handler)
		})
	}
}

func TestUtil_executeRequestResponse(t *testing.T) {
	type fields struct {
		Router         *mux.Router
		CircuitBreaker *breaker.Breaker
		Timeout        time.Duration
	}

	req, _ := http.NewRequest("GET", "/health", nil)
	req.Header.Set("Authorization", "aaa")

	handlerSuccess := func(w http.ResponseWriter, r *http.Request) (resp *types.JSONResponse, err error) {
		resp = &types.JSONResponse{
			Header: types.JSONRespHeader{
				Code: 200,
			},
		}
		return
	}
	handlerErr := func(w http.ResponseWriter, r *http.Request) (resp *types.JSONResponse, err error) {
		resp = &types.JSONResponse{
			Header: types.JSONRespHeader{
				Code: 400,
			},
		}
		return resp, errors.New("error response")
	}
	handlerErr2 := func(w http.ResponseWriter, r *http.Request) (resp *types.JSONResponse, err error) {
		resp = &types.JSONResponse{
			Header: types.JSONRespHeader{
				Code: 0,
			},
		}
		return resp, errors.New("error response")
	}

	type args struct {
		handler MuxHandler
		w       http.ResponseWriter
		r       *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "TestxecuteRequestResponse_success",
			fields: fields{
				Router:         mux.NewRouter(),
				CircuitBreaker: nil,
				Timeout:        0,
			},
			args: args{
				handler: handlerSuccess,
				w:       httptest.NewRecorder(),
				r:       req,
			},
		},
		{
			name: "TestxecuteRequestResponse_error",
			fields: fields{
				Router:         mux.NewRouter(),
				CircuitBreaker: nil,
				Timeout:        0,
			},
			args: args{
				handler: handlerErr,
				w:       httptest.NewRecorder(),
				r:       req,
			},
		},
		{
			name: "TestxecuteRequestResponse_error2",
			fields: fields{
				Router:         mux.NewRouter(),
				CircuitBreaker: nil,
				Timeout:        0,
			},
			args: args{
				handler: handlerErr2,
				w:       httptest.NewRecorder(),
				r:       req,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &HttpMiddleware{
				Router:         tt.fields.Router,
				CircuitBreaker: tt.fields.CircuitBreaker,
				Timeout:        tt.fields.Timeout,
			}
			m.executeRequestResponse(tt.args.handler, tt.args.w, tt.args.r)
		})
	}
}
