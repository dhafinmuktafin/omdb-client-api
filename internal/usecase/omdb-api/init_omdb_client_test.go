package omdbClient

import (
	"reflect"
	"testing"

	"github.com/dhafinmuktafin/omdb-client-api/internal/model/core/config"
	"github.com/dhafinmuktafin/omdb-client-api/internal/repository"
)

func TestOmdb_New(t *testing.T) {
	cfg := &config.Config{}
	repo := &repository.RepositoryTypes{}
	want := New(cfg, repo)

	type args struct {
		cfg  *config.Config
		repo *repository.RepositoryTypes
	}
	tests := []struct {
		name string
		args args
		want *OmdbClient
	}{
		{
			name: "TestOmdbNew_success",
			args: args{
				cfg:  cfg,
				repo: repo,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.cfg, tt.args.repo)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}
