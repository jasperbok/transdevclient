package transdevclient

type Trip struct {
	ID                   int    `json:"trip_id"`
	RealtimeID           string `json:"realtime_trip_id"`
	Headsign             string `json:"trip_headsign"`
	RouteID              int    `json:"route_id"`
	DirectionID          string `json:"direction_id"`
	ShapeID              int    `json:"shape_id"`
	WheelchairAccessible int    `json:"wheelchair_accessible"`
	BikesAllowed         bool   `json:"bikes_allowed"`
	Route                Route  `json:"route"`
}
