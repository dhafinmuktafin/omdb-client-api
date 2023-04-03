package mysql

import (
	"context"
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/types"
	"github.com/jmoiron/sqlx"
	"sync"
	"testing"
)

var conn MySQL

type mock struct {
	query  string
	column []string
	result [][]driver.Value
	err    error
}

func InitDBMock(dbConn string) (*sqlx.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	conn.dbMap = make(map[string]*sqlx.DB)
	conn.dbMap[dbConn] = sqlx.NewDb(db, "sqlmock")
	return conn.dbMap[dbConn], mock, nil
}

func TestMySQL_InsertLogging(t *testing.T) {
	dbConn, dbMock, err := InitDBMock("master")
	if err != nil {
		t.Errorf("Error mocking DB: %s", err.Error())
	}
	qrs := &queries{
		master: dbConn,
	}

	inputSuccess := map[string]interface{}{
		"url":             "aa",
		"requestPayload":  "aa",
		"responsePayload": "aa",
	}
	inputFail := map[string]interface{}{
		"url":             "aa",
		"requestPayload":  "aa",
		"responsePayload": 11,
	}

	type fields struct {
		driverName string
		cfg        *types.DatabaseConfig
		mtxDBMap   *sync.Mutex
		dbMap      map[string]*sqlx.DB
		queries    *queries
	}
	type args struct {
		ctx   context.Context
		input map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "TestInsertLogging: success",
			args: args{
				ctx:   context.Background(),
				input: inputSuccess,
			},
			fields: fields{
				queries: qrs,
			},
			wantErr: false,
			mock: func() {
				queryStr := `INSERT INTO (.+) VALUES (.+)`
				dbMock.ExpectPrepare(queryStr).ExpectExec().WithArgs(
					"aa", "aa", "aa",
				).WillReturnResult(sqlmock.NewResult(0, 1)).WillReturnError(nil)
				qrs.insertLogging, err = qrs.prepareNamedStatement(context.Background(), qrs.master, types.QrInsertLog)
				if err != nil {
					t.Errorf("Error mocking qrs: %s", err.Error())
				}
			},
		},
		{
			name: "TestInsertLogging: fail",
			args: args{
				ctx:   context.Background(),
				input: inputFail,
			},
			fields: fields{
				queries: qrs,
			},
			wantErr: true,
			mock: func() {
				queryStr := `INSERT INTO (.+) VALUES (.+)`
				dbMock.ExpectPrepare(queryStr).ExpectExec().WithArgs(
					"aa", "aa", 11,
				).WillReturnResult(sqlmock.NewResult(0, 1)).WillReturnError(nil)
				qrs.insertLogging, err = qrs.prepareNamedStatement(context.Background(), qrs.master, types.QrInsertLog)
				if err != nil {
					t.Errorf("Error mocking qrs: %s", err.Error())
				}
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			mysql := &MySQL{
				driverName: tt.fields.driverName,
				cfg:        tt.fields.cfg,
				mutexDBMap: tt.fields.mtxDBMap,
				dbMap:      tt.fields.dbMap,
				queries:    tt.fields.queries,
			}
			err := mysql.InsertLogging(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("MySQL.InsertLogging() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
