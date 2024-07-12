package db

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Config represents application configuration settings.
type Config struct {
	DatabaseDSN string `mapstructure:"database_dsn"`
}

// LoadConfig loads configuration from a file.
func LoadConfig() (*Config, error) {
	// Initialize viper
	viper.SetConfigName("config")   // Name of the configuration file (config.yaml)
	viper.AddConfigPath("./config") // Search the config directory for the configuration file
	viper.SetConfigType("yaml")     // Config file format

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Unmarshal the configuration into a struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}

	return &config, nil
}

// ConnectDatabase initializes and returns a connection to the SQL Server database.
// The function uses GORM for ORM capabilities and logs the connection process.
//
// Returns:
// - *gorm.DB: a pointer to the database connection instance.
// - error: an error object if the connection fails, nil otherwise.
func ConnectDatabase() (*gorm.DB, error) {

	// Load configuration from file
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	// Open the database connection using the specified DSN (Data Source Name)
	db, dbErr := gorm.Open(sqlserver.Open(config.DatabaseDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Set logger to log detailed information
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   // No prefix for table names
			SingularTable: true, // Use singular table names
			NoLowerCase:   true, // Preserve original case for table names
		},
	})
	if dbErr != nil {
		log.Fatal("failed to connect database:", dbErr)
		return nil, dbErr
	}

	log.Println("Database connection established")
	return db, nil
}
