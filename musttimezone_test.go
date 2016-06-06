package musttimezone_test

import (
	"errors"
	"testing"
	"time"

	"github.com/mzimmerman/musttimezone"
)

type timeStruct struct {
	source string
	want   string
	err    error
}

func TestParse(t *testing.T) {
	times := []timeStruct{
		{source: "TUE MAY 31 23:59:52 CDT 2016", want: "2016-06-01 04:59:52", err: nil},
		{source: "TUE MAY 31 23:59:52 EDT 2016", want: "2016-06-01 03:59:52", err: nil},
		{source: "Wed Jan 07 10:22:15 EST 2015", want: "2015-01-07 15:22:15", err: nil},
		{source: "Wed Jan 07 10:22:15 CST 2015", want: "2015-01-07 16:22:15", err: nil},
		{source: "Wed Jan 07 10:22:15 DoesNotExist 2015", want: "2015-01-07 16:22:15", err: errors.New("parsing time \"Wed Jan 07 10:22:15 DoesNotExist 2015\" as \"Mon Jan 02 15:04:05 MST 2006\": cannot parse \"DoesNotExist 2015\" as \"MST\"")},
		{source: "We Jan 07 10:22:15 CST 2015", want: "2015-01-07 16:22:15", err: errors.New("parsing time \"We Jan 07 10:22:15 CST 2015\" as \"Mon Jan 02 15:04:05 MST 2006\": cannot parse \"We Jan 07 10:22:15 CST 2015\" as \"Mon\"")},
	}
	for _, x := range times {
		d, err := musttimezone.Parse("Mon Jan 02 15:04:05 MST 2006", x.source)
		if err != nil {
			if x.err == nil {
				t.Errorf("%#v - Expected no error but got %q", x, err)
			} else {
				if err.Error() != x.err.Error() {
					t.Errorf("%#v - Expected error %s, but got %s", x, x.err, err)
				}
			}
		} else {
			if want, got := x.want, d.In(time.UTC).Format("2006-01-02 15:04:05"); want != got {
				t.Errorf("%#v - Expected %s, but got %q", x, want, got)
			}
		}
	}
}
