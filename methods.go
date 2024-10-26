package main

import (
	"log"
	"signailai/initializers"
)

type Measurement struct {
	Value    int     `json:"value"`
	X        int     `json:"x"`
	Y        int     `json:"y"`
	Distance float32 `json:"distance"`
}

func GetMeasurements() ([]Measurement, error) {

	// Execute the query
	rows, err := initializers.DB.Query("SELECT * FROM matavimai")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	measurements := make([]Measurement, 0)

	for rows.Next() {
		// Define variables to store each column
		var value int
		var x, y int
		var distance float32

		// Scan the row into variables
		if err := rows.Scan(&value, &x, &y, &distance); err != nil {
			log.Fatal(err)
		}

		// Print the values
		measurements = append(measurements, Measurement{
			Distance: distance,
			X:        x,
			Y:        y,
			Value:    value,
		})
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return measurements, nil
}

type Strength struct {
	ID          int    `json:"id"`
	Measurement int    `json:"measurement"`
	Sensor      string `json:"sensor"`
	Strength    int    `json:"strength"`
}

func GetStrengths() ([]Strength, error) {

	// Execute the query
	rows, err := initializers.DB.Query("SELECT * FROM stiprumai")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	intensities := make([]Strength, 0)

	for rows.Next() {
		// Define variables to store each column
		var id, strength, measurement int
		var sensorName string

		// Scan the row into variables
		if err := rows.Scan(&id, &measurement, &sensorName, &strength); err != nil {
			log.Fatal(err)
		}

		// Print the values
		intensities = append(intensities, Strength{
			ID:          id,
			Measurement: measurement,
			Sensor:      sensorName,
			Strength:    strength,
		})
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return intensities, nil
}

type User struct {
	ID       int    `json:"id"`
	MAC      string `json:"mac"`
	Sensor   string `json:"sensor"`
	Strength int    `json:"strength"`
}

func GetUsers() ([]User, error) {

	// Execute the query
	rows, err := initializers.DB.Query("SELECT * FROM vartotojai")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		// Define variables to store each column
		var id, strength int
		var sensorName, mac string

		// Scan the row into variables
		if err := rows.Scan(&id, &mac, &sensorName, &strength); err != nil {
			log.Fatal(err)
		}

		// Print the values
		users = append(users, User{
			ID:       id,
			MAC:      mac,
			Sensor:   sensorName,
			Strength: strength,
		})
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return users, nil
}
