package parser

import "testing"
import "reflect"

func TestParser(t *testing.T) {
	t.Run("parse function", func(t *testing.T) {
		iniText := `; last modified 1 April 2001 by John Doe
[owner]
name = John Doe
organization = Acme Widgets Inc.
[database]
; use IP address in case network name resolution is not working
server = 192.0.2.62
port = 143
file = "payroll.dat"
line = `
		parser := NewParser()
		_ = parser.parse(iniText)
		got := parser.ini
		want := map[string]map[string]string{
			"owner":    {"name ": " John Doe", "organization ": " Acme Widgets Inc."},
			"database": {"server ": " 192.0.2.62", "port ": " 143", "file ": " \"payroll.dat\"", "line ": " "},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v want\n %#v", got, want)
		}

	})
	t.Run("parse function", func(t *testing.T) {
		iniText := `; last modified 1 April 2001 by John Doe
[owner]
name == John Doe
organization = Acme Widgets Inc.`
		parser := NewParser()
		_ = parser.parse(iniText)
		got := parser.ini
		want := map[string]map[string]string{
			"owner":    {"name ": " John Doe", "organization ": " Acme Widgets Inc."},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v want\n %#v", got, want)
		}

	})
	t.Run("get from string", func(t *testing.T) {
		iniText := `; last modified 1 April 2001 by John Doe
[owner]
name = John Doe
organization = Acme Widgets Inc.
[database]
; use IP address in case network name resolution is not working
server = 192.0.2.62
port = 143
file = "payroll.dat"
line = `

		parser := NewParser()
		_ = parser.LoadFromString(iniText)
		got := parser.ini
		want := map[string]map[string]string{
			"owner":    {"name ": " John Doe", "organization ": " Acme Widgets Inc."},
			"database": {"server ": " 192.0.2.62", "port ": " 143", "file ": " \"payroll.dat\"", "line ": " "},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v want\n %#v", got, want)
		}

	})
	t.Run("get from file", func(t *testing.T) {
		
		parser := NewParser()
		_ = parser.LoadFromFile("input.ini")
		got := parser.ini
		want := map[string]map[string]string{
			"owner":    {"name ": " John Doe", "organization ": " Acme Widgets Inc."},
			"database": {"server ": " 192.0.2.62", "port ": " 143", "file ": " \"payroll.dat\""},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v want\n %#v", got, want)
		}

	})
	t.Run("check correct section name", func(t *testing.T) {
		
		got := isValidSectionName("[owner]")
		want := true

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v want\n %#v", got, want)
		}

	})
	t.Run("check incorrect section name", func(t *testing.T) {
		
		got := isValidSectionName("owner]")
		want := false

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v want\n %#v", got, want)
		}

	})
	t.Run("get section names", func(t *testing.T) {
		
		parser := NewParser()
		_ = parser.LoadFromFile("input.ini")
		got := parser.GetSectionNames()
		want := []string{"owner","database"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v want\n %#v", got, want)
		}
	})
	t.Run("get values", func(t *testing.T) {
		
		parser := NewParser()
		_ = parser.LoadFromFile("input.ini")
		got, _ := parser.Get("owner","name ")
		want := " John Doe"

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v want\n %#v", got, want)
		}
	})
	t.Run("get section", func(t *testing.T) {
		
		parser := NewParser()
		_ = parser.LoadFromFile("input.ini")
		got, _ := parser.GetSection("owner")
		want := map[string]string {"name ": " John Doe", "organization ": " Acme Widgets Inc."}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v want\n %#v", got, want)
		}
	})
	t.Run("set value to existent section and key", func(t *testing.T) {
		
		parser := NewParser()
		_ = parser.LoadFromFile("input.ini")
		_ = parser.Set("owner", "name ", "Maryam Nouh")
		got := parser.ini
		want := map[string]map[string]string{
			"owner":    {"name ": "Maryam Nouh", "organization ": " Acme Widgets Inc."},
			"database": {"server ": " 192.0.2.62", "port ": " 143", "file ": " \"payroll.dat\""},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v want\n %#v", got, want)
		}
	})
	t.Run("set value to nonexistent section and key", func(t *testing.T) {
		
		parser := NewParser()
		_ = parser.LoadFromFile("input.ini")
		_ = parser.Set("customer", "age", "27")
		got := parser.ini
		want := map[string]map[string]string{
			"customer": {"age":"27"},
			"owner":    {"name ": " John Doe", "organization ": " Acme Widgets Inc."},
			"database": {"server ": " 192.0.2.62", "port ": " 143", "file ": " \"payroll.dat\""},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %#v want\n %#v", got, want)
		}
	})
}