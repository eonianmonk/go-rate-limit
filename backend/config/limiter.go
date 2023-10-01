package config

type Limiter interface {
	Limit() *int16
}

type limiter struct {
	limit *int16
}

func (l *limiter) Limit() *int16 {
	return l.limit
}

func NewLimiter(limit *int16) Limiter {
	return &limiter{limit}
}
