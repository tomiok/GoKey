package gokey

import "encoding/json"

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

func (c *Cache) SGet(key string) ([]any, error) {
	_set := c.setEntries[key]

	_values := make([]any, 0, len(_set))
	for k, _ := range _set {
		_values = append(_values, k)
	}
	return _values, nil
}

func (c *Cache) SDelete(key string) bool {
	//TODO implement me
	panic("implement me")
}
