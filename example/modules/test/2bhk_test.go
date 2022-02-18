package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	customAssert "github.com/jskswamy/terraform-providers-dwarka/example/modules/test/assert"
)

func TestTerraform2bhk(t *testing.T) {
	t.Run("should be able to create with 2bhk", func(t *testing.T) {
		terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
			TerraformDir: "../2bhk",

			Vars: map[string]interface{}{
				"building_name": "New Residences",
				"floor_name":    "first floor",
			},
		})

		defer terraform.Destroy(t, terraformOptions)

		terraform.InitAndApply(t, terraformOptions)

		customAssert.Outputs(t, terraformOptions, map[string]interface{}{
			"building_id": "new-residences",
			"floor_id":    "first-floor",
		})
	})

	t.Run("should be able to create 2bhk with default floor name", func(t *testing.T) {
		terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
			TerraformDir: "../2bhk",

			Vars: map[string]interface{}{
				"building_name": "New Residences",
			},
		})

		defer terraform.Destroy(t, terraformOptions)

		terraform.InitAndApply(t, terraformOptions)

		customAssert.Outputs(t, terraformOptions, map[string]interface{}{
			"building_id": "new-residences",
			"floor_id":    "ground-floor",
		})
	})
}
