package models

import (
	"reflect"
	"testing"
)

const configTestContent = `setup:
  go_private: https://github.com/ledongthuc/privatelib
  custom_commands: |
    echo 'abc'
  hidden: true
module:
  https://github.com/ledongthuc/privatelib:
    default_replaces:
     - https://github.com/ledongthuc/privatelib
     - https://github.com/ledongthuc/privatelib
    default_updates:
     - v1.0.0
     - develop
     - master
    hidden: true
`

func TestParseConfigurations(t *testing.T) {
	tests := []struct {
		name    string
		content []byte
		want    Config
		wantErr bool
	}{
		{
			name:    "Empty",
			content: []byte{},
			want:    Config{},
			wantErr: false,
		},
		{
			name:    "Default",
			content: []byte(configTestContent),
			want: Config{
				Setup: ConfigSetup{
					GoPrivate:      "https://github.com/ledongthuc/privatelib",
					CustomCommands: "echo 'abc'\n",
					Hidden:         true,
				},
				Module: map[ModuleName]ConfigModule{
					"https://github.com/ledongthuc/privatelib": ConfigModule{
						DefaultReplaces: []ConfigModuleDefaultReplace{
							"https://github.com/ledongthuc/privatelib",
							"https://github.com/ledongthuc/privatelib",
						},
						DefaultUpdates: []ConfigModuleDefaultUpdateVersion{
							"v1.0.0",
							"develop",
							"master",
						},
						Hidden: true,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseConfigurations(tt.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseConfigurations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseConfigurations() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
