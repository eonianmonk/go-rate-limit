package config

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/kkyr/fig"
)

type Databaser interface {
	DB() *sql.DB
	DBType() string
}

type DatabaseConfig struct {
	Fields DatabaseConfigFields `fig:"db"`
}
type DatabaseConfigFields struct {
	Url  string `fig:"url" default:"postgres://postgres:postgres@localhost:5432/database"`
	Type string `fig:"type" default:"postgres"`
}

type databaser struct {
	action sync.Once
	db     *sql.DB
	typ    string
	file   fig.Option
}

func NewDatabaser(file fig.Option) Databaser {
	dbser := databaser{file: file}
	return &dbser
}

func (dbser *databaser) figure() {
	dbser.action.Do(func() {
		dbcfg := DatabaseConfig{}
		err := fig.Load(&dbcfg, dbser.file)
		if err != nil {
			panic(err)
		}
		cfg := &dbcfg.Fields
		db, err := sql.Open(cfg.Type, cfg.Url)
		if err != nil {
			panic(fmt.Errorf("failed to establish database connection: %s", err.Error()))
		}
		dbser.typ = cfg.Type
		dbser.db = db
	})
}

func (dbser *databaser) DB() *sql.DB {
	dbser.figure()
	return dbser.db
}

func (dbser *databaser) DBType() string {
	dbser.figure()
	return dbser.typ
}
