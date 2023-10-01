package config

import (
	"database/sql"
	"sync"

	"github.com/kkyr/fig"
)

type Databaser interface {
	DB() *sql.DB
	DBType() string
}

type DatabaseConfig struct {
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
		cfg := DatabaseConfig{}
		err := fig.Load(cfg, dbser.file)
		if err != nil {
			panic(err)
		}
		db, err := sql.Open(cfg.Type, cfg.Url)
		if err != nil {
			panic(err)
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
