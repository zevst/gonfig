package gonfig

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

var envs = map[string]string{
	"TEST_GET_ENV_STR":     "THIS IS TEST",
	"TEST_GET_ENV_ARR_STR": "THIS;IS;TEST",
	"TEST_GET_ENV_INT":     "202",
	"PORT":                 "8080",
	"APP_MODE":             "test",
}

func init() {
	for key, value := range envs {
		if err := os.Setenv(key, value); err != nil {
			panic(err)
		}
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
		{
			args: args{
				key: "",
			},
			want: "",
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

func TestGetEnvArrStr(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				key: "TEST_GET_ENV_ARR_STR",
			},
			want: []string{"THIS", "IS", "TEST"},
		},
		{
			args: args{
				key: "",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnvArrStr(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEnvArrStr() = %v, want %v", got, tt.want)
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

func TestGetEnvInt(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				key: "TEST_GET_ENV_INT",
			},
			want: 202,
		},
		{
			args: args{
				key: "TEST_GET_ENV_STR",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnvInt(tt.args.key); got != tt.want {
				t.Errorf("GetEnvInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnvIntWithDefault(t *testing.T) {
	type args struct {
		key          string
		defaultValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				key:          "TEST_GET_ENV_INT",
				defaultValue: 202,
			},
			want: 202,
		},
		{
			args: args{
				key:          "TEST_GET_ENV_STR",
				defaultValue: 202,
			},
			want: 202,
		},
		{
			args: args{
				key: "TEST_GET_ENV_STR",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnvIntWithDefault(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetEnvIntWithDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnvStrWithDefault(t *testing.T) {
	type args struct {
		key          string
		defaultValue string
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
		{
			args: args{
				key:          "DEFAULT_VALUE",
				defaultValue: "DEFAULT_VALUE",
			},
			want: "DEFAULT_VALUE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnvStrWithDefault(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetEnvStrWithDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
