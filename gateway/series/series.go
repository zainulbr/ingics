package series

type BLEAds struct {
	ReportType string `json:"report_type,omitempty" bson:"report_type,omitempty"`
	// TrackingId ID equal with Beacon ID
	TrackingId string `json:"trackingId,omitempty" bson:"trackingId,omitempty"`
	// Source ID equal with gateway ID
	SourceId  string `json:"sourceId,omitempty" bson:"sourceId,omitempty"`
	RSSI      int    `json:"rssi" bson:"rssi"`
	RawPacket string `json:"raw_packet,omitempty" bson:"raw_packet,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
}

type Service interface {
	Validate() error
	Parse() (*BLEAds, error)
}
