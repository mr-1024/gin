package orm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var DbConn *gorm.DB

// InitDb 初始化数据库
func InitDb() (err error) {
	// 加载配置文件
	viper.New()
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err != nil {
		//log.Fatalf("read config failed :	%v ", err)
		fmt.Println(fmt.Sprintf("read config failed :	%v ", err))
	}

	fmt.Sprintln(viper.Get("app_name"))
	var config = make(map[string]interface{})
	var dbConfig = viper.GetStringMapString("database")
	config["app_name"] = viper.Get("app_name")
	config["db_host"] = dbConfig["db_host"]
	config["db_port"] = dbConfig["db_port"]
	config["db_database"] = dbConfig["db_database"]
	config["db_username"] = dbConfig["db_username"]
	config["db_password"] = dbConfig["db_password"]
	config["db_prefix"] = dbConfig["db_prefix"]
	config["db_driver"] = dbConfig["db_driver"]

	var args string
	args = fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=true", config["db_username"], config["db_password"], config["db_database"])
	fmt.Println(fmt.Sprintf("args:%s", args))

	DbConn, err = gorm.Open("mysql", args)
	if err != nil {
		panic(err)
	}

	DbConn.AutoMigrate(Post{})
	return
}
