package teststack

type Config struct {
	AccessKey string
	SecretKey string
	Region    string
}

func (c *Config) Client() (*Client, error) {
	var client Client
	client.Region = c.Region
	return &client, nil
}
