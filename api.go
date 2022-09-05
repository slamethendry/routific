// Package routific is a wrapper for
// [Routific Engine API]: https://docs.routific.com/reference/api-reference.
// The package implements the API calls for vehicle routing problem
// (VRP), pickup-and-delivery problem (PDP), long-running VRP, and
// long-running PDP.
//
// The calls take a plan (visits, fleet, and options) and return a Schedule.
// Prerequisite: Routific API token.
package routific

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const vrpURL string = "https://api.routific.com/v1/vrp"
const pdpURL string = "https://api.routific.com/v1/pdp"
const vrpLongURL string = "https://api.routific.com/v1/vrp-long"
const pdpLongURL string = "https://api.routific.com/v1/pdp-long"
const longJobURL string = "https://api.routific.com/jobs"

// VRP is a wrapper for Routific API for vehicle routing problem solver.
func VRP(visits VRPlan, token string) (Schedule, error) {

	jsonOut, err := post(visits, vrpURL, token)
	if err != nil {
		return Schedule{}, err
	}

	var plan Schedule
	if err := json.Unmarshal(jsonOut, &plan); err != nil {
		return Schedule{}, err
	}

	return plan, nil
}

// PDP is a wrapper for Routific API for pickup-and-delivery problem solver.
func PDP(visits PDPlan, token string) (Schedule, error) {

	jsonOut, err := post(visits, pdpURL, token)
	if err != nil {
		return Schedule{}, err
	}

	var plan Schedule
	if err := json.Unmarshal(jsonOut, &plan); err != nil {
		return Schedule{}, err
	}

	return plan, nil
}

// LongVRP is a wrapper for Routific API for long-running vehicle routing
// problem solver.
// See [Interval]: https://docs.routific.com/reference/vrp-long to determine
// how many seconds to wait according to the size of the input list.
// If Routific server is not finished in (interval x maxRetry) seconds, then
// the function returns empty schedule with error message ("Timed out").
func LongVRP(
	visits VRPlan,
	token string,
	interval uint16, // seconds
	maxRetry uint8,
) (Schedule, error) {

	return longJob(visits, vrpLongURL, token, interval, maxRetry)
	//return job.Output, err
}

// LongPDP is a wrapper for Routific API for long-running pickup-and-delivery
// problem solver.
// See [Interval]: https://docs.routific.com/reference/vrp-long to determine
// how many seconds to wait according to the size of the input list.
// If Routific server is not finished in (interval x maxRetry) seconds, then
// the function returns empty schedule with error message ("Timed out").
func LongPDP(
	visits PDPlan,
	token string,
	interval uint16, // seconds
	maxRetry uint8,
) (Schedule, error) {

	return longJob(visits, pdpLongURL, token, interval, maxRetry)
	//return job.Output, err
}

func longJob(
	visits interface{},
	url string,
	token string,
	interval uint16,
	maxRetry uint8,
) (Schedule, error) {

	jobJSON, err := post(visits, url, token)
	if err != nil {
		return Schedule{}, err
	}

	var job struct {
		ID string `json:"job_id"`
	}

	if err := json.Unmarshal(jobJSON, &job); err != nil {
		return Schedule{}, err
	}

	jobURL := fmt.Sprintf("%s/%s", longJobURL, job.ID)

	response, err := get(jobURL, token)
	if err != nil {
		return Schedule{}, err
	}

	var check struct {
		Status string `json:"status"`
	}
	if err := json.Unmarshal(response, &check); err != nil {
		return Schedule{}, err
	}

	for try := uint8(0); check.Status != "finished" && try < maxRetry; try++ {
		time.Sleep(time.Duration(interval) * time.Second)
		response, err = get(jobURL, token)
		if err := json.Unmarshal(response, &check); err != nil {
			return Schedule{}, err
		}
	}

	if check.Status == "finished" {
		var plan struct {
			Status string   `json:"status"`
			ID     string   `json:"id"`
			Output Schedule `json:"output,omitempty"`
		}
		if err := json.Unmarshal(response, &plan); err != nil {
			return Schedule{}, err
		}
		return plan.Output, nil
	}

	if check.Status == "error" {
		var errMsg struct {
			Status string `json:"status"`
			Output string `json:"output,omitempty"`
		}
		if err := json.Unmarshal(response, &errMsg); err != nil {
			return Schedule{}, err
		}
		return Schedule{}, errors.New(errMsg.Output)
	}

	e := fmt.Sprintf("Timed out after %d x %d seconds", maxRetry, interval)
	return Schedule{}, errors.New(e)
}
