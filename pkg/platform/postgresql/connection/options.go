package connection

type Option func(*Connection)

func WithUser(user string) Option {
	return func(c *Connection) {
		c.User = user
	}
}

func WithPassword(password string) Option {
	return func(c *Connection) {
		c.Password = password
	}
}

func WithDatabase(database string) Option {
	return func(c *Connection) {
		c.Database = database
	}
}

func WithHost(host string) Option {
	return func(c *Connection) {
		c.Host = host
	}
}

func WithPort(port string) Option {
	return func(c *Connection) {
		c.Port = port
	}
}

func WithSSLMode(sslMode string) Option {
	return func(c *Connection) {
		c.SSLMode = sslMode
	}
}

func WithConnectionName(name string) Option {
	return func(c *Connection) {
		c.ConnectionName = name
	}
}

func WithDialFunc(dialFunc DialFunc) Option {
	return func(c *Connection) {
		c.DialFunc = dialFunc
	}
}
