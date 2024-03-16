package arunadb

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/samonzeweb/godb/adapters"
	"github.com/samonzeweb/godb/tablenamer"

	"context"
)

type DB struct {
	adapter      adapters.Adapter
	sqlDB        *sql.DB
	sqlTx        *sql.Tx
	logger       Logger
	consumedTime time.Duration
	// Called to format db table name if TableName() func is not defined for model struct
	defaultTableNamer tablenamer.NamerFn
	// Prepared Statement cache for DB and Tx
	stmtCacheDB *StmtCache
	stmtCacheTx *StmtCache
	// Optional error parsing by adapters (false by default = legacy mode)
	// Will probably be the default behavior in new major release.
	useErrorParser bool
}