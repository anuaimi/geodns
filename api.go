package main

import (
	"fmt"
	"net/http"
	"time"
)

/*
   Copyright 2015 Athir Nuaimi

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

// need access to config

func startApi(zones Zones) bool {

	// flagapiport
	// flagapicert
		// if no cert, use http

	http.HandleFunc("/rapi/index", GetCurrentTime)

	http.HandleFunc("/rapi/getdomains", GetDomains)
	http.HandleFunc("/rapi/getdomain", GetDomain)
	http.HandleFunc("/rapi/getdomainbyname", GetDomainByName)
	http.HandleFunc("/rapi/getdomain", CreateRegularDomain)
	http.HandleFunc("/rapi/getdomain", CreateReverseDomain)
	http.HandleFunc("/rapi/getdomain", DeleteDomain)
	http.HandleFunc("/rapi/getdomain", Update2Domain)

	http.HandleFunc("/rapi/getrecords", GetRecords)
	http.HandleFunc("/rapi/getrecord", GetRecord)
	http.HandleFunc("/rapi/createrecord", CreateRecord)
	http.HandleFunc("/rapi/updaterecord2", UpdateRecord2)
	http.HandleFunc("/rapi/deleterecord", DeleteRecord)

	return true
}

func GetCurrentTime( writer http.ResponseWriter, req *http.Request) {
	currentTime := time.Now().Format(time.RFC3339Nano)

	responseString :=  fmt.Sprintf( "{ \"utctime\" : \"%s\", \"version\" : \"%s\" }", currentTime, VERSION)
	writer.Write([]byte(responseString))
}

func GetDomains( writer http.ResponseWriter, req *http.Request) {

	//array of 
		// {
		// "id" // Int32,
		// "name" // String,
		// "owner_email" // String,
		// "type" // Int16,
		// "subnet_mask" // Int16,
		// "default_ns1" // String,
		// "default_ns2" // String
		// }
	responseString :=  fmt.Sprintf( "{  }")
	writer.Write([]byte(responseString))
}

func GetDomain( writer http.ResponseWriter, req *http.Request) {
	//id
	//return domain
	responseString :=  fmt.Sprintf( "{  }")
	writer.Write([]byte(responseString))
}

func GetDomainByName( writer http.ResponseWriter, req *http.Request) {
	//name
	//return domain
	responseString :=  fmt.Sprintf( "{  }")
	writer.Write([]byte(responseString))
}

func CreateRegularDomain( writer http.ResponseWriter, req *http.Request) {
	responseString :=  fmt.Sprintf( "{  }")
	writer.Write([]byte(responseString))
}

func CreateReverseDomain( writer http.ResponseWriter, req *http.Request) {
	responseString :=  fmt.Sprintf( "{  }")
	writer.Write([]byte(responseString))
}

func DeleteDomain( writer http.ResponseWriter, req *http.Request) {
	responseString :=  fmt.Sprintf( "{  }")
	writer.Write([]byte(responseString))
}

func Update2Domain( writer http.ResponseWriter, req *http.Request) {
	responseString :=  fmt.Sprintf( "{  }")
	writer.Write([]byte(responseString))
}

func GetRecords( writer http.ResponseWriter, req *http.Request) {
	//id= domainId
	//name= record's name
	responseString :=  fmt.Sprintf( "{  }")
	writer.Write([]byte(responseString))
}

func GetRecord( writer http.ResponseWriter, req *http.Request) {
	//id= domainId
	responseString :=  fmt.Sprintf( "{  }")
	writer.Write([]byte(responseString))
}

func CreateRecord( writer http.ResponseWriter, req *http.Request) {
	//required: id, name, content, type, priority
	//optional: active, ttl, geozone, geolock, geolat, geolong, udplimit
	responseString :=  fmt.Sprintf( "{  }")
	writer.Write([]byte(responseString))
}

func UpdateRecord2( writer http.ResponseWriter, req *http.Request) {
	//required: id, name, content, type, priority
	//optional: active, ttl, geozone, geolock, geolat, geolong, udplimit
	responseString :=  fmt.Sprintf( "{  }")
	writer.Write([]byte(responseString))
}

func DeleteRecord( writer http.ResponseWriter, req *http.Request) {
	//id
	responseString :=  fmt.Sprintf( "{  }")
	writer.Write([]byte(responseString))
}
