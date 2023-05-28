package database

import (
	"database/sql"
	"fmt"
	"goGRAPH/graph/model"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	conn *sql.DB
}

func OppenConnection() (*sql.DB, error) {

	sc := ""

	conn, err := sql.Open("mysql", sc)

	if err != nil {
		panic(err)
	}
	err = conn.Ping()

	return conn, err
}

func GetEndpoint(endpointName string) *model.Endpoint {
	conn, err := OppenConnection()
	if err != nil {
		panic(fmt.Errorf(err.Error()))
	}
	defer conn.Close()

	var endpoint model.Endpoint

	id := GetEndpointID(endpointName)

	row := conn.QueryRow(fmt.Sprintf("SELECT EndpointID,EndpointName,Entries FROM count_log WHERE EndpointID = '%d'", id))

	row.Scan(&endpoint.EndpointID, &endpoint.EndpointName, &endpoint.Entries)

	return &endpoint
}

func GetEndpoints() []*model.Endpoint {
	conn, err := OppenConnection()
	if err != nil {
		panic(fmt.Errorf("Error"))
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT EndpointID,EndpointName,Entries FROM count_log`)

	if err != nil {
		panic(fmt.Errorf(err.Error()))
	}
	var endpoints []*model.Endpoint

	for rows.Next() {
		var endpoint model.Endpoint
		err = rows.Scan(&endpoint.EndpointID, &endpoint.EndpointName, &endpoint.Entries)
		if err != nil {
			continue
		}

		endpoints = append(endpoints, &endpoint)
	}

	return endpoints
}

func UpdateEndpoint(endpointName string) *model.Endpoint {
	conn, err := OppenConnection()
	if err != nil {
		panic(fmt.Errorf("Error"))
	}
	defer conn.Close()
	var updatedEndpoint model.Endpoint
	id := GetEndpointID(endpointName)
	res, _ := conn.Exec(fmt.Sprintf("UPDATE count_log SET Entries = count_log.entries + 1 WHERE EndpointID = '%d'", id))
	sucess, _ := res.RowsAffected()
	if sucess > 0 {
		row := conn.QueryRow(fmt.Sprintf("SELECT EndpointID,EndpointName,Entries FROM count_log WHERE EndpointID = '%d'", id))
		row.Scan(&updatedEndpoint.EndpointID, &updatedEndpoint.EndpointName, &updatedEndpoint.Entries)
	}
	return &updatedEndpoint
}

func GetEndpointID(name string) (EndpointID int32) {
	conn, err := OppenConnection()
	if err != nil {
		panic(fmt.Errorf(err.Error()))
	}
	defer conn.Close()

	row := conn.QueryRow(fmt.Sprintf("SELECT EndpointID FROM count_log WHERE EndpointName = '%s'", name))

	row.Scan(&EndpointID)

	return

}
