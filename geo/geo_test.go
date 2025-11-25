package geo

import (
	"testing"
)



func TestIP(t *testing.T) {
	ret := getIPInfo("120.244.159.47")
	t.Logf("%+v\n", ret)
}
