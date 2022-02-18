package assert

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func Output(t *testing.T, options *terraform.Options, key string, expected interface{}) {
	buildingId := terraform.Output(t, options, key)
	assert.Equal(t, expected, buildingId)
}

func Outputs(t *testing.T, options *terraform.Options, expectations map[string]interface{}) {
	for key, value := range expectations {
		Output(t, options, key, value)
	}
}
