package config

/**
* Package used to provide access to the bank
* @package config
* @author Jordan Duarte
**/

import (
    "fmt"
    "log"
    "github.com/BurntSushi/toml"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var config = ConfigDB{}

// ConfigDB db seting
type ConfigDB struct {
    User             string
    Password         string
    Host             string
    Port             string
    Dbname           string
}

type Options struct {
    IsDefaultDbName bool
}

// ConnectDB returns initialized gorm.DB
func ConnectDB(options *Options) (*gorm.DB, error) {
    config.Read()

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.User, config.Password, config.Host, config.Port, "mysql")

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if (options.IsDefaultDbName) {
        useDatatable := fmt.Sprintf("USE %s", config.Dbname)
        db.Exec(useDatatable)
    }
    if err != nil {
        return nil, err
    }
    return db, nil
}

// Read and parse the configuration file
func (c *ConfigDB) Read() {
    if _, err := toml.DecodeFile("config.toml", &c); err != nil {
        log.Fatal(err)
    }
}