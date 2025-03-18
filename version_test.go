package version

import (
	"reflect"
	"testing"
)

func TestParseSemanticVersion(t *testing.T) {
	type args struct {
		version string
	}
	tests := []struct {
		name    string
		args    args
		want    SemanticVersion
		wantErr bool
	}{
		{name: "valid version", args: args{version: "1.0.0"}, want: SemanticVersion{major: 1, minor: 0, patch: 0}, wantErr: false},
		{name: "valid version with pre-release", args: args{version: "1.0.0-alpha"}, want: SemanticVersion{major: 1, minor: 0, patch: 0, preReleases: []string{"alpha"}}, wantErr: false},
		{name: "valid version with build metadata", args: args{version: "1.0.0+build.123"}, want: SemanticVersion{major: 1, minor: 0, patch: 0, buildMetadata: []string{"build", "123"}}, wantErr: false},
		{name: "valid version with pre-release and build metadata", args: args{version: "1.0.0-alpha+build.123"}, want: SemanticVersion{major: 1, minor: 0, patch: 0, preReleases: []string{"alpha"}, buildMetadata: []string{"build", "123"}}, wantErr: false},
		{name: "invalid version", args: args{version: "1.0"}, want: SemanticVersion{}, wantErr: true},

		{name: "valid version", args: args{version: "v1.0.0"}, want: SemanticVersion{}, wantErr: true},
		{name: "valid version with pre-release", args: args{version: "v1.0.0-alpha"}, want: SemanticVersion{}, wantErr: true},
		{name: "valid version with build metadata", args: args{version: "v1.0.0+build.123"}, want: SemanticVersion{}, wantErr: true},
		{name: "valid version with pre-release and build metadata", args: args{version: "v1.0.0-alpha+build.123"}, want: SemanticVersion{}, wantErr: true},
		{name: "invalid version", args: args{version: "v1.0"}, want: SemanticVersion{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseSemanticVersion(tt.args.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSemanticVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseSemanticVersion() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseSemanticVersionLenient(t *testing.T) {
	type args struct {
		version string
	}
	tests := []struct {
		name    string
		args    args
		want    SemanticVersion
		wantErr bool
	}{
		{name: "valid version", args: args{version: "1.0.0"}, want: SemanticVersion{major: 1, minor: 0, patch: 0}, wantErr: false},
		{name: "valid version with pre-release", args: args{version: "1.0.0-alpha"}, want: SemanticVersion{major: 1, minor: 0, patch: 0, preReleases: []string{"alpha"}}, wantErr: false},
		{name: "valid version with build metadata", args: args{version: "1.0.0+build.123"}, want: SemanticVersion{major: 1, minor: 0, patch: 0, buildMetadata: []string{"build", "123"}}, wantErr: false},
		{name: "valid version with pre-release and build metadata", args: args{version: "1.0.0-alpha+build.123"}, want: SemanticVersion{major: 1, minor: 0, patch: 0, preReleases: []string{"alpha"}, buildMetadata: []string{"build", "123"}}, wantErr: false},
		{name: "invalid version", args: args{version: "1.0"}, want: SemanticVersion{}, wantErr: true},

		{name: "valid version", args: args{version: "v1.0.0"}, want: SemanticVersion{major: 1, minor: 0, patch: 0}, wantErr: false},
		{name: "valid version with pre-release", args: args{version: "v1.0.0-alpha"}, want: SemanticVersion{major: 1, minor: 0, patch: 0, preReleases: []string{"alpha"}}, wantErr: false},
		{name: "valid version with build metadata", args: args{version: "v1.0.0+build.123"}, want: SemanticVersion{major: 1, minor: 0, patch: 0, buildMetadata: []string{"build", "123"}}, wantErr: false},
		{name: "valid version with pre-release and build metadata", args: args{version: "v1.0.0-alpha+build.123"}, want: SemanticVersion{major: 1, minor: 0, patch: 0, preReleases: []string{"alpha"}, buildMetadata: []string{"build", "123"}}, wantErr: false},
		{name: "invalid version", args: args{version: "v1.0"}, want: SemanticVersion{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseSemanticVersionLenient(tt.args.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSemanticVersionLenient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseSemanticVersionLenient() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustParseSemanticVersion(t *testing.T) {
	type args struct {
		version string
	}
	tests := []struct {
		name  string
		args  args
		want  SemanticVersion
		panic bool
	}{
		{name: "valid version", args: args{version: "1.0.0"}, want: SemanticVersion{major: 1, minor: 0, patch: 0}, panic: false},
		{name: "valid version with pre-release", args: args{version: "1.0.0-alpha"}, want: SemanticVersion{major: 1, minor: 0, patch: 0, preReleases: []string{"alpha"}}, panic: false},
		{name: "valid version with build metadata", args: args{version: "1.0.0+build.123"}, want: SemanticVersion{major: 1, minor: 0, patch: 0, buildMetadata: []string{"build", "123"}}, panic: false},
		{name: "valid version with pre-release and build metadata", args: args{version: "1.0.0-alpha+build.123"}, want: SemanticVersion{major: 1, minor: 0, patch: 0, preReleases: []string{"alpha"}, buildMetadata: []string{"build", "123"}}, panic: false},
		{name: "invalid version", args: args{version: "1.0"}, want: SemanticVersion{}, panic: true},

		{name: "valid version", args: args{version: "v1.0.0"}, want: SemanticVersion{}, panic: true},
		{name: "valid version with pre-release", args: args{version: "v1.0.0-alpha"}, want: SemanticVersion{}, panic: true},
		{name: "valid version with build metadata", args: args{version: "v1.0.0+build.123"}, want: SemanticVersion{}, panic: true},
		{name: "valid version with pre-release and build metadata", args: args{version: "v1.0.0-alpha+build.123"}, want: SemanticVersion{}, panic: true},
		{name: "invalid version", args: args{version: "v1.0"}, want: SemanticVersion{}, panic: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.panic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("MustSemanticVersion() did not panic")
					}
				}()
			}
			got := MustParseSemanticVersion(tt.args.version)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustSemanticVersion() got = %v, want %v", got, tt.want)
			}
		})
	}
}
