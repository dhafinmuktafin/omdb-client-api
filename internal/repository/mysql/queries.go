package mysql

import (
	"context"
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/types"
	"github.com/prometheus/common/log"

	"github.com/jmoiron/sqlx"
)

type queries struct {
	insertLogging *sqlx.NamedStmt
	master, slave *sqlx.DB
}

func newQueries(ctx context.Context, master, slave *sqlx.DB) (qr *queries, err error) {
	queries := &queries{}

	queries.insertLogging, err = queries.prepareNamedStatement(ctx, master, types.QrInsertLog)
	if err != nil {
		log.Errorf("[newQueries]")
		return
	}

	queries.master = master
	queries.slave = slave
	return queries, err
}

func (q *queries) prepareNamedStatement(ctx context.Context, db *sqlx.DB, query string) (*sqlx.NamedStmt, error) {

	stmt, err := db.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, err
	}
	return stmt, err
}
