package app

import (
	"fmt"
	"log"

	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/config"
	"github.com/uzzalhcse/go-gin-gorm-mvc-boilerplate/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func createConnectionPool(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func buildConnectionString(dbConfig config.DBConfig) string {
	connString := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s",
		dbConfig.DBHost,
		dbConfig.DBPort,
		dbConfig.DBDatabase,
		dbConfig.DBUsername,
		dbConfig.DBPassword,
	)
	//connString := fmt.Sprintf(dbConfig.DBUsername + ":" + dbConfig.DBPassword + "@tcp(" + dbConfig.DBHost + ":" + strconv.Itoa(dbConfig.DBPort) + ")/" + dbConfig.DBDatabase + "?charset=utf8&parseTime=true&loc=Local")

	return connString
}

func connectDB() error {
	rwConnString := buildConnectionString(config.Database.MasterDB)

	var err error
	DB, err = createConnectionPool(rwConnString)
	if err != nil {
		return fmt.Errorf("[INIT] failed to connect to the database: %v", err)
	}

	// roConnString := buildConnectionString(config.Database.SlaveDB)
	// var err2 error
	// DB, err2 = createConnectionPool(roConnString)
	// if err2 != nil {
	// 	return fmt.Errorf("[INIT] failed to connect to the database2: %v", err)
	// }
	// err = DB.Use(
	// 	dbresolver.Register(dbresolver.Config{
	// 		Replicas: []gorm.Dialector{
	// 			mysql.Open(roConnString),
	// 		},
	// 		Policy: dbresolver.RandomPolicy{},
	// 	}),
	// )
	// if err != nil {
	// 	return fmt.Errorf("[INIT] failed to configure database resolver: %v", err)
	// }

	log.Println("[INIT] database connected")

	err = database.Migrate(DB)
	if err != nil {
		log.Printf("[INIT] failed/skipping migrating database: %v\n", err)
	}

	return nil
}
