package routific_test

import (
	"encoding/json"
	"testing"

	r "github.com/slamethendry/routific"
	"github.com/stretchr/testify/assert"
)

// types_test checks the type definitions for Routific objects against the
// Test data is defined in setup_test.

func TestParseVRPInput(t *testing.T) {

	// Convert InputJSON into a Routific request object
	var v r.VRPlan
	err := json.Unmarshal([]byte(vrpInputJSON), &v)
	assert.Nil(t, err)

	// Compare the JSON conversion vs manually created object
	assert.Equal(t, v, vrpInput)
}

func TestParseVRPOutput(t *testing.T) {

	// Convert outputJSON into a Routific request object
	var output r.Schedule
	err := json.Unmarshal([]byte(vrpOutputJSON), &output)
	assert.Nil(t, err)

	// Compare the JSON conversion vs internally created object
	assert.Equal(t, output, vrpOutput)
}

func TestParsePDPInput(t *testing.T) {

	// Convert InputJSON into a Routific request object
	var p r.PDPlan
	err := json.Unmarshal([]byte(pdpInputJSON), &p)
	assert.Nil(t, err)

	// Compare the JSON conversion vs internally created object
	assert.Equal(t, p, pdpInput)
}

func TestParsePDPOutput(t *testing.T) {

	// Convert outputJSON into a Routific request object
	var output r.Schedule
	err := json.Unmarshal([]byte(pdpOutputJSON), &output)
	assert.Nil(t, err)

	// Compare the JSON conversion vs manually created object
	assert.Equal(t, output, pdpOutput)
}

func TestParseOptionsInput(t *testing.T) {

	// Convert optionsJSON
	var options r.VRPlan
	err := json.Unmarshal([]byte(optionsJSON), &options)
	assert.Nil(t, err)

	// Compare the JSON conversion vs manually created object
	assert.Equal(t, options, optionsInput)
}
