package thing

// New is a factory function for creating a new Config
func New(opts ...func(*Config)) *Config {
	config := &Config{}

	// Set a default value
	config.color = "blue"

	// apply the list of options to Config
	for _, opt := range opts {
		opt(config)
	}
	return config
}

// GetColor returns the color of the thing
func (c *Config) GetColor() string {
	return c.color
}

// GetDescription returns the description of the thing
func (c *Config) GetDescription() string {
	return c.description
}

// GetLabel returns the label of the thing
func (c *Config) GetLabel() string {
	return c.label
}

// GetSize returns the size of the thing
func (c *Config) GetSize() int {
	return c.size
}

// GetVersion returns the version of the thing
func (c *Config) GetVersion() string {
	return c.version
}

// SetColor sets the color of the thing
func SetColor(color string) Option {
	return func(c *Config) {
		c.color = color
	}
}

// SetDescription set the description of the thing
func SetDescription(description string) Option {
	return func(c *Config) {
		c.description = description
	}
}

// SetLabel sets the label of the thing
func SetLabel(label string) Option {
	return func(c *Config) {
		c.label = label
	}
}

// SetSize sets the size of the thing
func SetSize(size int) Option {
	return func(c *Config) {
		c.size = size
	}
}

// SetVersion sets the version of the thing
func SetVersion(version string) Option {
	return func(c *Config) {
		c.version = version
	}
}
