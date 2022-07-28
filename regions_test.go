package idnumber

import (
	"testing"
)

func TestRegionsData(t *testing.T) {
	regions, err := parseRegions(regionsData)
	if err != nil {
		t.Fatal(err)
	}
	for _, region := range regions {
		t.Logf("region: %+v\n", region)
	}
}
