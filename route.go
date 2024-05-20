package transdevclient

type Route struct {
	ID                 int    `json:"route_id"`
	DataOwnerCode      string `json:"dataowner_code"`
	RouteColor         string `json:"route_color"`
	RouteShortName     string `json:"route_short_name"`
	LinePlanningNumber string `json:"line_planning_number"`
	RouteType          int    `json:"route_type"`
	RouteTextColor     string `json:"route_text_color"`
	AgencyID           string `json:"agency_id"`
	RouteLongName      string `json:"route_long_name"`
	// TODO: Find implementations of the fields below so they can be assigned their correct type.
	//DisruptionAdvice ? `json:"disruption_advice"`
	//AlertsActive       ?   `json:"alerts_active"`
}
