package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var (
	defaultDB *sqlx.DB
)

type MySQL struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

func InitMySQL() {
	mysqlCfg := MySQL{
		Host:     viper.GetString("mysql.host"),
		Port:     viper.GetInt("mysql.port"),
		User:     viper.GetString("mysql.user"),
		Password: viper.GetString("mysql.password"),
		Database: viper.GetString("mysql.database"),
	}
	
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true&loc=Local",
		mysqlCfg.User, mysqlCfg.Password, mysqlCfg.Host, mysqlCfg.Port, mysqlCfg.Database))
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		panic("Failed to ping database: %v" + err.Error())
	}

	// 数据库表初始化
	if err := initMysqlTable(db); err != nil {
		panic("Failed to init database: " + err.Error())
	}

	defaultDB = db
}

func initMysqlTable(db *sqlx.DB) error {
	// order 订单表
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS orders  (
			ID int(11) NOT NULL AUTO_INCREMENT,
			ProductID int(11) NOT NULL,
			CustomerID int(11) NOT NULL,
			OrderID varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
			OrderDate date NOT NULL,
			CdKey varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
			LogisticsStatus int(11) NOT NULL,
			Addr varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
			PRIMARY KEY (ID) USING BTREE
		) ENGINE = InnoDB;`)
	if err != nil {
		return err
	}

	return nil
}

func GetMySQLClient() *sqlx.DB {
	return defaultDB
}
