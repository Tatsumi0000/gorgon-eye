package config

import (
	"reflect"
	"testing"
)

func Test_Read(t *testing.T) {
	t.Run("Can not open yaml file.", func(t *testing.T) {
		config := New("hogehoge.yaml")
		config.Read()
		got := config.Urls
		want := []string{"hogehoge", "fugafuga"}
		if reflect.DeepEqual(got, want) {
			t.Errorf("config.Urls == %s, want %s", got, want)
		}
	})
}
