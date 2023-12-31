package state_test

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		src    string
		width  int
		height int
	}{
		{"testdata/image/httpd.png", 531, 839},
	}

	for _, c := range cases {
		t.Run(c.src, func(t *testing.T) {
			reader, err := os.Open(c.src)
			defer reader.Close()

			if err != nil {
				t.Fatal(err)
			}

			im, _, err := image.DecodeConfig(reader)
			if err != nil {
				t.Fatal(err)
			}

			if im.Height != c.height {
				t.Errorf("expected height = %d but got %d", c.height, im.Height)
			}
			if im.Width != c.width {
				t.Errorf("expected width = %d but got %d", c.width, im.Width)
			}
		})
	}
}
