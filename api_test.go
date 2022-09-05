package routific_test

import (
	"encoding/json"
	"os"
	"testing"

	r "github.com/slamethendry/routific"
	"github.com/stretchr/testify/assert"
)

// api_test calls Routific server API and compare the result
// between input from JSON and input from Routific object.
// Prerequisite: Auth Bearer Token, which, for testing purpose
// only, is assumed to be available as environment variable
// "Routific_Token".
// Test data is defined in setup_test.

var token = os.Getenv("Routific_Token")

func TestVRP(t *testing.T) {

	assert.NotEmpty(t, token)

	// Compare the JSON conversion vs internally created object
	var v r.VRPlan
	err := json.Unmarshal([]byte(vrpInputJSON), &v)
	assert.Nil(t, err)
	assert.Equal(t, v, vrpInput)

	// VRP call using JSON
	output1, err := r.VRP(v, token)
	assert.Nil(t, err)
	assert.NotEmpty(t, output1)
	assert.Equal(t, output1.Status, "success")

	// VRP call using internally created object
	output2, err := r.VRP(vrpInput, token)
	assert.Nil(t, err)
	assert.NotEmpty(t, output2)

	assert.Equal(t, output1, output2)
}

func TestPDP(t *testing.T) {

	assert.NotEmpty(t, token)

	// Compare the JSON conversion vs internally created object
	var p r.PDPlan
	err := json.Unmarshal([]byte(pdpInputJSON), &p)
	assert.Nil(t, err)
	assert.Equal(t, p, pdpInput)

	// PDP call using JSON
	output1, err := r.PDP(p, token)
	assert.Nil(t, err)
	assert.NotEmpty(t, output1)
	assert.Equal(t, output1.Status, "success")

	// PDP call using internally created object
	output2, err := r.PDP(pdpInput, token)
	assert.Nil(t, err)
	assert.NotEmpty(t, output2)

	assert.Equal(t, output1, output2)
}

func TestLongVRP(t *testing.T) {

	assert.NotEmpty(t, token)

	vrp, err := r.LongVRP(vrpInput, token, 3, 5)
	assert.Nil(t, err)
	assert.Equal(t, "success", vrp.Status)
	assert.NotEmpty(t, vrp.Solution)
}

func TestLongPDP(t *testing.T) {

	assert.NotEmpty(t, token)

	pdp, err := r.LongPDP(pdpInput, token, 3, 2)
	assert.Nil(t, err)
	assert.Equal(t, "success", pdp.Status)
	assert.NotEmpty(t, pdp.Solution)
}
