package gokey

import "encoding/json"

type setEntry struct {
	entry string
}

func (s setEntry) MarshalText() ([]byte, error) {
	return json.Marshal(s)
}

func (s *setEntry) UnmarshalText(b []byte) error {
	return json.Unmarshal(b, s)
}
