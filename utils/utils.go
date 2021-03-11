package utils

import (
	"errors"
	"github.com/xjh22222228/ip"
	"math/rand"
	"net"
	"os"
	"strings"
)

const letters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
func RandString(n uint) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int63() % int64(len(letters))]
	}
	return string(b)
}

func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

var publicIp string
func GetPublicIp() (string, error) {
	if publicIp != "" {
		return publicIp, nil
	}
	ipv4, err := ip.V4()
	if err != nil {
		return "", err
	}
	publicIp = ipv4
	return publicIp, nil
}

var internalIp string
func GetInternalIp() (string, error) {
	if internalIp != "" {
		return internalIp, nil
	}
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", errors.New("internal IP fetch failed, detail:" + err.Error())
	}
	defer conn.Close()

	addr := conn.LocalAddr().String()
	internalIp = strings.Split(addr, ":")[0]
	return internalIp, nil
}