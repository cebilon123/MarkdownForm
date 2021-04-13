package decode

import (
	"core"
	"reflect"
	"testing"
)

func TestMarkdownDecoder_Decode(t *testing.T) {
	type fields struct {
		FileContentText string
	}

	tests := []struct {
		name        string
		fields      fields
		want        *core.Form
		shouldPanic bool
	}{
		{name: "NoContentToDecode_PanicReturnNil", fields: struct{ FileContentText string }{FileContentText: ""}, want: nil, shouldPanic: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MarkdownDecoder{
				FileContentText: tt.fields.FileContentText,
			}
			if tt.shouldPanic {
				defer func() { recover() }()
				m.Decode()
				t.Errorf("should have panicked")
				return
			}
			if got := m.Decode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func shouldPanic(t *testing.T, f func()) {
	defer func() { recover() }()
	f()
	t.Errorf("should have panicked")
}
