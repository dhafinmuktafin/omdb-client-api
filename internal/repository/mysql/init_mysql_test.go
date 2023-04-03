package mysql

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/types"
	"github.com/go-test/deep"
	"github.com/jmoiron/sqlx"
	"reflect"
	"sync"
	"testing"
)

func TestMySQL_New(t *testing.T) {
	cfg := &types.DatabaseConfig{}
	want, _ := New(cfg)

	type args struct {
		config *types.DatabaseConfig
	}
	tests := []struct {
		name    string
		args    args
		want    *MySQL
		wantErr bool
	}{
		{
			name: "TestMySQL_New_error",
			args: args{
				config: &types.DatabaseConfig{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "TestMySQL_New_success",
			args: args{
				config: cfg,
			},
			want:    want,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.config)
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

func TestMySQL_GetConnection(t *testing.T) {
	mockMySQL, _, err := sqlmock.New()
	if err != nil {
		t.Error("failed mock db")
		return
	}
	defer mockMySQL.Close()

	sqlxDB := sqlx.NewDb(mockMySQL, "mockdb")
	mockDbMap := make(map[string]*sqlx.DB)
	mockDbMap["Master"] = sqlxDB

	type field struct {
		mysql *MySQL
	}
	type args struct {
		ctx  context.Context
		conn string
	}
	tests := []struct {
		name    string
		field   field
		args    args
		want    *sqlx.DB
		wantErr bool
	}{
		{
			name: "TestMySQL_GetConnection_success",
			field: field{
				mysql: &MySQL{
					driverName: "mysql",
					mutexDBMap: &sync.Mutex{},
					dbMap:      mockDbMap,
				},
			},
			args: args{
				ctx:  context.Background(),
				conn: "Master",
			},
			want:    sqlxDB,
			wantErr: false,
		},
		{
			name: "GetConnection_failed",
			field: field{
				mysql: &MySQL{
					driverName: "mysql",
					cfg: &types.DatabaseConfig{
						Slave:       "user=mockuser password=mockpass dbname=mockdbname host=mockhost port=5432 sslmode=disable",
						Master:      "user=mockuser password=mockpass dbname=mockdbname host=mockhost port=5432 sslmode=disable",
						MaxOpenConn: 10,
						MaxIdleConn: 200,
					},
					mutexDBMap: &sync.Mutex{},
					dbMap:      mockDbMap,
				},
			},
			args: args{
				ctx:  context.Background(),
				conn: "Master",
			},
			want:    sqlxDB,
			wantErr: true,
		},
		{
			name: "GetConnection_failed2",
			field: field{
				mysql: &MySQL{
					driverName: "mysql",
					mutexDBMap: &sync.Mutex{},
					cfg: &types.DatabaseConfig{
						Slave:       "user=mockuser password=mockpass dbname=mockdbname host=mockhost port=5432 sslmode=disable",
						Master:      "user=mockuser password=mockpass dbname=mockdbname host=mockhost port=5432 sslmode=disable",
						MaxOpenConn: 10,
						MaxIdleConn: 200,
					},
				},
			},
			args: args{
				ctx:  context.Background(),
				conn: "Master",
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.field.mysql.GetConnection(tt.args.ctx, tt.args.conn)
			if err != nil && tt.wantErr == false {
				t.Error(err)
			}

			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}
