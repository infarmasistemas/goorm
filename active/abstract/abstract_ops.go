package abstract

import (
	"database/sql"
	"github.com/infarmasistemas/go-abstract-record/active/query"
)

type AbstractOps struct {
	SqlOps query.SQLOps
}

func (d *AbstractOps) Prepare(object interface{}, objectArray interface{}, db *sql.DB, extraOptions ...interface{}) {
	d.SqlOps = *query.NewSQLOps(object, objectArray, db, extraOptions...)
}

func (d *AbstractOps) All() error {
	return d.SqlOps.Where()
}

func (d *AbstractOps) Count(values ...interface{}) (int, error) {
	return d.SqlOps.Count(values...)
}

func (d *AbstractOps) Max(value interface{}) (int, error) {
	return d.SqlOps.Max(value)
}

func (d *AbstractOps) Find(values ...interface{}) error {
	return d.SqlOps.Select(values...)
}

func (d *AbstractOps) Where(values ...interface{}) error {
	return d.SqlOps.Where(values...)
}

func (d *AbstractOps) Save(values ...interface{}) error {
	return d.SqlOps.Insert()
}

func (d *AbstractOps) Update(values ...interface{}) error {
	return d.SqlOps.Update(values...)
}

func (d *AbstractOps) Delete(values ...interface{}) error {
	return d.SqlOps.Delete()
}

func (d *AbstractOps) SQL(sqlQuery string, values...interface{}) ([]map[string]interface{}, error) {
	return d.SqlOps.SQL(sqlQuery, values...)
}

