package main

import (
    "testing"
)


func TestInit(t *testing.T) {
  buffer := "test"
  s := &Sysla{Buffer: string(buffer)}
  s.Init()
}
