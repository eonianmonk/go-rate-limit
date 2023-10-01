package config

import "github.com/kkyr/fig"

type Config interface {
	Databaser
	Listener
}

type config struct {
	Databaser
	Listener
}

func New(file fig.Option) Config {
	return &config{
		Databaser: NewDatabaser(file),
		Listener:  NewListener(file),
	}
}
