package config

import "github.com/kkyr/fig"

type Config interface {
	Databaser
	Listener
	Limiter
}

type config struct {
	Databaser
	Listener
	Limiter
}

func New(file fig.Option, lim *int16) Config {
	return &config{
		Databaser: NewDatabaser(file),
		Listener:  NewListener(file),
		Limiter:   NewLimiter(lim),
	}
}
