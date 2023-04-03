package config

import (
	"reflect"
	"testing"
)

func TestConfig_Init(t *testing.T) {
	cfg, _ := Init()
	tests := []struct {
		name    string
		want    *Config
		wantErr bool
	}{
		{
			name:    "test1_init_cfg",
			want:    cfg,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Init()
			if (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Init() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_ReadConfig(t *testing.T) {
	type args struct {
		config *Config
		path   string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestReadConfig-Fail",
			args: args{
				config: &Config{},
				path:   "/",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadConfig(tt.args.config); got != tt.want {
				t.Errorf("ReadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
