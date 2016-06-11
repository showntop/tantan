package config

//the server and db's configure
//can save in the file when start load it

var (
	Configure = map[string]interface{}{
		"port": "10000",
		"database": map[string]string{
			"addr":     "localhost:9000",
			"user":     "showntop",
			"password": "1",
			"dbname":   "tantan",
			"sslmode":  "false",
		},
	}
)
