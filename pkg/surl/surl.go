package surl

import (
	"fmt"
	"starry/pkg/snowflake"
)

var base62 = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GetShortCode() string {
	id := snowflake.GenIDInt()
	code := toBase62(id, base62)
	if len(code) > 7 {
		code = code[len(code)-7:]
	}
	for i := 0; i < 7-len(code); i++ {
		code += string(base62[0])
	}

	fmt.Println("Snowflake ID in base 62:", code)
	return code
}

func toBase62(id int64, base62 string) string {
	var code string
	for id > 0 {
		code = string(base62[id%62]) + code
		id /= 62
	}
	return code
}
