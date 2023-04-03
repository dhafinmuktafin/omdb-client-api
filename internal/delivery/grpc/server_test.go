package grpc

import (
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/core/config"
	"github.com/dhafinmuktafin/omdb-client-api/internal/usecase"
	mockUsecase "github.com/dhafinmuktafin/omdb-client-api/testfile/usecase"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockOmdbUsecase := mockUsecase.NewMockOmdbClientUseCase(mockCtrl)

	type args struct {
		o           *Option
		conf        *config.Config
		productCase usecase.OmdbClientUseCase
	}

	grpServ, _ := New(&Option{}, &config.Config{}, usecase.UseCase{
		OmdbClientUseCase: mockOmdbUsecase,
	})

	tests := []struct {
		name string
		args args
		want *GRPCServer
		mock func()
	}{
		{
			name: "TestNewServer_success",
			args: args{
				o:           &Option{},
				productCase: mockOmdbUsecase,
			},
			want: grpServ,
			mock: func() {

			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			if got := grpServ; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSignalInterrupt(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockOmdbCase := mockUsecase.NewMockOmdbClientUseCase(mockCtrl)

	type args struct {
		o           *Option
		conf        *config.Config
		productCase usecase.OmdbClientUseCase
	}

	grpServ, _ := New(&Option{}, &config.Config{}, usecase.UseCase{
		OmdbClientUseCase: mockOmdbCase,
	})

	tests := []struct {
		name string
		args args
		want *GRPCServer
		mock func()
	}{
		{
			name: "TestNewServer_success",
			args: args{
				o:           &Option{},
				conf:        &config.Config{},
				productCase: mockOmdbCase,
			},
			want: grpServ,
			mock: func() {

			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			if got := grpServ; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
