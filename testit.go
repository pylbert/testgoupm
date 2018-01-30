package simplelib_test

import (
    "github.com/pylbert/testgoupm"
    "testing"
)

func TestHello(t *testing.T) {
    c := goupm_testit.NewTESTIT()
    s := c.Blibbity()
    if s != 666 {
        t.Error("unexpected value: ", s)
    }
}

