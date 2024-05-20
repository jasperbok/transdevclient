package transdevclient

type Departure struct {
	ID                     string         `json:"Id"`
	ArrivalTime            UnixTimeString `json:"arrival_time"`
	DepartureTime          UnixTimeString `json:"departure_time"`
	RealtimeArrivalTime    UnixTimeInt    `json:"realtime_arrival_time"`
	RealtimeDepartureTime  UnixTimeInt    `json:"realtime_departure_time"`
	OperationDate          Date           `json:"operation_date"`
	TripID                 int            `json:"trip_id"`
	RealtimeTripID         string         `json:"realtime_trip_id"`
	StopSequence           int            `json:"stop_sequence"`
	GTFSStopID             string         `json:"gtfs_stop_id"`
	StopHeadsign           string         `json:"stop_headsign"`
	PickupType             int            `json:"pickup_type"`
	DropOffType            int            `json:"drop_off_type"`
	TimePoint              bool           `json:"timepoint"`
	StopPlaceCode          string         `json:"stop_place_code"`
	QuayCode               string         `json:"quay_code"`
	StopSideCode           string         `json:"stopsidecode"`
	QuayVisuallyAccessible bool           `json:"quay_visually_accessible"`
	QuayDisabledAccessible bool           `json:"quay_disabled_accessible"`
	Occupancy              int            `json:"occupancy"`
	ScheduleRelationship   string         `json:"schedule_relationship"`
	TripData               Trip           `json:"trip_data"`
}
