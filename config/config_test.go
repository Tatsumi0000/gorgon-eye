package config_test

import (
	"testing"

	"github.com/Tatsumi0000/gorgon-eye/config"
)

func Test_Read(t *testing.T) {
	t.Run("Can not open yaml file.", func(t *testing.T) {
		fileName := "hogehoge.yaml"
		config := config.New(fileName)
		err := config.Read()
		got := config.Urls
		want := []string{"hogehoge", "fugafuga"}
		if err != nil {
			t.Log(err)
			t.Errorf("config.Urls == %s, want %s", got, want)
		}
	})
}
