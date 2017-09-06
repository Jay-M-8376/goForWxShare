package main

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"time"
)

func hashSha1(str string) string {
	t := sha1.New()
	t.Write([]byte(str))
	return fmt.Sprintf("%x", t.Sum(nil))
}

func getRandom(n int) string {
	const allChar string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(allChar)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
