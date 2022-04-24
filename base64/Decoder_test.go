package base64

import (
	"reflect"
	"testing"
)

func Test_getIndices(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name        string
		args        args
		wantIndices []int
	}{
		{"simple", args{s: "QUJD"}, []int{16, 20, 9, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIndices := getIndices(tt.args.s); !reflect.DeepEqual(gotIndices, tt.wantIndices) {
				t.Errorf("getIndices() = %v, want %v", gotIndices, tt.wantIndices)
			}
		})
	}
}

func Test_indices2BinaryNotation(t *testing.T) {
	type args struct {
		indices []int
	}
	tests := []struct {
		name         string
		args         args
		wantBinaries []string
	}{
		{
			"simple",
			args{indices: []int{16, 20, 9, 3}},
			[]string{
				"010000",
				"010100",
				"001001",
				"000011",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBinaries := indices2BinaryNotation(tt.args.indices); !reflect.DeepEqual(gotBinaries, tt.wantBinaries) {
				t.Errorf("indices2BinaryNotation() = %v, want %v", gotBinaries, tt.wantBinaries)
			}
		})
	}
}

func Test_extractBytes(t *testing.T) {
	type args struct {
		bs string
	}
	tests := []struct {
		name      string
		args      args
		wantBytes []string
	}{
		{
			"simple",
			args{bs: "010000010100001001000011"},
			[]string{
				"01000001",
				"01000010",
				"01000011",
			},
		},
		{
			"discard incomplete bytes",
			args{bs: "01000001010000100100001100"},
			[]string{
				"01000001",
				"01000010",
				"01000011",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBytes := extractBytes(tt.args.bs); !reflect.DeepEqual(gotBytes, tt.wantBytes) {
				t.Errorf("extractBytes() = %v, want %v", gotBytes, tt.wantBytes)
			}
		})
	}
}

func Test_binaryNotationToASCII(t *testing.T) {
	type args struct {
		bytes []string
	}
	tests := []struct {
		name         string
		args         args
		wantSDecoded string
	}{
		{
			"simple",
			args{bytes: []string{
				"01000001",
				"01000010",
				"01000011",
			},
			},
			"ABC"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSDecoded := binaryNotationToASCII(tt.args.bytes); gotSDecoded != tt.wantSDecoded {
				t.Errorf("binaryNotationToASCII() = %v, want %v", gotSDecoded, tt.wantSDecoded)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name         string
		args         args
		wantSDecoded string
	}{
		{"simple", args{s: "QUJD"}, "ABC"},
		{"missing padding", args{s: "cGFkZGluZw"}, "padding"},
		{"with padding", args{s: "cGFkZGluZw=="}, "padding"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSDecoded := Decode(tt.args.s); gotSDecoded != tt.wantSDecoded {
				t.Errorf("Decode() = %v, want %v", gotSDecoded, tt.wantSDecoded)
			}
		})
	}
}
