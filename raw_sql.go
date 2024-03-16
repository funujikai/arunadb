package arunadb

import "database/sql"

type RawSQL struct {
	db        *DB
	sql       string
	arguments []interface{}
}

func (db *DB) RawSQL(sql string, args ...interface{}) *RawSQL {
	return &RawSQL{
		db:        db,
		sql:       sql,
		arguments: args,
	}
}

func (raw *RawSQL) Do(record interface{}) error {
	recordInfo, err := buildRecordDescription(record)
	if err != nil {
		return err
	}

	// the function which will return the pointers according to the given columns
	pointersGetter := func(record interface{}, columns []string) ([]interface{}, error) {
		var pointers []interface{}
		pointers, err := recordInfo.structMapping.GetPointersForColumns(record, columns...)
		return pointers, err
	}

	rowsCount, err := raw.db.doSelectOrWithReturning(raw.sql, raw.arguments, recordInfo, pointersGetter)
	if err != nil {
		return err
	}

	// When a single instance is requested but not found, sql.ErrNoRows is
	// returned like QueryRow in database/sql package.
	if !recordInfo.isSlice && rowsCount == 0 {
		err = sql.ErrNoRows
	}

	return err
}