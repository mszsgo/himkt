package tcc

import (
	"errors"
	"testing"
	"time"
)

func TestTx(t *testing.T) {
	var a = "10"
	err := Tx(
		NewTcc(func() error {
			t.Log("T1....try" + a)
			return nil
		}, func() error {
			t.Log("T1....confirm" + a)
			return nil
		}, func() error {
			t.Log("T1....cancel" + a)
			return nil
		}),
		NewTcc(func() error {
			t.Log("T2....try")
			return nil
		}, func() error {
			t.Log("T2....confirm")
			return nil
		}, func() error {
			t.Log("T2....cancel")
			return errors.New("t2-cancel-err....")
		}),
		NewTcc(func() error {
			t.Log("T3....try")
			return errors.New("t3-try-报个错")
		}, func() error {
			t.Log("T3....confirm")
			return nil
		}, func() error {
			t.Log("T3....cancel")
			return nil
		}),
	)
	if err != nil {
		t.Error(err)
	}
	time.Sleep(5 * time.Second)
}
