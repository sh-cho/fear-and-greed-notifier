package main

import (
	"encoding/json"
	"testing"
)

func TestFgiParse(t *testing.T) {
	res := `{
  "lastUpdated": {
    "epochUnixSeconds": 1687874279,
    "humanDate": "2023-06-27T13:57:59.000Z"
  },
  "fgi": {
    "now": {
      "value": 73,
      "valueText": "Greed"
    },
    "previousClose": {
      "value": 71,
      "valueText": "Greed"
    },
    "oneWeekAgo": {
      "value": 78,
      "valueText": "Extreme Greed"
    },
    "oneMonthAgo": {
      "value": 69,
      "valueText": "Greed"
    },
    "oneYearAgo": {
      "value": 30,
      "valueText": "Fear"
    }
  }
}`
	r := FgiResult{}
	err := json.Unmarshal([]byte(res), &r)
	if err != nil {
		t.Errorf("parse failed")
	}
}
