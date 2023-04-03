package mysql

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/dhafinmuktafin/omdb-client-api/internal/model/types"
	"github.com/jmoiron/sqlx"
)

type MySQL struct {
	driverName string
	cfg        *types.DatabaseConfig
	mutexDBMap *sync.Mutex
	dbMap      map[string]*sqlx.DB
	queries    *queries
}

func New(cfg *types.DatabaseConfig) (*MySQL, error) {
	var ctx = context.Background()

	var err error
	mysql := &MySQL{
		driverName: "mysql",
		cfg:        cfg,
		mutexDBMap: &sync.Mutex{},
		dbMap:      make(map[string]*sqlx.DB),
	}

	master, err := mysql.GetConnection(ctx, "Master")
	if err != nil {
		return nil, err
	}
	if master == nil {
		return nil, errors.New("[mysql][Init] Failed connecting to master")
	}

	slave, err := mysql.GetConnection(ctx, "Slave")
	if err != nil {
		return nil, err
	}
	if slave == nil {
		return nil, errors.New("[mysql][Init] Failed connecting to Slave")
	}

	if master != nil && slave != nil {
		mysql.queries, err = newQueries(ctx, master, slave)
		if err != nil {
			return nil, err
		}
	}

	return mysql, nil
}

func (mysql *MySQL) GetConnection(ctx context.Context, connection string) (*sqlx.DB, error) {

	var err error
	mysql.mutexDBMap.Lock()
	db := mysql.dbMap[connection]
	if db != nil {
		err = db.Ping()
		if err == nil {
			mysql.mutexDBMap.Unlock()
			return db, err
		}
	}
	mysql.mutexDBMap.Unlock()

	connName := fmt.Sprint(reflect.ValueOf(*mysql.cfg).FieldByName(connection))
	if connName == "" {
		return db, errors.New("[mysql][GetConnection] no dbMap conn for " + connection)
	}

	db, err = sqlx.Connect(mysql.driverName, connName)
	if err != nil {
		return db, err
	}

	db.SetMaxIdleConns(mysql.cfg.MaxIdleConn)
	db.SetMaxOpenConns(mysql.cfg.MaxOpenConn)
	db.SetConnMaxLifetime(time.Second * mysql.cfg.ConnMaxLifetimeSeconds)

	mysql.mutexDBMap.Lock()
	mysql.dbMap[connection] = db
	mysql.mutexDBMap.Unlock()

	return db, err
}
