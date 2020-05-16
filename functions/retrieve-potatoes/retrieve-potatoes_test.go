package main

import "testing"

func TestHandler(t *testing.T) {
	id 		:= float64(42)
	req		:= Request{RequestId: id}
	res, _	:= Handler(req)

	if res.ResponseId != id {
		t.Errorf("ResponseId should be %f", id)
	}
}
