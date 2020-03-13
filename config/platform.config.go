package config

import "github.com/kelseyhightower/envconfig"

// Config struct of configuration
type Config struct {
	ServiceName      string   `envconfig:"SERVICE_NAME" default:"enterprise-app-platform"`
	PORT             string   `envconfig:"PORT" default:"80"`
	GRPCPORT         string   `envconfig:"GRPC_PORT" default:"8080"`
	PREFIX           string   `envconfig:"PREFIX" default:"enterpriseplatform"`
	INTERNALPASSWORD string   `envconfig:"INTERNAL_PASSWORD" default:"INTERNALAPIPASSWORD"`
	// DB Information
	DBAddress 		 string   `envconfig:"DBADDRESS" default:"localhost"`
	DBName 			 string   `envconfig:"DBNAME" default:"test"`
	DBUser 			 string   `envconfig:"DBUSER" default:"root"`
	DBPassword 		 string   `envconfig:"DBPASSWORD" default:""`
	DBPort 		 	 int   	  `envconfig:"DBPORT" default:"3306"`
	// GRPC Gateway Information
	GRPCClient		 string   `envconfig:"GRPCClient" default:"localhost"`
}

// singleton of data
var data *Config

// Get configuration of data
func Get() *Config {
	if data == nil {
		data = &Config{}
		envconfig.MustProcess("", data)
	}

	// returing configuration
	return data
}


