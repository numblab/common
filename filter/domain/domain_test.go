package domain

import (
	"testing"
)

func TestNewMatcher(t *testing.T) {
	m := NewMatcher([]string{"baidu.com"})
	if m.Match("tttbaidu.com") {
		t.Error("tttbaidu.com should not be matched")
	}
	if !m.Match("baidu.com") {
		t.Error("baidu.com should be matched")
	}
	if !m.Match("www.baidu.com") {
		t.Error("www.baidu.com should be matched")
	}
	if m.Match("www.baidu.com.cn") {
		t.Error("www.baidu.com.cn should not be matched")
	}
}

func TestNewMatcher2(t *testing.T) {
	m := NewMatcher([]string{""})
	if !m.Match("baidu.com") {
		t.Error("baidu.com should be matched")
	}
}

func TestNewMatcher3(t *testing.T) {
	m := NewMatcher([]string{"123.com", "."})
	if !m.Match("baidu.com") {
		t.Error("baidu.com should be matched")
	}
}
