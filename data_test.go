package gomnik

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestNewRequest(t *testing.T) {
	_ = NewRequest(604045891)
}

func TestDecodeResponse(t *testing.T) {
	dat, err := ioutil.ReadFile("samples/data_response_1563565442.19")
	if err != nil {
		t.Error(err)
	}
	resp, err := decodeResponse(dat)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", resp)
}
