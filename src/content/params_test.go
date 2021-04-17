package content

import (
	"reflect"
	"testing"
)

func TestTryMapToParams(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []Param
	}{
		{name: "empty input", args: args{s: ""}, want: nil},
		{name: "valid input", args: args{s: "[i=checkbox, alias=rulesAccepted]"}, want: []Param{
			{Key: "i", Value: "checkbox"},
			{Key: "alias", Value: "rulesAccepted"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TryMapToParams(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TryMapToParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
