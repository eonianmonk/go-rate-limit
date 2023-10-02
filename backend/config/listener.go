package config

import (
	"fmt"
	"net"
	"sync"

	"github.com/kkyr/fig"
)

type Listener interface {
	Listen() net.Listener
}

type ListenerConfig struct {
	Fields ListenerFields `fig:"svc"`
}
type ListenerFields struct {
	Address string `fig:"addr" default:":8000"`
}

type listener struct {
	ln     net.Listener
	action sync.Once
	file   fig.Option
}

func (l *listener) Listen() net.Listener {
	l.action.Do(func() {
		lcfg := ListenerConfig{}
		err := fig.Load(&lcfg, l.file)
		if err != nil {
			panic(fmt.Errorf("failed to read listener config: %s", err.Error()))
		}
		cfg := &lcfg.Fields
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
