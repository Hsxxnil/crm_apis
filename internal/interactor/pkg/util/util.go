package util

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func PointerString(s string) *string     { return &s }
func PointerInt64(i int64) *int64        { return &i }
func PointerBool(b bool) *bool           { return &b }
func PointerTime(t time.Time) *time.Time { return &t }

func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max-min) + min
	return randNum
}

func Round(x float64) int64 {
	return int64(math.Floor(x + 0.5))
}

func InterfaceArrayToStringArray(arr []any) []string {
	strArr := make([]string, len(arr))
	for i, v := range arr {
		switch s := v.(type) {
		case string:
			strArr[i] = s
		default:
			strArr[i] = fmt.Sprintf("%v", s)
		}
	}
	return strArr
}
