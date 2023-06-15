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
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			args: args{
				token: "hello 你好吗?? good",
			},
			want: []string{"你好", "吗", "hello", "good"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SplitWorkByLanguage(tt.args.token)
			log.Println(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitWorkByLanguage() got = %v, want %v", got, tt.want)
			}
		})
	}
}
