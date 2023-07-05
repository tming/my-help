package tools

import (
	"fmt"
	"my-help/src/common"
	"strconv"
	"time"
)

func validTimeString(s string) (time.Time, bool) {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, s)
	if err != nil {
		return t, false
	}
	return t, true
}

func integerToTime(sec, nansec int64) string {
	t := time.Unix(sec, nansec)
	str := t.Format("2006-01-02 15:04:05")
	return str
}

func ConvertTime(s string) error {
	if s == "" {
		return common.ErrorTarget
	}

	isInteger := true
	// to check whether is an integer
	i, err := strconv.Atoi(s)
	if err != nil {
		isInteger = false
	}

	if isInteger {
		// as second
		fmt.Printf("as second: %s\n", integerToTime(int64(i), 0))
		// as ms-second
		fmt.Printf("as ms-second: %s\n", integerToTime(int64(i)/1000, int64(i)%1000*int64(time.Millisecond)))
		// as nan-second
		fmt.Printf("as nan-second: %s\n", integerToTime(0, int64(i)))
	} else {
		t, isvalid := validTimeString(s)
		if !isvalid {
			fmt.Println("Invalid time string:", s)
		} else {
			fmt.Printf("Unix: %d\n", t.Unix())
			fmt.Printf("UnixNano: %d\n", t.UnixNano())
		}
	}

	return nil
}
