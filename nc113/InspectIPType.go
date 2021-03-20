package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * 验证IP地址
 * @param IP string字符串 一个IP地址字符串
 * @return string字符串
 */
func solve(IP string) string {
	// write code here

	if len(IP) == 0 {
		return "Neither"
	}

	if strings.Index(IP, ".") > 0 {
		parts := strings.Split(IP, ".")
		if len(parts) != 4 {
			return "Neither"
		}

		for _, part := range parts {
			if len(part) > 1 && part[0] == '0' {
				return "Neither"
			}
			n, err := strconv.ParseInt(part, 10, 0)
			if err != nil || n < 0 || n > 255 {
				return "Neither"
			}
		}

		return "IPv4"
	} else if strings.Index(IP, ":") > 0 {
		parts := strings.Split(IP, ":")
		if len(parts) != 8 {
			return "Neither"
		}

		for _, part := range parts {
			if len(part) == 0 || len(part) > 4 || part == "0000" {
				return "Neither"
			}
			n, err := strconv.ParseInt(part, 16, 0)
			if err != nil || n < 0 || n > 0xffff {
				return "Neither"
			}
		}

		return "IPv6"
	}

	return "Neither"
}

func main() {
	fmt.Println(solve(""))
	fmt.Println(solve("192.1"))
	fmt.Println(solve("192.01.2.3"))
	fmt.Println(solve("192.168.11.300"))
	fmt.Println(solve("192.168.11.abc"))
	fmt.Println(solve("192.168.11.43"))
	fmt.Println(solve("2001:0db8:85a3::8A2E:0370:7334"))
	fmt.Println(solve("02001:0db8:85a3:0000:0000:8a2e:0370:7334"))
	fmt.Println(solve("2001:0db8:85a3:0000:0000:8a2e:0370:7334"))
	fmt.Println(solve("2001:0db8:85a3:0:0:8A2E:0370:7334"))
}
