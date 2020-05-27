package goquery

import (
	"fmt"
	"regexp"
	"strings"
)

// wrap by back quote
func quoteField(name string) string {
	// 用空格分割，确定是否使用了别名
	ns := regexp.MustCompile(`\s+`).Split(name, -1)
	switch len(ns) {
	case 1:
		return quoteCascade(ns[0])
	case 2:
		return quoteCascade(ns[0]) + " AS " + quote(ns[1])
	case 3:
		fallthrough
	default:
		if "as" == strings.ToLower(ns[1]) {
			return quoteCascade(ns[0]) + " AS " + quote(ns[2])
		} else {
			return quoteCascade(ns[0]) + " AS " + quote(ns[1])
		}
	}

	return ""
}

func quoteFieldOrder(fo string) string {
	// 用空格分割
	fos := regexp.MustCompile(`\s+`).Split(fo, -1)
	switch len(fos) {
	case 1:
		return quoteField(fos[0])
	case 2:
		fallthrough
	default:
		if strings.ToLower(fos[1]) == "desc" {
			return quoteField(fos[0]) + " DESC"
		} else {
			return quoteField(fos[0]) + " ASC"
		}
	}
	return ""
}

// wrap by back quote
func quoteTable(name string) string {
	// 用空格分割
	ns := regexp.MustCompile(`\s+`).Split(name, -1)
	switch len(ns) {
	case 1:
		return quote(ns[0])
	case 2:
		return quote(ns[0]) + " AS " + quote(ns[1])
	case 3:
		fallthrough
	default:
		if "as" == strings.ToLower(ns[1]) {
			return quote(ns[0]) + " AS " + quote(ns[2])
		} else {
			return quote(ns[0]) + " AS " + quote(ns[1])
		}
	}
	return ""
}

func quoteCascade(str string) string {
	// 用点分割
	ns := strings.Split(str, ".") // []string
	// 包裹
	names := make([]string, len(ns))
	for i, n := range ns {
		if "*" == n {
			names[i] = n
		} else {
			names[i] = quote(n)
		}
	}
	// 点 连接
	return strings.Join(names, ".")
}

// wrap by back quote
func quote(str string) string {
	return fmt.Sprintf("`%s`", strings.Trim(str, "`"))
}
