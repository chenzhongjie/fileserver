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
func RandomString(n uint) string {
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

func GetPublicIp() (string, error) {
	ipv4, err := ip.V4()
	if err != nil {
		return "", err
	}
	return ipv4, nil
}

func GetInternalIp() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", errors.New("internal IP fetch failed, detail:" + err.Error())
	}
	defer conn.Close()

	addr := conn.LocalAddr().String()
	return strings.Split(addr, ":")[0], nil
}