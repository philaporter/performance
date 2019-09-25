package main

import (
	i "performance/server/initialize"
	"sync"
	"testing"
)

func TestBuildResponse(t *testing.T) {
	i.Init()
	s := buildResponse()
	if s != "" {
		return
	}
	t.Errorf("error building response")
}

func TestBuildResponseV2(t *testing.T) {
	i.Init()
	s := buildResponseV2()
	if s != "" {
		return
	}
	t.Errorf("error building response")
}

func TestBuildResponseV3(t *testing.T) {
	i.Init()
	s := buildResponseV3()
	if s != "" {
		return
	}
	t.Errorf("error building response")
}

func TestConverString(t *testing.T) {
	sm := sync.Map{}
	sm.Store("1", "one")
	sm.Store("2", 2)

	one, _ := sm.Load("1")
	two, _ := sm.Load("2")

	if convertString(one) != "one" {
		t.Errorf("convertString should have converted the sync.Map value to \"one\"")
	}
	if convertString(two) != "" {
		t.Errorf("convertString should have hit the default case for \"\" because the value isn't type string")
	}
}
