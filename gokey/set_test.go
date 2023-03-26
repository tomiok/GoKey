package gokey

import (
	"encoding/json"
	"testing"
)

func TestSetInts(t *testing.T) {
	client := newClient()

	_, err := client.Sadd("some key", []any{1, 2, 3, 4, 5})
	if err != nil {
		t.Fatal(err)
	}

	_values, err := client.Sget("some key")

	if err != nil {
		t.Fatal(err)
	}

	for _, v := range _values {
		var flag bool
		for _, i := range []string{"1", "2", "3", "4", "5"} {
			if v.(string) == i {
				flag = true
			}
		}
		if !flag {
			t.Fatal("is not in the list")
		}
	}
}

type V struct {
	ID   int
	Name string
}

func TestSetStructs(t *testing.T) {
	client := newClient()

	_, err := client.Sadd("key", []any{V{1, "tomas"}})

	if err != nil {
		t.Fatal(err)
	}

	res, err := client.Sget("key")

	if err != nil {
		t.Fatal(err)
	}

	val := res[0]

	var response V
	err = json.Unmarshal([]byte(val.(string)), &response)

	if err != nil {
		t.Fatal(err)
	}

	if response.ID != 1 {
		t.Fatal("id should be 1")
	}

	if response.Name != "tomas" {
		t.Fatal("name should be tomas")
	}
}
