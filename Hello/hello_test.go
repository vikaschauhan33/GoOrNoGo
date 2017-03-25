package main

import "testing"

func TestName(t *testing.T){
  if getName() != "universe!"{
    t.Error("Response from getName is wrong")
  }
}
