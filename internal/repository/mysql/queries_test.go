package mysql

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"reflect"
	"testing"
)

func TestMySQL_newQueries(t *testing.T) {
	ctx := context.Background()
	mockdbSlave, _, err := sqlmock.New()
	mockdbMaster, m, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}

	defer mockdbSlave.Close()
	defer mockdbMaster.Close()
	dbMaster := sqlx.NewDb(mockdbMaster, "sqlmock")
	dbSlave := sqlx.NewDb(mockdbSlave, "sqlmock")

	type args struct {
		master *sqlx.DB
		slave  *sqlx.DB
	}
	tests := []struct {
		name    string
		args    args
		want    *queries
		wantErr bool
		mock    func()
	}{
		{
			name:    "TestNewQuery_success",
			want:    nil,
			wantErr: false,
			mock: func() {
				//must called sequentially due to "append" operation inside "ExpectPrepare" function
				//sqlMock is case sensitive with query, SELECT != select
				//QrInsertCardByUserId
				m.ExpectPrepare(`INSERT INTO logging (
									request_payload, response_payload, url
								) VALUES (
									?, ?, ?
								)`)
			},
		},
		{
			name:    "TestNewQuery_error",
			want:    nil,
			wantErr: true,
			mock: func() {
				//must called sequentially due to "append" operation inside "ExpectPrepare" function
				//sqlMock is case sensitive with query, SELECT != select
				//QrInsertCardByUserId
				m.ExpectPrepare(`INSERT INTO logging (
									request_payload, response_payload, url
								) VALUES (
									?, ?, ?
								)`).WillReturnError(errors.New("ERROR"))
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		tt.want, err = newQueries(ctx, dbMaster, dbSlave)
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := newQueries(ctx, dbMaster, dbSlave)
			if (err != nil) != tt.wantErr {
				t.Errorf("newQueries() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newQueries() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMySQL_prepareNamedStatement(t *testing.T) {
	mockdb, m, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	defer mockdb.Close()

	first := "SELECT 1"
	sqlxDB := sqlx.NewDb(mockdb, "sqlmock")

	type args struct {
		ctx     context.Context
		db      *sqlx.DB
		queries string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestPrepareNamed_success",
			args: args{
				ctx:     context.Background(),
				db:      sqlxDB,
				queries: first,
			},
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.ExpectPrepare(tt.args.queries)
			q := &queries{}
			got, err := q.prepareNamedStatement(tt.args.ctx, tt.args.db, tt.args.queries)
			if err != nil {
				t.Errorf("prepareNamedStatement() error test case: %v", err)
			}
			if got == nil {
				t.Errorf("prepareNamedStatement() error test case: %d, function should not return nil", i)
			}
		})
	}
}
