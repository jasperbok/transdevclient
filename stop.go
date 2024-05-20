package transdevclient

type Stop struct {
	ID                string      `json:"stopid"`
	Type              string      `json:"type"`
	Name              string      `json:"name"`
	Town              string      `json:"town"`
	Street            string      `json:"street"`
	Location          Geolocation `json:"location"`
	DeparturesLink    string      `json:"departures_link"`
	IsTransferStation bool        `json:"is_transfer_station"`
	Routes            []Route     `json:"routes"`
	Departures        []Departure `json:"departures"`
	// TODO: Find implementations of the field below so it can be assigned its correct type.
	//MaxExtraWaitingTimeForTransfer ?      `json:"max_extra_waiting_time_for_transfer"`
}
