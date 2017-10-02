package servicenow

import (
	"bytes"
	"time"
)

type SNTime struct {
	time.Time
}

const snFormat = "2006-01-02 15:04:05"

func (s *SNTime) UnmarshalJSON(bs []byte) error {
	st := string(bytes.Trim(bs, "\""))

	if len(st) < len(snFormat) {
		return nil
	}

	var err error
	s.Time, err = time.Parse(snFormat, st)
	return err
}

func (s *SNTime) MarshalJSON() ([]byte, error) {
	return []byte(s.Time.Format(snFormat)), nil
}
