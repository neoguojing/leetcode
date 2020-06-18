package test

import(
	"testing"
	"leetcode/string"
)


func TestIsMatch(t *testing.T){
	text := "aa"
	p := "a*"

	ret := string.IsMatch(text,p)
	if(!ret) {
		t.Error("wrong")
	}

	text = "aa"
	p = ".*"

	ret = string.IsMatch(text,p)
	if(!ret) {
		t.Error("wrong")
	}

	text = "aa"
	p = "a"

	ret = string.IsMatch(text,p)
	if(ret) {
		t.Error("wrong")
	}
}