package usecase

import (
	"reflect"
	"testing"

	"github.com/dhafinmuktafin/omdb-client-api/internal/model/core/config"

	repository "github.com/dhafinmuktafin/omdb-client-api/internal/repository"
)

func TestUseCase_Init(t *testing.T) {
	cfg := &config.Config{}
	repo := &repository.RepositoryTypes{}
	want := Init(cfg, repo)

	type args struct {
		cfg  *config.Config
		repo *repository.RepositoryTypes
	}
	tests := []struct {
		name        string
		args        args
		wantUseCase UseCase
	}{
		{
			name: "TestUseCaseInit_success",
			args: args{
				cfg:  cfg,
				repo: repo,
			},
			wantUseCase: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Init(tt.args.cfg, tt.args.repo)
			if !reflect.DeepEqual(got, tt.wantUseCase) {
				t.Errorf("Init() got = %v, want %v", got, tt.wantUseCase)
			}
		})
	}
}
