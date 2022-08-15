package interleave

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_readLines(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    wordsList
		wantErr bool
	}{
		{
			name: "success readlines",
			args: args{
				path: "../inputs/tab/file1.txt",
			},
			want: []string{"I", "going", "Switzerland",
				"airplane", "2", "books", "library",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readLines(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("readLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				fmt.Printf("got :%T", got)
				fmt.Printf("want: %T", tt.want)
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
			wantResult: []string{"I", "am", "going", "to", "Switzerland",
				"by", "airplane", "Love", "2", "read", "books", "from", "library"},
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

				file1:     "../inputs/multidelimiter/file1.txt",
				file2:     "../inputs/multidelimiter/file2.txt",
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
			gotResult, err := InterleaveFiles(tt.args.file1, tt.args.file2)
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

func Test_wordsList_removeSingleSymbols(t *testing.T) {
	tests := []struct {
		name string
		j    *wordsList
		want *wordsList
	}{
		{
			name: "success testcase",
			j:    &wordsList{"f1", "!f3", "$25", "f^7", "$", "f"},
			want: &wordsList{"f1", "!f3", "$25", "f^7", "f"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.j.removeSingleSymbols()
			if !reflect.DeepEqual(tt.j, tt.want) {
				t.Errorf("InterleaveFiles() = %v, want %v", tt.j, tt.want)
			}
		})
	}
}
