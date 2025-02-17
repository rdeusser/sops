package formats

import (
	"path/filepath"
	"strings"
)

// Format is an enum type
type Format int

const (
	Binary Format = iota
	Dotenv
	Ini
	Json
	Yaml
)

var stringToFormat = map[string]Format{
	"binary": Binary,
	"dotenv": Dotenv,
	"ini":    Ini,
	"json":   Json,
	"yaml":   Yaml,
}

// FormatFromString returns a Format from a string.
// This is used for converting string cli options.
func FormatFromString(formatString string) Format {
	format, found := stringToFormat[formatString]
	if !found {
		return Binary
	}
	return format
}

// IsYAMLFile returns true if a given file path corresponds to a YAML file
func IsYAMLFile(path string) bool {
	return strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml")
}

// IsJSONFile returns true if a given file path corresponds to a JSON file
func IsJSONFile(path string) bool {
	return strings.HasSuffix(path, ".json")
}

// IsEnvFile returns true if a given file path corresponds to a .env file
func IsEnvFile(path string) bool {
	return strings.HasSuffix(path, ".env")
}

// IsIniFile returns true if a given file path corresponds to a INI file
func IsIniFile(path string) bool {
	return strings.HasSuffix(path, ".ini")
}

// FormatForPath returns the correct format given the path to a file
func FormatForPath(path string) Format {
	if filepath.Ext(path) == "" {
		return Binary // default to binary
	}
	switch {
	case IsYAMLFile(path):
		return Yaml
	case IsJSONFile(path):
		return Json
	case IsEnvFile(path):
		return Dotenv
	case IsIniFile(path):
		return Ini
	default:
		return FormatForPath(strings.TrimSuffix(path, filepath.Ext(path)))
	}
}

// FormatForPathOrString returns the correct format-specific implementation
// of the Store interface given the formatString if specified, or the path to a file.
// This is to support the cli, where both are provided.
func FormatForPathOrString(path, format string) Format {
	formatFmt, found := stringToFormat[format]
	if !found {
		formatFmt = FormatForPath(path)
	}
	return formatFmt
}
