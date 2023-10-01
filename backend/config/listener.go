package config

import (
	"net"
	"sync"

	"github.com/kkyr/fig"
)

type Listener interface {
	Listen() net.Listener
}

type ListenerConfig struct {
	Address string `fig:"addr" default:":8000"`
}

type listener struct {
	ln     net.Listener
	action sync.Once
	file   fig.Option
}

func (l *listener) Listen() net.Listener {
	l.action.Do(func() {
		cfg := ListenerConfig{}
		err := fig.Load(cfg, l.file)
		if err != nil {
			panic(err)
		}
		l.ln, err = net.Listen("tcp", cfg.Address)
		if err != nil {
			panic(err)
		}
	})
	return l.ln
}

func NewListener(file fig.Option) Listener {
	return &listener{file: file}
}
