package parser

import (
	"testing"
	"io/ioutil"
)

func TestParseCityList(t *testing.T) {
	content, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(content)

	const resultSize = 470

	if len(result.Requests) != resultSize {
		t.Errorf("error")
	}
	if len(result.Items) != resultSize {
		t.Errorf("error")
	}

}
