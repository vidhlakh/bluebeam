package interleave

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_readLines(t *testing.T) {
	type args struct {
		path      string
		delimiter string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "success readlines",
			args: args{
				path:      "../inputs/file1.txt",
				delimiter: ",",
			},
			want: []string{"f1", "f3", "f5", "f7", "f9",
				"f11", "f13",
				"f21",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readLines(tt.args.path, tt.args.delimiter)
			if (err != nil) != tt.wantErr {
				t.Errorf("readLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterleaveFiles(t *testing.T) {
	type args struct {
		file1     string
		file2     string
		delimiter string
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantResult []string
	}{
		{
			name: "success test case with tab delimiter",
			args: args{

				file1:     "../inputs/tab/file1.txt",
				file2:     "../inputs/tab/file2.txt",
				delimiter: " ",
			},
			wantErr: false,
			wantResult: []string{"f1", "g2", "f3", "g4", "f5", "g6", "f7",
				"g8", "f9", "g12", "f11", "g14", "f13", "g20",
				"f21", "g32", "g34", "g36"},
		},
		{
			name: "success test case with comma delimiter",
			args: args{

				file1:     "../inputs/comma/file1.txt",
				file2:     "../inputs/comma/file2.txt",
				delimiter: ",",
			},
			wantErr: false,
			wantResult: []string{"f1", "g2", "f3", "g4", "f5", "g6", "f7",
				"g8", "f9", "g12", "f11", "g14", "f13", "g20",
				"f21", "g32", "g34", "g36"},
		},
		{
			name: "test case with double delimiter",
			args: args{

				file1:     "../inputs/doubledelimiter/file1.txt",
				file2:     "../inputs/doubledelimiter/file2.txt",
				delimiter: ",",
			},
			wantErr: false,
			wantResult: []string{"f1", "g2", "f3", "g4", "f5", "g6", "f7",
				"g8", "f9", "g12", "f11", "g14", "f13", "g20",
				"f21", "g32", "g34", "g36"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := InterleaveFiles(tt.args.file1, tt.args.file2, tt.args.delimiter)
			if (err != nil) != tt.wantErr {

				t.Errorf("InterleaveFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				fmt.Println("gotResult:", gotResult)
				t.Errorf("InterleaveFiles() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
