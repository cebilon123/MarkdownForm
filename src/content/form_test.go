package content

import (
	"reflect"
	"testing"
)

func TestBaseForm_ToExtendedForm(t *testing.T) {
	type fields struct {
		Fields []LineField
	}
	tests := []struct {
		name   string
		fields fields
		want   *ExtendedForm
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &BaseForm{
				Fields: tt.fields.Fields,
			}
			if got := f.ToExtendedForm(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToExtendedForm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLineField_ToExtendedLineField(t *testing.T) {
	type fields struct {
		Type  ValueType
		Value string
	}
	tests := []struct {
		name   string
		fields fields
		want   ExtendedLineField
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lf := &LineField{
				Type:  tt.fields.Type,
				Value: tt.fields.Value,
			}
			if got := lf.ToExtendedLineField(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToExtendedLineField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCleanedValueBasedOnType(t *testing.T) {
	type args struct {
		s  string
		vt ValueType
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Is valid - heading line", args: args{
			s:  "##Hello",
			vt: HeadingLine,
		}, want: "Hello"},
		{name: "Is_valid_title_line", args: args{
			s:  "**Hello**",
			vt: TitleLine,
		}, want: "Hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCleanedValueBasedOnType(tt.args.s, tt.args.vt); got != tt.want {
				t.Errorf("getCleanedValueBasedOnType() = %v, want %v", got, tt.want)
			}
		})
	}
}
