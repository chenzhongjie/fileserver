package utils

import (
	"testing"
)

func Test_getPublicIp(t *testing.T) {
	publicIp, err := GetPublicIp()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("public ip: %s", publicIp)
}

func Test_getInternalIp(t *testing.T) {
	publicIp, err := GetInternalIp()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("internal ip: %s", publicIp)
}