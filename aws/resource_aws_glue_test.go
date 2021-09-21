package aws

import "testing"

func TestAccAWSGlue_serial(t *testing.T) {
	testCases := map[string]map[string]func(t *testing.T){
		"ResourcePolicy": {
			"basic_hybrid": testAccAWSGlueResourcePolicy_basic_hybrid,
			"update":       testAccAWSGlueResourcePolicy_update,
			"disappears":   testAccAWSGlueResourcePolicy_disappears,
		},
	}
	// TODO: basic should fail if lake formation permissions have been set. We should verify that is the case.
	t.Run("basic", testAccAWSGlueResourcePolicy_basic)

	// The rest of the test cases should succeed regardless of lake formation permissions because of enable_hybrid = true
	for group, m := range testCases {
		m := m
		t.Run(group, func(t *testing.T) {
			for name, tc := range m {
				tc := tc
				t.Run(name, func(t *testing.T) {
					tc(t)
				})
			}
		})
	}
}
