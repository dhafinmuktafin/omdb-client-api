package repository

import (
	"reflect"
	"testing"

	"github.com/dhafinmuktafin/omdb-client-api/internal/model/core/config"
)

func TestRepository_Init(t *testing.T) {
	cfg := &config.Config{}
	want, _ := Init(cfg)

	type args struct {
		cfg *config.Config
	}
	tests := []struct {
		name           string
		args           args
		wantRepository *RepositoryTypes
		wantErr        bool
	}{
		{
			name: "TestRepository_Init",
			args: args{
				cfg: cfg,
			},
			wantRepository: want,
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Init(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.wantRepository) {
				t.Errorf("Init() got = %v, want %v", got, tt.wantRepository)
			}
		})
	}
}
