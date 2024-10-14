package config

import (
	"encoding/json"
	"testing"
)

func TestParseTemplates(t *testing.T) {
	temps, err := ParseTemplates()
	if err != nil {
		t.Error(err)
	}

	bs, err := json.Marshal(temps)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%s", string(bs))
}
