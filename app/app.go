package app

import (
	"fmt"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/viper"
)

// Config stores the application-wide configurations
var Config appConfig

type appConfig struct {
	// authentication server port. Defaults to 8082
	AuthSrvPort int `mapstructure:"s_auth_port"`
	// timesheet server port. Defaults to 8083
	TSSrvPort int `mapstructure:"s_ts_port"`
	// the data source name (DSN) for connecting to the database. required. mongodb obligatory
	DSN string `mapstructure:"dsn"`
	// JWT signing key. required.
	JWTSigningKey string `mapstructure:"jwt_signing_key"`
	// JWT verification key. required.
	JWTVerificationKey string `mapstructure:"jwt_verification_key"`
	/// Scrypt hash function salt
	ScryptSalt string `mapstructure:"scrypt_salt"`
}

func (config appConfig) Validate() error {
	return validation.ValidateStruct(&config,
		validation.Field(&config.DSN, validation.Required),
		validation.Field(&config.JWTSigningKey, validation.Required),
		validation.Field(&config.JWTVerificationKey, validation.Required),
		validation.Field(&config.ScryptSalt, validation.Required),
	)
}

// LoadConfig loads configuration from the given list of paths and populates it into the Config variable.
// The configuration file(s) should be named as app.yaml.
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("singularity")
	v.AutomaticEnv()
	v.SetDefault("s_auth_port", 8082)
	v.SetDefault("s_ts_port", 8083)
	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("Failed to read the configuration file: %s", err)
	}
	if err := v.Unmarshal(&Config); err != nil {
		return err
	}
	return Config.Validate()
}
