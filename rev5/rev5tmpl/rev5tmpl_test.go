package rev5tmpl

import (
	"testing"
)

var renderTests = []struct {
	v      string
	params map[string]string
	want   string
}{
	{
		"abc {{ insert: param, foo.bar }} xyz",
		map[string]string{"foo.bar": "lmn"},
		"abc lmn xyz",
	},
}

func TestRender(t *testing.T) {
	for _, tt := range renderTests {
		try, err := Render(tt.v, tt.params)
		if err != nil {
			t.Errorf("rev5tmpl.Render(\"%s\",%v) error: (%s)", tt.v, tt.params, err.Error())
		}
		if try != tt.want {
			t.Errorf("rev5tmpl.Render(\"%s\",%v) mismatch: want (%s), got (%s)", tt.v, tt.params, tt.want, try)
		}
	}
}
