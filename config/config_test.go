package config

import (
	"testing"
)

func Test_Read(t *testing.T) {
	t.Run("Can not open yaml file.", func(t *testing.T) {
		fileName := "hogehoge.yaml"
		c := New(fileName)
		got := c.Read()
		if got == nil {
			t.Errorf("c.Read() == %s, want error", got)
		} else {
			t.Logf("got error message: %s", got)
		}
	})
}
