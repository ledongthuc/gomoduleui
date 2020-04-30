package models

import (
	"reflect"
	"testing"

	"golang.org/x/mod/modfile"
	"golang.org/x/mod/module"
)

func TestCreateMenuItemFromRequire(t *testing.T) {
	type args struct {
		module   modfile.Require
		replaced bool
	}
	tests := []struct {
		name string
		args args
		want MenuItem
	}{
		{
			name: "Empty",
			args: args{
				module:   modfile.Require{},
				replaced: false,
			},
			want: MenuItem{
				Label: "(Nothing)",
				Description: `  Module name: (Nothing)
  Version: (Nothing)
  Replacing: false
  Indirect: false`,
				Type: MenuTypeModule,
			},
		},
		{
			name: "Has data",
			args: args{
				module: modfile.Require{
					Mod: module.Version{
						Path:    "github.com/ledongthuc/gomoduleui",
						Version: "v0.0.1",
					},
					Indirect: true,
				},
				replaced: true,
			},
			want: MenuItem{
				Label: "github.com/ledongthuc/gomoduleui | v0.0.1",
				Description: `  Module name: github.com/ledongthuc/gomoduleui
  Version: v0.0.1
  Replacing: true
  Indirect: true`,
				Type: MenuTypeModule,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateMenuItemFromRequire(tt.args.module, tt.args.replaced); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateMenuItemFromRequire() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestCreateMenuItemsFromModFile(t *testing.T) {
	tests := []struct {
		name    string
		modfile modfile.File
		want    []MenuItem
	}{
		{
			name:    "Empty",
			modfile: modfile.File{},
			want:    []MenuItem{},
		},
		{
			name: "Has data",
			modfile: modfile.File{
				Require: []*modfile.Require{
					{
						Mod: module.Version{
							Path:    "github.com/ledongthuc/gomoduleui",
							Version: "v0.0.1",
						},
						Indirect: true,
					},
					{
						Mod: module.Version{
							Path:    "github.com/ledongthuc/test",
							Version: "v2.0.2",
						},
						Indirect: false,
					},
				},
			},
			want: []MenuItem{
				{
					Label: "github.com/ledongthuc/gomoduleui | v0.0.1",
					Description: `  Module name: github.com/ledongthuc/gomoduleui
  Version: v0.0.1
  Replacing: false
  Indirect: true`,
					Type: MenuTypeModule,
				},
				{
					Label: "github.com/ledongthuc/test | v2.0.2",
					Description: `  Module name: github.com/ledongthuc/test
  Version: v2.0.2
  Replacing: false
  Indirect: false`,
					Type: MenuTypeModule,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateMenuItemsFromModFile(&tt.modfile); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateMenuItemsFromModFile() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
