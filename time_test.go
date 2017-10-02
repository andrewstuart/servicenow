package servicenow

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

const tJson = `{"time": "2006-01-02 15:04:05"}`

func TestTime(t *testing.T) {
	asrt := assert.New(t)

	var tm struct{ Time SNTime }
	asrt.NoError(json.Unmarshal([]byte(tJson), &tm))

	t.Log(tm)
	asrt.Equal(2006, tm.Time.Year(), "Year should match")
	asrt.Equal(2, tm.Time.YearDay())
	asrt.Equal(15, tm.Time.Hour())
}
