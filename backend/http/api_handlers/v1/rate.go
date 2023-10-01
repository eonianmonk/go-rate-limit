package v1

import (
	"database/sql"
	"fmt"
	"net"

	"github.com/eonianmonk/go-rate-limit/backend/http/context"
	"github.com/eonianmonk/go-rate-limit/backend/http/responses"
	"github.com/eonianmonk/go-rate-limit/data"
	"github.com/gofiber/fiber/v2"
)

func GetRate(c *fiber.Ctx) error {
	q := context.GetLocal[*data.Queries](c, context.DbKey)
	limit := context.GetLocal[*int16](c, context.LimitKey)
	ip, err := ipv4ToInt(net.ParseIP(c.IP()))
	if err != nil {
		responses.RenderErr(c, err, 500)
	}
	params := data.HitRateParams{
		ID:   ip,
		Hits: sql.NullInt16{Int16: *limit, Valid: true},
	}
	rate, err := q.HitRate(c.Context(), params)
	if err != nil {
		responses.RenderErr(c, err, 500)
	}
	responses.Render(c, responses.GetRate{Rate: rate.Hits.Int16}, 200)
	return nil
}

func ipv4ToInt(ip net.IP) (int32, error) {
	if len(ip) != 4 {
		return 0, fmt.Errorf("unknown IP format")
	}

	ipInt32 := int32(ip[0])<<24 | int32(ip[1])<<16 | int32(ip[2])<<8 | int32(ip[3])

	return ipInt32, nil
}
