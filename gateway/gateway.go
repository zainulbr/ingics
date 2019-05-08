package gateway

import (
	"github.com/zainulbr/ingics/gateway/series"
	"github.com/zainulbr/ingics/gateway/series/gs01s"
)

const (
	GatewayTypeGS01S = "gs01s"
)

func New(gtype, data string) series.Service {
	switch gtype {
	case GatewayTypeGS01S:
		return gs01s.New(data)
	}
	return nil
}
