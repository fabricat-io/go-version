package version

import (
	"regexp"
	"strconv"
	"strings"
)

var semanticVersionRegex = func() *regexp.Regexp {
	var re *regexp.Regexp
	re, err := regexp.Compile(`^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)
	if err != nil {
		panic(err)
	}
	return re
}()

var semanticVersionRegexLenient = func() *regexp.Regexp {
	var re *regexp.Regexp
	re, err := regexp.Compile(`v?(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`)
	if err != nil {
		panic(err)
	}
	return re
}()

// SemanticVersion represents a semantic version as per the Semantic Versioning 2.0.0 specification.
type SemanticVersion struct {
	major         uint64
	minor         uint64
	patch         uint64
	preReleases   []string
	buildMetadata []string
}

// NewSemanticVersion creates a new SemanticVersion instance with the given major, minor, patch, pre-release and build metadata.
func NewSemanticVersion(major, minor, patch uint64, preReleases, buildMetadata []string) *SemanticVersion {
	return &SemanticVersion{
		major:         major,
		minor:         minor,
		patch:         patch,
		preReleases:   preReleases,
		buildMetadata: buildMetadata,
	}
}

// ParseSemanticVersion parses a semantic version string and returns a SemanticVersion instance.
func ParseSemanticVersion(version string) (SemanticVersion, error) {
	return parseSemanticVersion(version, semanticVersionRegex)
}

// ParseSemanticVersionLenient parses a semantic version string leniently and returns a SemanticVersion instance.
func ParseSemanticVersionLenient(version string) (SemanticVersion, error) {
	return parseSemanticVersion(version, semanticVersionRegexLenient)
}

func parseSemanticVersion(version string, reg *regexp.Regexp) (SemanticVersion, error) {
	matches := reg.FindStringSubmatch(version)
	if matches == nil {
		return SemanticVersion{}, ErrInvalidVersion
	}

	s := SemanticVersion{}
	for i, name := range semanticVersionRegex.SubexpNames() {
		if name == "" {
			continue
		}
		switch name {
		case "major":
			major, err := strconv.ParseUint(matches[i], 10, 64)
			if err != nil {
				return SemanticVersion{}, ErrInvalidMajorValue
			}
			s.major = major
		case "minor":
			minor, err := strconv.ParseUint(matches[i], 10, 64)
			if err != nil {
				return SemanticVersion{}, ErrInvalidMinorValue
			}
			s.minor = minor
		case "patch":
			patch, err := strconv.ParseUint(matches[i], 10, 64)
			if err != nil {
				return SemanticVersion{}, ErrInvalidPatchValue
			}
			s.patch = patch
		case "prerelease":
			if matches[i] != "" {
				s.preReleases = strings.Split(matches[i], ".")
			}
		case "buildmetadata":
			if matches[i] != "" {
				s.buildMetadata = strings.Split(matches[i], ".")
			}
		}
	}

	return s, nil
}

// MustParseSemanticVersion is a helper function that panics if the version string cannot be parsed.
func MustParseSemanticVersion(version string) SemanticVersion {
	v, err := ParseSemanticVersion(version)
	if err != nil {
		panic(err)
	}
	return v
}

func (v SemanticVersion) Major() uint64 {
	return v.major
}

func (v SemanticVersion) Minor() uint64 {
	return v.minor
}

func (v SemanticVersion) Patch() uint64 {
	return v.patch
}

func (v SemanticVersion) PreReleases() []string {
	return v.preReleases
}

func (v SemanticVersion) BuildMetadata() []string {
	return v.buildMetadata
}

func (v SemanticVersion) String() string {
	var sb strings.Builder
	sb.WriteString(strconv.FormatUint(v.major, 10))
	sb.WriteString(".")
	sb.WriteString(strconv.FormatUint(v.minor, 10))
	sb.WriteString(".")
	sb.WriteString(strconv.FormatUint(v.patch, 10))
	if len(v.preReleases) > 0 {
		for i, preRelease := range v.preReleases {
			if i > 0 {
				sb.WriteString(".")
			}
			sb.WriteString(preRelease)
		}
	}
	if len(v.buildMetadata) > 0 {
		sb.WriteString("+")
		for i, build := range v.buildMetadata {
			if i > 0 {
				sb.WriteString(".")
			}
			sb.WriteString(build)
		}
	}
	return sb.String()
}
