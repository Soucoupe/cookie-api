package db

import (
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// ConnectDb is used to create a new schema for each model
func ConnectDb(options *pg.Options) *pg.DB {
	db := pg.Connect(options)

	log.Print("Connected to the DB successfully.")

	err := createSchema(db)
	if err != nil {
		panic(err)
	}

	log.Print("Synchronized all the schemas.")

	return db
}

func InsertSensor(sensorData SensorData, db *pg.DB) error {
	_, err := db.Model(&sensorData).Insert()
	if err != nil {
		return err
	}

	return nil
}

func GetAllSensors(db *pg.DB) (*[]SensorData, error) {
	// Select all users.
	var sensors []SensorData
	err := db.Model(&sensors).Select()
	if err != nil {
		return nil, err
	}

	return &sensors, nil
}

func GetRandomSensor(db *pg.DB) (*SensorData, error) {
	var sensor SensorData

	err := db.Model(&sensor).
		Limit(1).
		OrderExpr("random()").
		Select()

	if err != nil {
		return nil, err
	}

	return &sensor, nil
}

// createSchema creates database schema for SensorData
func createSchema(db *pg.DB) error {
	err := db.Model(&SensorData{}).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})

	if err != nil {
		return err
	}

	return nil
}
