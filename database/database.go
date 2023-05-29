package database

import (
	"database/sql"
	"fmt"
	"goGRAPH/graph/model"

	_ "github.com/go-sql-driver/mysql"
)

const (
	_1 = "GET"
	_2 = "POST"
	_3 = "DELETE"
	_4 = "PUT"
)

func GetHttpMethodID(method string) int {
	switch method {
	case _1:
		return 1
	case _2:
		return 2
	case _3:
		return 3
	case _4:
		return 4
	}
	return 0
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

func GetEndpoint(endpointName string, http_method string) *model.Endpoint {
	conn, err := OppenConnection()
	if err != nil {
		panic(fmt.Errorf(err.Error()))
	}
	defer conn.Close()

	var endpoint model.Endpoint

	id := GetEndpointID(endpointName, http_method)

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

func UpdateEndpoint(endpointName string, http_method string) *model.Endpoint {
	conn, err := OppenConnection()
	if err != nil {
		panic(fmt.Errorf("Error"))
	}
	defer conn.Close()
	var updatedEndpoint model.Endpoint
	id := GetEndpointID(endpointName, http_method)
	res, _ := conn.Exec(fmt.Sprintf("UPDATE count_log SET Entries = count_log.entries + 1 WHERE EndpointID = '%d'", id))
	sucess, _ := res.RowsAffected()
	if sucess > 0 {
		row := conn.QueryRow(fmt.Sprintf("SELECT EndpointID,EndpointName,Entries FROM count_log WHERE EndpointID = '%d'", id))
		row.Scan(&updatedEndpoint.EndpointID, &updatedEndpoint.EndpointName, &updatedEndpoint.Entries)
	}
	return &updatedEndpoint
}

func GetEndpointID(name string, http_method string) (EndpointID int32) {
	conn, err := OppenConnection()
	if err != nil {
		panic(fmt.Errorf(err.Error()))
	}
	defer conn.Close()
	method_id := GetHttpMethodID(http_method)

	row := conn.QueryRow(fmt.Sprintf("SELECT EndpointID FROM count_log WHERE EndpointName = '%s' AND http_method = '%d'", name, method_id))

	row.Scan(&EndpointID)

	return
}
