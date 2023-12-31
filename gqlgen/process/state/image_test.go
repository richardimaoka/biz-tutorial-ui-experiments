package state

import (
	_ "image/jpeg"
	_ "image/png"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		src            string
		expectedWidth  int
		expectedHeight int
	}{
		{"testdata/image/httpd.png", 531, 839},
	}

	for _, c := range cases {
		t.Run(c.src, func(t *testing.T) {
			img, err := NewImage("tutorial", c.src, "")
			if err != nil {
				t.Fatalf("failed to load image, %s", err)
			}

			if img.height != c.expectedHeight {
				t.Errorf("expected height = %d but got %d", c.expectedHeight, img.height)
			}
			if img.width != c.expectedWidth {
				t.Errorf("expected width = %d but got %d", c.expectedWidth, img.width)
			}
		})
	}
}
