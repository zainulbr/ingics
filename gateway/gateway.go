package gateway

import "github.com/zainulbr/ingics/gateway/series"

type Service interface {
	Validate() error
	Parse() (*series.BLEAds, error)
}
