package configuration_test

import (
	"testing"

	awsConfig "github.com/raymondsquared/common-utilities/go/cloud/aws/configuration"
)

func TestImport_InputIsValid_OutputOK(t *testing.T) {
	var importTests = []struct {
		input    awsConfig.Configuration
		expected awsConfig.Configuration
	}{
		{awsConfig.Configuration{}, awsConfig.Configuration{}},
	}

	for _, it := range importTests {
		actual := it.input
		if actual != it.expected {
			t.Errorf("Creat %+v: expected %+v, actual %+v", it.input, it.expected, actual)
		}
	}
}
