package version

import "errors"

var (
	// ErrInvalidVersion indicates that the version is invalid
	ErrInvalidVersion = errors.New("invalid version format")

	// ErrInvalidMajorValue indicates that the major version is invalid
	ErrInvalidMajorValue = errors.New("invalid major version value")

	// ErrInvalidMinorValue indicates that the minor version is invalid
	ErrInvalidMinorValue = errors.New("invalid minor version value")

	// ErrInvalidPatchValue indicates that the patch version is invalid
	ErrInvalidPatchValue = errors.New("invalid patch version value")
)
