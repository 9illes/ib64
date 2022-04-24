package base64

import (
	"reflect"
	"testing"
)

func Test_stringToBinaryNotation(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name             string
		args             args
		wantBinaryString string
	}{
		{
			"simple letter 1",
			args{s: "A"},
			"01000001",
		},
		{
			"simple string 1",
			args{s: "QUJD"},
			"01010001010101010100101001000100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBinaryString := stringToBinaryNotation(tt.args.s); gotBinaryString != tt.wantBinaryString {
				t.Errorf("stringToBinaryNotation() = %v, want %v", gotBinaryString, tt.wantBinaryString)
			}
		})
	}
}

func Test_createStringsOf6bits(t *testing.T) {
	type args struct {
		bs string
	}
	tests := []struct {
		name       string
		args       args
		wantGroups []string
	}{
		{
			"sample 1 (QUJD)",
			args{
				bs: "01010001010101010100101001000100"},
			[]string{
				"010100",
				"010101",
				"010101",
				"001010",
				"010001",
				"000000",
			},
		},
		{
			"sample 1 with padding",
			args{
				bs: "00000011111100000011111100000011"},
			[]string{
				"000000",
				"111111",
				"000000",
				"111111",
				"000000",
				"110000",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotGroups := createStringsOf6bits(tt.args.bs); !reflect.DeepEqual(gotGroups, tt.wantGroups) {
				t.Errorf("createStringsOf6bits() = %v, want %v", gotGroups, tt.wantGroups)
			}
		})
	}
}

func Test_padBits(t *testing.T) {
	type args struct {
		groups []string
	}
	tests := []struct {
		name            string
		args            args
		wantBytesGroups []string
	}{
		{"simple", args{
			groups: []string{
				"010100",
				"010101",
				"010101",
				"001010",
				"010001",
				"000000",
			}},
			[]string{
				"00010100",
				"00010101",
				"00010101",
				"00001010",
				"00010001",
				"00000000",
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBytesGroups := padBits(tt.args.groups); !reflect.DeepEqual(gotBytesGroups, tt.wantBytesGroups) {
				t.Errorf("padBits() = %v, want %v", gotBytesGroups, tt.wantBytesGroups)
			}
		})
	}
}

func Test_indices2string(t *testing.T) {
	type args struct {
		indices []string
	}
	tests := []struct {
		name  string
		args  args
		wantS string
	}{
		{
			"simple",
			args{
				indices: []string{
					"00010100",
					"00010101",
					"00010101",
					"00001010",
					"00010001",
					"00000000",
				}},
			"UVVKRA"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := indices2string(tt.args.indices); gotS != tt.wantS {
				t.Errorf("indices2string() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name         string
		args         args
		wantSEncoded string
	}{
		{
			"test string with padding",
			args{s: "padding"},
			"cGFkZGluZw==",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSEncoded := Encode(tt.args.s); gotSEncoded != tt.wantSEncoded {
				t.Errorf("Encode() = %v, want %v", gotSEncoded, tt.wantSEncoded)
			}
		})
	}
}
