package gs01s

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zainulbr/ingics/gateway/series"
)

const (
	ReportTypeGPRP     = "GPRP"
	ReportTypeSRRP     = "SRRP"
	ReportHeader       = "$"
	ReportSparator     = ","
	ReportMinimalArray = 5
)

// format pattern
// `$<report type>,<tag id>,<gateway id>,<rssi>,<raw packet content>,*<unix epoch timestamp>\r\n`
// timestamp is optional
type service struct {
	bleAds string
}

func (ss *service) Validate() error {
	v := ss.bleAds
	if !hasHeader(v) {
		return fmt.Errorf("Invalid Header")
	}
	v = strings.TrimLeft(v, ReportHeader)
	vv := strings.Split(v, ReportSparator)
	if !minimalArray(vv) {
		return fmt.Errorf("Invalid data long")
	}

	if !(vv[0] == ReportTypeGPRP || vv[0] == ReportTypeSRRP) {
		return fmt.Errorf("Invalid header")
	}

	return nil
}

func hasHeader(v string) bool {
	return strings.HasPrefix(v, ReportHeader)
}

func minimalArray(v []string) bool {
	return len(v) >= ReportMinimalArray
}

func (ss *service) Parse() (*series.BLEAds, error) {
	data := new(series.BLEAds)
	if err := ss.Validate(); err != nil {
		return nil, err
	}

	v := strings.TrimLeft(ss.bleAds, ReportHeader)
	vv := strings.Split(v, ReportSparator)
	data.ReportType = vv[0]
	data.TagID = vv[1]
	data.GatewayID = vv[2]
	rssi, err := strconv.Atoi(vv[3])
	if err != nil {
		return nil, err
	}

	data.RSSI = rssi
	data.RawPacket = vv[4]
	if len(vv) > ReportMinimalArray {
		timestamp, err := strconv.Atoi(vv[3])
		if err != nil {
			return nil, err
		}
		data.Timestamp = int64(timestamp)
	}

	return data, nil
}

func New(data string) series.Service {
	return &service{data}
}
