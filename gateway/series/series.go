package series

type BLEAds struct {
	ReportType string `json:"report_type,omitempty" bson:"report_type,omitempty"`
	TagID      string `json:"tag_id,omitempty" bson:"tag_id,omitempty"`
	GatewayID  string `json:"gateway_id,omitempty" bson:"gateway_id,omitempty"`
	RSSI       int    `json:"rssi" bson:"rssi"`
	RawPacket  string `json:"raw_packet,omitempty" bson:"raw_packet,omitempty"`
	Timestamp  int64  `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
}

type Service interface {
	Validate() error
	Parse() (*BLEAds, error)
}
