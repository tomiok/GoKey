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

func (c *Cache) SAdd(key string, v values) (bool, error) {
	setEntries := make(map[string]struct{})
	for _, r := range v {
		bytes, err := json.Marshal(r)
		if err != nil {
			panic(err)
		}

		setEntries[string(bytes)] = struct{}{}
	}

	c.setEntries[key] = setEntries
	return true, nil
}

func (c *Cache) SGet() ([]any, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Cache) SDelete() {
	//TODO implement me
	panic("implement me")
}
