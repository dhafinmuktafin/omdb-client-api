package external_api

import (
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/core/config"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestEXTAPI_New(t *testing.T) {
	cfg := &config.Config{}
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	want := New(cfg, client)
	want1 := New(cfg, nil)

	type args struct {
		config *config.Config
		client *http.Client
	}
	tests := []struct {
		name string
		args args
		want *ExternalAPI
	}{
		{
			name: "TestExtAPI_New",
			args: args{
				config: cfg,
				client: client,
			},
			want: want,
		},
		{
			name: "TestExtAPI_New_incomplete",
			args: args{
				config: cfg,
				client: nil,
			},
			want: want1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.config, tt.args.client)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Init() got = %v, want %v", got, tt.want)
			}
		})
	}
}
