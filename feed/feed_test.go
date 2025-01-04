package feed

import "testing"

func Test_ParseYaml(t *testing.T) {
	t.Run("Can not open yaml file.", func(t *testing.T) {
		fileName := "hogehoge.yaml"
		got, err := ParseYaml(fileName)
		if err == nil {
			t.Errorf("Want error. Got value is %v", got)
		} else {
			t.Logf("Got error message: %s", err)
		}
	})

	t.Run("Can open yaml file.", func(t *testing.T) {
		fileName := "testdata/test_feed.yaml"
		got, err := ParseYaml(fileName)
		if err != nil {
			t.Errorf("Not want error. Got value is %v", err)
		} else {
			t.Logf("Got parse url is %v", got)
		}
	})
}

// func Test_ParseRss(t *testing.T) {
// 	t.Run("Debug", func(t *testing.T) {
// 		fileName := "testdata/test_feed.yaml"
// 		feeds, _ := ParseYaml(fileName)
// 		ParseRss(feeds)
// 	})
// }
