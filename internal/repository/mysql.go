package repository

import "context"

//MySQLRepository is interface
type MySQLRepository interface {
	InsertLogging(context.Context, map[string]interface{}) error
}
