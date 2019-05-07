package test

import (
	"fmt"
	"github.com/idata-shopee/mc_service/mc"
	"testing"
)

func assertEqual(t *testing.T, expect interface{}, actual interface{}, message string) {
	if expect == actual {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("expect %v !=  actual %v", expect, actual)
	}
	t.Fatal(message)
}

// simple case
func TestBase(t *testing.T) {
	mm := mc.GetMemMap()
	mm.Set("a", 1)
	v, _ := mm.Get("a")

	assertEqual(t, v, 1, "")
}

func TestConcurrent(t *testing.T) {
	mm := mc.GetMemMap()
	for i := 0; i < 10000; i++ {
		go (func() {
			mm.Set("a", 1)
			v, _ := mm.Get("a")
			assertEqual(t, v, 1, "")
		})()
	}
}

func TestJsonPath(t *testing.T) {
	mm := mc.GetMemMap()
	mm.Set("a.b", 1)
	v, _ := mm.Get("a.b")

	assertEqual(t, v, 1, "")
}

func TestJsonPath2(t *testing.T) {
	mm := mc.GetMemMap()
	mm.Set("a.b.c", 1)
	v, _ := mm.Get("a.b.c")
	assertEqual(t, v, 1, "")
}

func TestConcurrent2(t *testing.T) {
	mm := mc.GetMemMap()
	for i := 0; i < 10000; i++ {
		go (func() {
			mm.Set("a.b.c", 1)
			v, _ := mm.Get("a.b.c")
			assertEqual(t, v, 1, "")
		})()
	}
}
