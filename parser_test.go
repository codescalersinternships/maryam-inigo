package main

import "testing"
import "reflect"

func TestParse(t *testing.T) {
	t.Run("ini text", func(t *testing.T) {
		iniText := `; last modified 1 April 2001 by John Doe
[owner]
name = John Doe
organization = Acme Widgets Inc.
[database]
; use IP address in case network name resolution is not working
server = 192.0.2.62     
port = 143
file = "payroll.dat"`

		got, err := parse(iniText)
		want := map[string]map[string]string{
			"owner":    {"name": "John Doe", "organization": "Acme Widgets Inc."},
			"database": {"server": "192.0.2.62", "port": "143", "file": "\"payroll.dat\""},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v want\n %#v", got, want)
		}

	})
}