package holidays

import (
	"fmt"
	"testing"
	"time"
)

func TestIsHoliday(t *testing.T) {
	d := time.Date(2019, 10, 1, 0, 0, 0, 0, location)
	result, err := IsHoliday(d)
	if err != nil {
		t.Error(err)
	}

	if !result {
		t.Fail()
	}

	fmt.Println(Holiday(d))

}

func TestIsWorkingday(t *testing.T) {
	d := time.Date(2019, 9, 30, 0, 0, 0, 0, location)
	result, err := IsWorkingday(d)
	if err != nil {
		t.Error(err)
	}

	if !result {
		t.Fail()
	}
}
