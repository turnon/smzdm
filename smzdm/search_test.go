package smzdm

import "testing"

func TestQuery(t *testing.T) {
	s := Query("macbook pro")

	if len(s.Entries) == 0 {
		t.Error("Nothing responsed")
	}

	for _, e := range s.Entries {
		t.Log(e.Time + ": " + e.Title)
	}
}
