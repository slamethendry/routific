package routific

// TimeWindow defines the time window when a location can be visited.
type TimeWindow struct {
	Start string `json:"start,omitempty"` // "hh:mm"
	End   string `json:"end,omitempty"`   // "hh:mm"
}

// Location describes the GPS coordinate of a location.
type Location struct {
	ID        string  `json:"id,omitempty"`
	Name      string  `json:"name,omitempty"`
	Latitude  float32 `json:"lat,omitempty"`
	Longitude float32 `json:"lng,omitempty"`
}

// Visit describes the targeted visit.
// See [Visits]: https://docs.routific.com/reference/input
type Visit struct {
	Location    Location     `json:"location,omitempty"`
	Start       string       `json:"start,omitempty"`    // "hh:mm"
	End         string       `json:"end,omitempty"`      // "hh:mm"
	Duration    uint8        `json:"duration,omitempty"` // minutes
	Load        interface{}  `json:"load,omitempty"`
	Type        string       `json:"type,omitempty"`
	Priority    string       `json:"priority,omitempty"`
	TimeWindows []TimeWindow `json:"time_windows,omitempty"`
	Notes       string       `json:"notes,omitempty"`
	CustomNotes interface{}  `json:"customNotes,omitempty"`
}

// Vehicle describes the vehicle or driver in the fleet.
// See [Fleet]: https://docs.routific.com/reference/fleet
type Vehicle struct {
	StartLocation Location    `json:"start_location,omitempty"`
	EndLocation   Location    `json:"end_location,omitempty"`
	ShiftStart    string      `json:"shift_start,omitempty"` // "hh:mm"
	ShiftEnd      string      `json:"shift_end,omitempty"`   // "hh:mm"
	Capacity      uint8       `json:"capacity,omitempty"`
	Type          string      `json:"type,omitempty"`
	Speed         string      `json:"speed,omitempty"`
	StrictStart   bool        `json:"strict_start,omitempty"`
	MinVisits     uint8       `json:"min_visits,omitempty"`
	Breaks        interface{} `json:"breaks,omitempty"`
}

// VRPlan is the vehicle routing plan that we want Routific to solve / optimise.
type VRPlan struct {
	Visits  map[string]Visit   `json:"visits"`
	Fleet   map[string]Vehicle `json:"fleet"`
	Options Options            `json:"options,omitempty"`
}

// Destination describes the location for pickup and dropoff.
type Destination struct {
	Location Location `json:"location"`
	Start    string   `json:"start,omitempty"`    // "hh:mm"
	End      string   `json:"end,omitempty"`      // "hh:mm"
	Duration uint8    `json:"duration,omitempty"` // minutes
}

// PickDropOrder describes the targeted pickup and dropoff.
// See [Orders]: https://docs.routific.com/reference/defining-orders
type PickDropOrder struct {
	Load    uint8       `json:"load,omitempty"`
	PickUp  Destination `json:"pickup,omitempty"`
	DropOff Destination `json:"dropoff,omitempty"`
	Type    []string    `json:"type,omitempty"`
}

// PDPlan is the pickup and dropoff plan that we want Routific to solve /
// optimise.
type PDPlan struct {
	Visits  map[string]PickDropOrder `json:"visits"`
	Fleet   map[string]Vehicle       `json:"fleet"`
	Options Options                  `json:"options,omitempty"`
}

// Stop defines the stop during the route for pickup or dropoff.
type Stop struct {
	ID          string  `json:"location_id,omitempty"`
	Name        string  `json:"location_name,omitempty"`
	ArrivalTime string  `json:"arrival_time,omitempty"` // "hh:mm"
	FinishTime  string  `json:"finish_time,omitempty"`  // "hh:mm"
	Type        string  `json:"type,omitempty"`
	Late        bool    `json:"too_late,omitempty"`
	LateBy      float32 `json:"late_by,omitempty"`
}

// Stops defines the order of stops.
type Stops []Stop

// VehicleOvertime describes how many minutes each vehicle may have overtime in
// the schedule.
type VehicleOvertime map[string]float32

// Schedule describes the Routific optimised schedule and route for the
// requested / input plan.
// See [Output]: https://docs.routific.com/reference/output
type Schedule struct {
	Status        string            `json:"status"`
	TravelTime    float32           `json:"total_travel_time"` // minutes
	IdleTime      float32           `json:"total_idle_time"`   // minutes
	Fitness       uint8             `json:"fitness,omitempty"`
	NumUnserved   uint8             `json:"num_unserved"`
	Unserved      map[string]string `json:"unserved"`
	Solution      map[string]Stops  `json:"solution"`
	NumLateVisits uint8             `json:"num_late_visits,omitempty"`
	TotalLateness float32           `json:"total_visit_lateness,omitempty"` // minutes
	Overtime      VehicleOvertime   `json:"vehicle_overtime,omitempty"`
	TotalOvertime float32           `json:"total_overtime,omitempty"` // minutes
}

// Options tweak how the Routific Engine performs the optimisation.
// See [Input Options]: https://docs.routific.com/reference/options
type Options struct {
	Traffic                 string  `json:"traffic,omitempty"`
	MinVisitsPerVehicle     uint8   `json:"min_visits_per_vehicle,omitempty"`
	Balance                 bool    `json:"balance,omitempty"`
	VisitBalanceCoefficient float32 `json:"visit_balance_coefficient,omitempty"`
	MinVehicles             bool    `json:"min_vehicles,omitempty"`
	ShortestDistance        bool    `json:"shortest_distance,omitempty"`
	SquashDurations         uint8   `json:"squash_durations,omitempty"`
	MaxVehicleOvertime      uint8   `json:"max_vehicle_overtime,omitempty"` // minutes
	MaxVisitLateness        uint8   `json:"max_visit_lateness,omitempty"`   // minutes
	Polylines               bool    `json:"polylines,omitempty"`
	AvoidTolls              bool    `json:"avoid_tolls,omitempty"`
	GeoCoder                string  `json:"geocoder,omitempty"`
}
