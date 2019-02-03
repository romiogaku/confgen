package backends

import (
	"testing"
)

func TestNewInvalidBackend(t *testing.T) {
	cases := map[string]struct {
		config Config
	}{
		"backend empty": {
			config: Config{
				TemplateFilePath: "",
				Backend:          "",
				BackendNodes:     []string{"", ""},
			},
		},
		"invalid backend": {
			config: Config{
				TemplateFilePath: "",
				Backend:          "INVALID_BACKEND",
				BackendNodes:     []string{"", ""},
			},
		},
	}

	for k, tt := range cases {
		t.Run(k, func(t *testing.T) {
			_, err := New(tt.config)
			if err == nil {
				t.Errorf("It should be error")
			}
		})
	}
}
