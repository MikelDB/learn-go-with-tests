package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello("Mikel")
    want := "Hello, Mikel"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}