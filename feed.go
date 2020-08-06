package gbfs

import (
	"strings"
	"time"
)

type (
	// Boolean ...
	Boolean bool
	// Date ...
	Date string
	// Time ...
	Time string
	// Timestamp ...
	Timestamp int64
	// Feed ...
	Feed interface {
		Name() string
		GetLanguage() string
		SetLanguage(string) Feed
		GetLastUpdated() Timestamp
		SetLastUpdated(Timestamp) Feed
		GetTTL() int
		SetTTL(int) Feed
		GetVersion() string
		SetVersion(string) Feed
		GetData() interface{}
		SetData(interface{}) Feed
		Expired() bool
	}
	// FeedCommon ...
	FeedCommon struct {
		Language    string      `json:"-"` // Unofficial helper parameter
		LastUpdated Timestamp   `json:"last_updated"`
		TTL         int         `json:"ttl"`
		Version     string      `json:"version,omitempty"` // (v1.1)
		Data        interface{} `json:"data"`
	}
	// RentalURIs ...
	RentalURIs struct {
		Android string `json:"android,omitempty"`
		IOS     string `json:"ios,omitempty"`
		Web     string `json:"web,omitempty"`
	}
	// RentalApps ...
	RentalApps struct {
		Android *RentalApp `json:"android,omitempty"`
		IOS     *RentalApp `json:"ios,omitempty"`
	}
	// RentalApp ...
	RentalApp struct {
		StoreURI     string `json:"store_uri,omitempty"`     // (v1.1)
		DiscoveryURI string `json:"discovery_uri,omitempty"` // (v1.1)
	}
)

const (
	DateFormat = "2006-01-02"
	TimeFormat = "15:04:05"

	V10 string = "1.0"
	V11 string = "1.1"
	V20 string = "2.0"
	V21 string = "2.1"
	V30 string = "3.0"

	FeedNameGbfs               = "gbfs"
	FeedNameGbfsVersions       = "gbfs_versions"
	FeedNameSystemInformation  = "system_information"
	FeedNameVehicleTypes       = "vehicle_types"
	FeedNameStationInformation = "station_information"
	FeedNameStationStatus      = "station_status"
	FeedNameFreeBikeStatus     = "free_bike_status"
	FeedNameSystemHours        = "system_hours"
	FeedNameSystemCalendar     = "system_calendar"
	FeedNameSystemRegions      = "system_regions"
	FeedNameSystemPricingPlans = "system_pricing_plans"
	FeedNameSystemAlerts       = "system_alerts"
	FeedNameGeofencingZones    = "geofencing_zones"
)

const (
	FormFactorBicycle = "bicycle"
	FormFactorCar     = "car"
	FormFactorMoped   = "moped"
	FormFactorOther   = "other"
	FormFactorScooter = "scooter"
)

// FormFactorAll ...
func FormFactorAll() []string {
	return []string{
		FormFactorBicycle,
		FormFactorCar,
		FormFactorMoped,
		FormFactorOther,
		FormFactorScooter,
	}
}

const (
	PropulsionTypeHuman          = "human"
	PropulsionTypeElectricAssist = "electric_assist"
	PropulsionTypeElectric       = "electric"
	PropulsionTypeCombustion     = "combustion"
)

// PropulsionTypeAll ...
func PropulsionTypeAll() []string {
	return []string{
		PropulsionTypeHuman,
		PropulsionTypeElectricAssist,
		PropulsionTypeElectric,
		PropulsionTypeCombustion,
	}
}

const (
	AlertTypeSystemClosure  = "SYSTEM_CLOSURE"
	AlertTypeStationClosure = "STATION_CLOSURE"
	AlertTypeStationMove    = "STATION_MOVE"
	AlertTypeOther          = "OTHER"
)

// AlertTypeAll ...
func AlertTypeAll() []string {
	return []string{
		AlertTypeSystemClosure,
		AlertTypeStationClosure,
		AlertTypeStationMove,
		AlertTypeOther,
	}
}

const (
	RentalMethodKey           = "KEY"
	RentalMethodCreditCard    = "CREDITCARD"
	RentalMethodPayPass       = "PAYPASS"
	RentalMethodApplePay      = "APPLEPAY"
	RentalMethodAndroidPay    = "ANDROIDPAY"
	RentalMethodTransitCard   = "TRANSITCARD"
	RentalMethodAccountNumber = "ACCOUNTNUMBER"
	RentalMethodPhone         = "PHONE"
)

// RentalMethodAll ...
func RentalMethodAll() []string {
	return []string{
		RentalMethodKey,
		RentalMethodCreditCard,
		RentalMethodPayPass,
		RentalMethodApplePay,
		RentalMethodAndroidPay,
		RentalMethodTransitCard,
		RentalMethodAccountNumber,
		RentalMethodPhone,
	}
}

const (
	UserTypeMember    = "member"
	UserTypeNonMember = "nonmember"
)

// UserTypeAll ...
func UserTypeAll() []string {
	return []string{UserTypeMember, UserTypeNonMember}
}

const (
	DayMon = "mon"
	DayTue = "tue"
	DayWed = "wed"
	DayThu = "thu"
	DayFri = "fri"
	DaySat = "sat"
	DaySun = "sun"
)

// DayAll ...
func DayAll() []string {
	return []string{DayMon, DayTue, DayWed, DayThu, DayFri, DaySat, DaySun}
}

// GeoJSONGeometry ...
type GeoJSONGeometry struct {
	Type        string      `json:"type"`
	Coordinates interface{} `json:"coordinates"`
	Properties  interface{} `json:"properties,omitempty"`
}

// GeoJSONFeature ...
type GeoJSONFeature struct {
	Type       string           `json:"type"`
	Geometry   *GeoJSONGeometry `json:"geometry"`
	Properties interface{}      `json:"properties,omitempty"`
}

// GeoJSONFeatureCollection ...
type GeoJSONFeatureCollection struct {
	Type     string            `json:"type"`
	Features []*GeoJSONFeature `json:"features"`
}

// NewGeoJSONFeatureCollection ...
func NewGeoJSONFeatureCollection(features []*GeoJSONFeature) *GeoJSONFeatureCollection {
	return &GeoJSONFeatureCollection{
		Type:     "FeatureCollection",
		Features: features,
	}
}

// NewGeoJSONFeature ...
func NewGeoJSONFeature(geometry *GeoJSONGeometry, properties interface{}) *GeoJSONFeature {
	return &GeoJSONFeature{
		Type:       "Feature",
		Geometry:   geometry,
		Properties: properties,
	}
}

// NewGeoJSONGeometryMultiPolygon ...
func NewGeoJSONGeometryMultiPolygon(coordinates interface{}, properties interface{}) *GeoJSONGeometry {
	return &GeoJSONGeometry{
		Type:        "MultiPolygon",
		Coordinates: coordinates,
		Properties:  properties,
	}
}

// FeedNameAll ...
func FeedNameAll() []string {
	return []string{
		FeedNameGbfs,
		FeedNameGbfsVersions,
		FeedNameSystemInformation,
		FeedNameVehicleTypes,
		FeedNameStationInformation,
		FeedNameStationStatus,
		FeedNameFreeBikeStatus,
		FeedNameSystemHours,
		FeedNameSystemCalendar,
		FeedNameSystemRegions,
		FeedNameSystemPricingPlans,
		FeedNameSystemAlerts,
		FeedNameGeofencingZones,
	}
}

// FeedStruct ...
func FeedStruct(name string) Feed {
	switch name {
	case FeedNameGbfs:
		return &FeedGbfs{}
	case FeedNameGbfsVersions:
		return &FeedGbfsVersions{}
	case FeedNameSystemInformation:
		return &FeedSystemInformation{}
	case FeedNameVehicleTypes:
		return &FeedVehicleTypes{}
	case FeedNameStationInformation:
		return &FeedStationInformation{}
	case FeedNameStationStatus:
		return &FeedStationStatus{}
	case FeedNameFreeBikeStatus:
		return &FeedFreeBikeStatus{}
	case FeedNameSystemHours:
		return &FeedSystemHours{}
	case FeedNameSystemCalendar:
		return &FeedSystemCalendar{}
	case FeedNameSystemRegions:
		return &FeedSystemRegions{}
	case FeedNameSystemPricingPlans:
		return &FeedSystemPricingPlans{}
	case FeedNameSystemAlerts:
		return &FeedSystemAlerts{}
	case FeedNameGeofencingZones:
		return &FeedGeofencingZones{}
	}
	return nil
}

// UnmarshalJSON ...
func (t *Boolean) UnmarshalJSON(b []byte) error {
	switch v := strings.ToLower(strings.Trim(string(b), `"`)); v {
	case "1", "true":
		*t = true
	default:
		*t = false
	}
	return nil
}

// Time ...
func (t Timestamp) Time() time.Time {
	return time.Unix(int64(t), 0)
}

// Name ...
func (s FeedCommon) Name() string {
	return ""
}

// GetLanguage ...
func (s FeedCommon) GetLanguage() string {
	return s.Language
}

// SetLanguage ...
func (s *FeedCommon) SetLanguage(l string) Feed {
	s.Language = l
	return s
}

// GetLastUpdated ...
func (s FeedCommon) GetLastUpdated() Timestamp {
	return s.LastUpdated
}

// SetLastUpdated ...
func (s *FeedCommon) SetLastUpdated(t Timestamp) Feed {
	s.LastUpdated = t
	return s
}

// GetTTL ...
func (s FeedCommon) GetTTL() int {
	return s.TTL
}

// SetTTL ...
func (s *FeedCommon) SetTTL(t int) Feed {
	s.TTL = t
	return s
}

// GetVersion ...
func (s FeedCommon) GetVersion() string {
	return s.Version
}

// SetVersion ...
func (s *FeedCommon) SetVersion(v string) Feed {
	s.Version = v
	return s
}

// GetData ...
func (s FeedCommon) GetData() interface{} {
	return s.Data
}

// SetData ...
func (s *FeedCommon) SetData(v interface{}) Feed {
	s.Data = v
	return s
}

// Expired ...
func (s FeedCommon) Expired() bool {
	if s.TTL == 0 {
		return false
	}
	return int64(s.LastUpdated)+int64(s.TTL) < time.Now().Unix()
}
