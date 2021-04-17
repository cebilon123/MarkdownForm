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
		{name: "not valid prefix", args: args{s: "i=test,alias=best]"}, want: nil},
		{name: "not valid input body", args: args{s: "[i,alias]"}, want: []Param{}},
		{name: "valid input", args: args{s: "[i=checkbox, alias=rulesAccepted]"}, want: []Param{
			{Key: "i", Value: "checkbox"},
			{Key: "alias", Value: "rulesAccepted"},
		}},
		{name: "semi-valid input", args: args{s: "[i,alias=btn]"}, want: []Param{
			{Key: "alias", Value: "btn"},
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
