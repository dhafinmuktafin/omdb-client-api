package api

import (
	"github.com/dhafinmuktafin/omdb-client-api/internal/delivery/util"
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/core/config"
	"github.com/dhafinmuktafin/omdb-client-api/internal/usecase"
	mockUsecase "github.com/dhafinmuktafin/omdb-client-api/testfile/usecase"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"reflect"
	"testing"
)

func TestAPI_New(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockCase := mockUsecase.NewMockOmdbClientUseCase(mockCtrl)
	router := mux.NewRouter()
	mWare := util.NewMiddleware(router, 10)
	cfg := &config.Config{}
	want := &OmdbClientAPI{
		Middleware: mWare,
		UseCase:    mockCase,
		Config:     cfg,
	}

	type args struct {
		middleware *util.HttpMiddleware
		usecase    usecase.OmdbClientUseCase
		cfg        *config.Config
	}
	tests := []struct {
		name string
		args args
		want *OmdbClientAPI
	}{
		{
			name: "TestAPINew_success",
			args: args{
				middleware: mWare,
				usecase:    mockCase,
				cfg:        cfg,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.middleware, tt.args.usecase, tt.args.cfg)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOmdbClientAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPI_RegisterRoutes(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockCase := mockUsecase.NewMockOmdbClientUseCase(mockCtrl)
	router := mux.NewRouter()
	mWare := util.NewMiddleware(router, 10)

	type fields struct {
		middleware *util.HttpMiddleware
		usecase    usecase.OmdbClientUseCase
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "TestAPIRegisterRoutes_success",
			fields: fields{
				middleware: mWare,
				usecase:    mockCase,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &OmdbClientAPI{
				UseCase:    tt.fields.usecase,
				Middleware: tt.fields.middleware,
			}
			d.RegisterRoutes()
		})
	}
}
