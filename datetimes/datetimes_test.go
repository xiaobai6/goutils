package datetimes

import (
	"testing"
)

func TestGetNowDate(t *testing.T) {
	nowDate := GetNowDate()

	t.Log(nowDate)
}
