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

func TestIsParamContainableType(t *testing.T) {
	type args struct {
		valueType ValueType
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Is containable - button", args: args{valueType: ButtonLine}, want: true},
		{name: "Is containable - input", args: args{valueType: InputLine}, want: true},
		{name: "Is not containable - heading", args: args{valueType: HeadingLine}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsParamContainableType(tt.args.valueType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsParamContainableType() = %v, want %v", got, tt.want)
			}
		})
	}
}
