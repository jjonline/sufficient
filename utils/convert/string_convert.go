package convert

import (
	"strconv"
)

// String 自定义类型以实现字符串转换
type String string

// Int string to int
func (s String) Int() int {
	if s == "" {
		return 0
	}
	i, _ := strconv.Atoi(string(s))
	return i
}

// UInt8 string to uint8
func (s String) UInt8() uint8 {
	return uint8(s.Int64())
}

// UInt32 string to uint32
func (s String) UInt32() uint32 {
	return uint32(s.Int64())
}

// UInt string to uint
func (s String) UInt() uint {
	return uint(s.Int())
}

// Int64 string to int64
func (s String) Int64() int64 {
	if s == "" {
		return 0
	}
	i, _ := strconv.ParseInt(string(s), 10, 64)
	return i
}

// UInt64 string to uint64
func (s String) UInt64() uint64 {
	if s == "" {
		return 0
	}
	i, _ := strconv.ParseInt(string(s), 10, 64)
	return uint64(i)
}
