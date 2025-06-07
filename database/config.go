package database

import (
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host       string `yaml:"host"`
	DBName     string `yaml:"dbName"`
	DBUsername string `yaml:"dbUsername"`
	DBPassword string `yaml:"dbPassword"`
	DBPort     string `yaml:"dbPort"`
}

func loadDatabaseConfig() string {
	_, testMode := os.LookupEnv("test_mode")
	if testMode {
		log.SetOutput(io.Discard)
	}

	custom_config_path, isset := os.LookupEnv("CONFIG_PATH")
	default_path := ""
	if isset {
		log.Printf("Loading custom config path: %s\n", custom_config_path)
		if custom_config_path[len(custom_config_path)-1] != '/' {
			custom_config_path += "/"
		}
		default_path = custom_config_path
	}
	log.Printf("loading config from path: " + default_path + "configy.yaml")
	file, err := os.ReadFile(default_path + "config.yaml")
	if err != nil {
		log.Fatal("Config file is missing", err)
	}
	var cfg DBConfig
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		log.Fatal("Error unmarshalling config file", err)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Host, cfg.DBUsername, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	return dsn
}
func SetupAndGetDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(loadDatabaseConfig()), &gorm.Config{})
	if err != nil {
		log.Fatal("Couldn't connect to db", err)
		panic("")
	}
	//creating extension is a superuser priviliege, so this will only check if the extension exists or not
	var postgisExists bool
	createExtensionResult := db.Raw("select exists(select 1 from pg_extension where extname = 'postgis')").Scan(&postgisExists)
	if createExtensionResult.Error != nil{
		//TODOish
		//this feels rather restrictive, the extension might exist but the privilege to query this might not? so it's maybe a good idea to disable this?
		log.Fatal("Error querying for postgis extension\n", createExtensionResult.Error)
	}
	if !postgisExists{
		log.Fatal("Postgis not available on the database")
	}
	log.Print("Database connected successfully!")
	return db
}
