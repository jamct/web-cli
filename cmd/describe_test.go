package cmd

import "testing"

var dates = []struct {
	in   string
	want string
}{
	{"2006-01-02T15:04:05-07:00", "02.01.2006 um 23:04"},
	{"2020-01-01T10:00:00+01:00", "01.01.2020 um 10:00"},
	{"2020-05-31T10:04:06+02:00", "31.05.2020 um 10:04"},
	{"2111-11-11T11:11:11+01:00", "11.11.2111 um 11:11"},
	//{"2020-05-31T", ""},
}

func TestFormatDate(t *testing.T) {

	for _, date := range dates {
		out := formatDate(date.in)
		if out != date.want {
			t.Errorf("Date was wrong, got: %s, want: %s.", out, date.want)
		}
	}
}
