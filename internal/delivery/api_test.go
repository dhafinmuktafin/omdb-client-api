package delivery

import (
	"testing"

	mockUsecase "github.com/dhafinmuktafin/omdb-client-api/testfile/usecase"

	"github.com/dhafinmuktafin/omdb-client-api/internal/delivery/util"
	httpUtil "github.com/dhafinmuktafin/omdb-client-api/internal/delivery/util"
	"github.com/dhafinmuktafin/omdb-client-api/internal/usecase"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

func TestDelivery_InitRoutes(t *testing.T) {
	router := mux.NewRouter()
	mWare := util.NewMiddleware(router, 10)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockCase := mockUsecase.NewMockOmdbClientUseCase(mockCtrl)

	type args struct {
		UseCase    usecase.OmdbClientUseCase
		Middleware *httpUtil.HttpMiddleware
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestDelivery_InitRoutes_success",
			args: args{
				UseCase:    mockCase,
				Middleware: mWare,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &OmdbClientDelivery{
				OmdbUseCase: tt.args.UseCase,
				Middleware:  tt.args.Middleware,
			}
			c.InitRoutes()
		})
	}
}
