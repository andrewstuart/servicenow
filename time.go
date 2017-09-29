package servicenow

import (
	"encoding/json"
	"time"
)

type SNTime struct {
	time.Time
}

const snFormat = "2006-01-02 15:04:05"

func (s *SNTime) UnmarshalJSON(bs []byte) error {
	var st string

	err := json.Unmarshal(bs, &st)
	if err != nil {
		return err
	}

	if len(st) < len(snFormat) {
		return nil
	}

	s.Time, err = time.Parse(snFormat, st)
	return err
}

func (s *SNTime) MarshalJSON() ([]byte, error) {
	return []byte(s.Time.Format(snFormat)), nil
}
