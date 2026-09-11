package ttl_test

import (
	"fmt"
	"testing"

	"github.com/mantyr/i18n/humanize/ttl"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDuration(t *testing.T) {
	Convey("ttl.Duration", t, func() {
		tests := []ttl.Duration{
			{
				Days:    10,
				Hours:   20,
				Minutes: 40,
				Seconds: 50,
			},
			{
				Days:    0,
				Hours:   0,
				Minutes: 0,
				Seconds: 50,
			},
			{
				Days:    10,
				Hours:   0,
				Minutes: 0,
				Seconds: 50,
			},
		}
		for key, test := range tests {
			Convey(fmt.Sprintf("#%d - %s", key, test.Duration().String()), func() {
				in := test.Duration()
				out, err := ttl.Split(in)
				So(err, ShouldBeNil)
				So(out, ShouldResemble, test)
			})
		}
	})
}
