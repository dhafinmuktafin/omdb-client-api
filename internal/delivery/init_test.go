package delivery

import (
	"reflect"
	"testing"

	"github.com/dhafinmuktafin/omdb-client-api/internal/model/core/config"
	"github.com/dhafinmuktafin/omdb-client-api/internal/usecase"
	"github.com/gorilla/mux"
)

func TestDelivery_Init(t *testing.T) {
	useCase := usecase.UseCase{}
	want := Init(&config.Config{}, nil, useCase)

	type args struct {
		cfg     *config.Config
		useCase usecase.UseCase
		router  *mux.Router
	}
	tests := []struct {
		name string
		args args
		want *OmdbClientDelivery
	}{
		{
			name: "TestDelivery_New_success",
			args: args{
				cfg:     &config.Config{},
				useCase: usecase.UseCase{},
				router:  nil,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Init(tt.args.cfg, tt.args.router, tt.args.useCase)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Init() = %v, want %v", got, tt.want)
			}
		})
	}
}
