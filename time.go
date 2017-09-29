package servicenow

import (
	"encoding/json"
	"time"
)

type SNTime struct {
	time.Time
}

func (s *SNTime) UnmarshalJSON(bs []byte) error {
	var st string
	err := json.Unmarshal(bs, &st)
	if err != nil {
		return err
	}
	s.Time, err = time.Parse("2006-01-02 15:04:05", st)
	return err
}
