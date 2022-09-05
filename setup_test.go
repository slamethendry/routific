package routific_test

import (
	r "github.com/slamethendry/routific"
)

// Common data setup for the unit test suites
// JSON data is copied from Routific doc.

const vrpInputJSON = `
{
  "visits": {
    "order_1": {
      "location": {
        "name": "6800 Cambie",
        "lat": 49.227107,
        "lng": -123.1163085
      }
    },
    "order_2": {
      "location": {
        "name": "3780 Arbutus",
        "lat": 49.2474624,
        "lng": -123.1532338
      }
    },
    "order_3": {
      "location": {
        "name": "800 Robson",
        "lat": 49.2819229,
        "lng": -123.1211844
      }
    }
  },
  "fleet": {
    "vehicle_1": {
      "start_location": {
        "id": "depot",
        "name": "800 Kingsway",
        "lat": 49.2553636,
        "lng": -123.0873365
      },
      "end_location": {
        "id": "depot",
        "name": "800 Kingsway",
        "lat": 49.2553636,
        "lng": -123.0873365
      }
    }
  }
}`

const vrpOutputJSON = `
{
  "status": "success",
  "total_travel_time": 31.983334,
  "total_idle_time": 0,
  "num_unserved": 0,
  "unserved": null,
  "solution": {
    "vehicle_1": [
      {
        "location_id": "depot",
        "location_name": "800 Kingsway"
      },
      {
        "location_id": "order_3",
        "location_name": "800 Robson"
      },
      {
        "location_id": "order_2",
        "location_name": "3780 Arbutus"
      },
      {
        "location_id": "order_1",
        "location_name": "6800 Cambie"
      },
      {
        "location_id": "depot",
        "location_name": "800 Kingsway"
      }
    ]
  }
}
`

var kingswayDepot = r.Location{
	ID:        "depot",
	Name:      "800 Kingsway",
	Latitude:  49.2553636,
	Longitude: -123.0873365,
}

var robsonDepot = r.Location{
	ID:        "depot 2",
	Name:      "800 Robson",
	Latitude:  49.2819229,
	Longitude: -123.1211844,
}

var arbutus = r.Location{
	Name:      "3780 Arbutus",
	Latitude:  49.2474624,
	Longitude: -123.1532338,
}

var robson = r.Location{
	Name:      "800 Robson",
	Latitude:  49.2819229,
	Longitude: -123.1211844,
}

var cambie = r.Location{
	Name:      "6800 Cambie",
	Latitude:  49.227107,
	Longitude: -123.1163085,
}

// Define the Routific request object, i.e. list of visits to be solved
var vrpInput = r.VRPlan{
	Visits: map[string]r.Visit{
		"order_1": {Location: cambie},
		"order_2": {Location: arbutus},
		"order_3": {Location: robson},
	},
	Fleet: map[string]r.Vehicle{
		"vehicle_1": {
			StartLocation: kingswayDepot,
			EndLocation:   kingswayDepot,
		},
	},
}

// Define the Routific solution object, i.e. ordered route
var vrpDepot = r.Stop{ID: kingswayDepot.ID, Name: kingswayDepot.Name}
var vrpStop1 = r.Stop{ID: "order_3", Name: "800 Robson"}
var vrpStop2 = r.Stop{ID: "order_2", Name: "3780 Arbutus"}
var vrpStop3 = r.Stop{ID: "order_1", Name: "6800 Cambie"}
var vrpRoute = []r.Stop{vrpDepot, vrpStop1, vrpStop2, vrpStop3, vrpDepot}
var vrpOutput = r.Schedule{
	Status:      "success",
	TravelTime:  31.983334,
	NumUnserved: 0,
	Unserved:    nil,
	Solution:    map[string]r.Stops{"vehicle_1": vrpRoute},
}

var pdpInputJSON = `
{
    "visits": {
        "order_1": {
            "load": 1,
            "pickup": {
                "location": {
                    "name": "3780 Arbutus",
                    "lat": 49.2474624,
                    "lng": -123.1532338
                },
                "start": "9:00",
                "end": "12:00",
                "duration": 10
            },
            "dropoff": {
                "location": {
                    "name": "6800 Cambie",
                    "lat": 49.227107,
                    "lng": -123.1163085
                },
                "start": "9:00",
                "end": "12:00",
                "duration": 10
            }
        },
        "order_2": {
            "load": 1,
            "pickup": {
                "location": {
                    "name": "3780 Arbutus",
                    "lat": 49.2474624,
                    "lng": -123.1532338
                },
                "start": "9:00",
                "end": "12:00",
                "duration": 10
            },
            "dropoff": {
                "location": {
                    "name": "800 Robson",
                    "lat": 49.2819229,
                    "lng": -123.1211844
                },
                "start": "9:00",
                "end": "12:00",
                "duration": 10
            }
        }
    },
    "fleet": {
        "vehicle_1": {
            "start_location": {
                "id": "depot",
                "name": "800 Kingsway",
                "lat": 49.2553636,
                "lng": -123.0873365
            },
            "end_location": {
                "id": "depot",
                "name": "800 Kingsway",
                "lat": 49.2553636,
                "lng": -123.0873365
            },
            "shift_start": "8:00",
            "shift_end": "12:00",
            "capacity": 2
        },
        "vehicle_2": {
            "start_location": {
                "id": "depot 2",
                "name": "800 Robson",
                "lat": 49.2819229,
                "lng": -123.1211844
            },
            "end_location": {
                "id": "depot",
                "name": "800 Kingsway",
                "lat": 49.2553636,
                "lng": -123.0873365
            },
            "shift_start": "8:00",
            "shift_end": "12:00",
            "capacity": 1
        }
    }
}
`

var pdpOutputJSON = `
{
  "status": "success",
  "total_travel_time": 31.283333,
  "total_idle_time": 0,
  "num_unserved": 0,
  "unserved": null,
  "solution": {
    "vehicle_1": [
      {
        "location_id": "depot",
        "location_name": "800 Kingsway",
        "arrival_time": "08:50"
      },
      {
        "location_id": "order_2",
        "location_name": "3780 Arbutus",
        "arrival_time": "09:00",
        "finish_time": "09:10",
        "type": "pickup"
      },
      {
        "location_id": "order_1",
        "location_name": "3780 Arbutus",
        "arrival_time": "09:10",
        "finish_time": "09:20",
        "type": "pickup"
      },
      {
        "location_id": "order_1",
        "location_name": "6800 Cambie",
        "arrival_time": "09:26",
        "finish_time": "09:36",
        "type": "dropoff"
      },
      {
        "location_id": "order_2",
        "location_name": "800 Robson",
        "arrival_time": "09:45",
        "finish_time": "09:55",
        "type": "dropoff"
      },
      {
        "location_id": "depot",
        "location_name": "800 Kingsway",
        "arrival_time": "10:02"
      }
    ],
    "vehicle_2": [
      {
        "location_id": "depot 2",
        "location_name": "800 Robson",
        "arrival_time": "08:00"
      },
      {
        "location_id": "depot",
        "location_name": "800 Kingsway",
        "arrival_time": "08:06"
      }
    ]
  }
}
`

var pdpInput = r.PDPlan{
	Visits: map[string]r.PickDropOrder{
		"order_1": {
			Load: 1,
			PickUp: r.Destination{
				Location: arbutus,
				Start:    "9:00",
				End:      "12:00",
				Duration: 10,
			},
			DropOff: r.Destination{
				Location: cambie,
				Start:    "9:00",
				End:      "12:00",
				Duration: 10,
			},
		},
		"order_2": {
			Load: 1,
			PickUp: r.Destination{
				Location: arbutus,
				Start:    "9:00",
				End:      "12:00",
				Duration: 10,
			},
			DropOff: r.Destination{
				Location: robson,
				Start:    "9:00",
				End:      "12:00",
				Duration: 10,
			},
		},
	},
	Fleet: map[string]r.Vehicle{
		"vehicle_1": {
			StartLocation: kingswayDepot,
			EndLocation:   kingswayDepot,
			ShiftStart:    "8:00",
			ShiftEnd:      "12:00",
			Capacity:      2,
		},
		"vehicle_2": {
			StartLocation: robsonDepot,
			EndLocation:   kingswayDepot,
			ShiftStart:    "8:00",
			ShiftEnd:      "12:00",
			Capacity:      1,
		},
	},
}

var pdpRoute1 = []r.Stop{
	{
		ID:          "depot",
		Name:        "800 Kingsway",
		ArrivalTime: "08:50",
	},
	{
		ID:          "order_2",
		Name:        "3780 Arbutus",
		ArrivalTime: "09:00",
		FinishTime:  "09:10",
		Type:        "pickup",
	},
	{
		ID:          "order_1",
		Name:        "3780 Arbutus",
		ArrivalTime: "09:10",
		FinishTime:  "09:20",
		Type:        "pickup",
	},
	{
		ID:          "order_1",
		Name:        "6800 Cambie",
		ArrivalTime: "09:26",
		FinishTime:  "09:36",
		Type:        "dropoff",
	},
	{
		ID:          "order_2",
		Name:        "800 Robson",
		ArrivalTime: "09:45",
		FinishTime:  "09:55",
		Type:        "dropoff",
	},
	{
		ID:          "depot",
		Name:        "800 Kingsway",
		ArrivalTime: "10:02",
	},
}

var pdpRoute2 = []r.Stop{
	{
		ID:          "depot 2",
		Name:        "800 Robson",
		ArrivalTime: "08:00",
	},
	{
		ID:          "depot",
		Name:        "800 Kingsway",
		ArrivalTime: "08:06",
	},
}

var pdpOutput = r.Schedule{
	Status:      "success",
	TravelTime:  31.283333,
	IdleTime:    0,
	NumUnserved: 0,
	Unserved:    nil,
	Solution: map[string]r.Stops{
		"vehicle_1": pdpRoute1,
		"vehicle_2": pdpRoute2,
	},
}

var optionsJSON = `
{
  "visits": {},
  "fleet": {},
  "options": {
    "traffic": "slow",
    "min_visits_per_vehicle": 5,
    "balance": true,
    "min_vehicles": true,
    "shortest_distance": true,
    "squash_durations": 1,
    "max_vehicle_overtime": 30,
    "max_visit_lateness": 15,
    "polylines": true
  }
}
`

var optionsInput = r.VRPlan{
	Visits: map[string]r.Visit{},
	Fleet:  map[string]r.Vehicle{},
	Options: r.Options{
		Traffic:             "slow",
		MinVisitsPerVehicle: 5,
		Balance:             true,
		MinVehicles:         true,
		ShortestDistance:    true,
		SquashDurations:     1,
		MaxVehicleOvertime:  30,
		MaxVisitLateness:    15,
		Polylines:           true,
	},
}
