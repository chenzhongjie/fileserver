package handler

import "testing"

func Test_getPublicIp(t *testing.T) {
	publicIp, err := getPublicIp()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("public ip: %s", publicIp)
}