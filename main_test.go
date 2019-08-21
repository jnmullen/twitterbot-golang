package main

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	output := m.Run()
	fmt.Println(output)
	//if strings.Compare(output, "hello world") != 0 {
	//}
}
