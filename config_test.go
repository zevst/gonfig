package gonfig

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func init() {
	if err := os.Setenv("TEST_GET_ENV_STR", "THIS IS TEST"); err != nil {
		panic(err)
	}
	if err := os.Setenv("PORT", "8080"); err != nil {
		panic(err)
	}
	if err := os.Setenv("APP_MODE", "test"); err != nil {
		panic(err)
	}
}

func TestGetEnvStr(t *testing.T) {
	type args struct {
		key string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				key: "TEST_GET_ENV_STR",
			},
			want: "THIS IS TEST",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnvStr(tt.args.key); got != tt.want {
				t.Errorf("GetEnvStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetListenPort(t *testing.T) {
	port := "8080"
	tests := []struct {
		name string
		want *string
	}{
		{
			want: &port,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetListenPort(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetListenPort() = %v, want %v", got, tt.want)
			} else {
				fmt.Println("Good")
			}
		})
	}
}

func TestGetApplicationMode(t *testing.T) {
	mode := "test"
	tests := []struct {
		name string
		want *string
	}{
		{
			want: &mode,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetApplicationMode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetApplicationMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
