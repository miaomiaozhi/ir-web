package pkg

import (
	"log"
	"reflect"
	"testing"
)

func TestSplitWorkByLanguage(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 []string
	}{
		// TODO: Add test cases.
		{
			args: args{
				token: "hello 你好吗?? good",
			},
			want:  []string{"你好", "吗"},
			want1: []string{"hello", "good"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SplitWorkByLanguage(tt.args.token)
			log.Println(got)
			log.Println(got1)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitWorkByLanguage() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("SplitWorkByLanguage() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
