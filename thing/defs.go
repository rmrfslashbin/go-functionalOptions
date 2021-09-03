package thing

// Used to manage varidic options
type Option func(c *Config)

// Config is the configuration for the thing.
type Config struct {
	// Color
	color string

	// Description
	description string

	// Label
	label string

	// Size
	size int

	// Version
	version string
}
