package mysql

import (
	"context"
	"log"
)

func (mysql *MySQL) InsertLogging(ctx context.Context, input map[string]interface{}) (err error) {

	_, err = mysql.queries.insertLogging.ExecContext(ctx, input)
	if err != nil {
		log.Printf("[InsertLogging] input: %v\n err ExecContext: %+v\n ", input, err)
	}
	return
}
