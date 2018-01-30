package testgoupm_test

import (
    "github.com/pylbert/testgoupm"
    "testing"
)

func TestHello(t *testing.T) {
    c := testgoupm.NewTESTIT(5)
    s := c.Blibbity()
    if s != 666 {
        t.Error("unexpected value: ", s)
    }
}

