package config

import (
	"flag"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

const (
	space string = " "
)

func Test_LoadConfig(t *testing.T) {
	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
	}()
	tests := []struct {
		name    string
		args    string
		want    Config
		wantErr bool
	}{
		{
			name: "empty args",
			args: "cmd",
			want: Config{
				Addr:     "localhost:8080",
				BaseAddr: "http://localhost:8080/",
			},
			wantErr: false,
		},
		{
			name: "only addr",
			args: "cmd -a ya.ru",
			want: Config{
				Addr:     "ya.ru",
				BaseAddr: "http://localhost:8080/",
			},
			wantErr: false,
		},
		{
			name: "only base",
			args: "cmd -b ya.ru",
			want: Config{
				Addr:     "localhost:8080",
				BaseAddr: "ya.ru/",
			},
			wantErr: false,
		},
		{
			name: "good args",
			args: "cmd -a ya.ru -b go.go",
			want: Config{
				Addr:     "ya.ru",
				BaseAddr: "go.go/",
			},
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			os.Args = strings.Split(test.args, space)
			config, err := LoadConfig()
			if !test.wantErr {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
			diff := cmp.Diff(test.want, config)
			assert.Equal(t, "", diff)
		})
	}

}
