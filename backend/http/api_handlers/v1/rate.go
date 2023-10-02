package v1

import (
	"fmt"
	"net"

	mw "github.com/eonianmonk/go-rate-limit/backend/http/middleware"
	"github.com/eonianmonk/go-rate-limit/backend/http/responses"
	"github.com/eonianmonk/go-rate-limit/data"
	"github.com/gofiber/fiber/v2"
)

func GetRate(c *fiber.Ctx) error {
	q := mw.GetLocal[*data.Queries](c, mw.DbKey)
	limit := mw.GetLocal[*int16](c, mw.LimitKey)
	ip, err := ipv4ToInt(net.ParseIP(c.IP()))
	if err != nil {
		responses.RenderErr(c, err, 500)
		return nil
	}
	params := data.HitRateParams{
		ID:   ip,
		Hits: *limit,
	}
	rate, err := q.HitRate(c.Context(), params)
	if err != nil {
		responses.RenderErr(c, err, 500)
		return nil
	}
	if rate >= *limit {
		err = fmt.Errorf("rate limit exceeded")
		responses.RenderErr(c, err, 500)
		return nil
	}
	responses.Render(c, responses.GetRate{Rate: rate}, 200)
	return nil
}

func ipv4ToInt(ip net.IP) (int32, error) {
	if len(ip) != 4 {
		if len(ip) < 4 {
			return 0, fmt.Errorf("unknown IP format")
		}
		ip = ip[len(ip)-4:]
		//return 0, fmt.Errorf("unknown IP format")
	}
	ipInt32 := int32(ip[0])<<24 | int32(ip[1])<<16 | int32(ip[2])<<8 | int32(ip[3])

	return ipInt32, nil
}
