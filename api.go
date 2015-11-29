package main

// problems
//   geodns does not have an id for each zone

import (
	"fmt"
	"log"
	"net/http"
	"strings"
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

type Domain struct {
  Id          int       `json:"id"`
  Name        string    `json:"name"`
  Email       string    `json:"owner_email"`
  Type        int       `json:"type"`
  SubnetMask  int       `json:"subnet_mask"`
  DefaultNS1  string    `json:"default_ns1"`
  DefaultNS2  string    `json:"default_ns2"`    
}

type ApiResponse struct {
  Status	string    	`json:"status"`
  Id    	int    		`json:"id"`
  Error    	string    	`json:"error"`
}


func startApi(zones Zones) bool {

	// setup the API methods
	http.HandleFunc("/rapi/index", GetCurrentTime)

	http.HandleFunc("/rapi/getdomains", GetDomains(zones))
	// http.HandleFunc("/rapi/getdomain", GetDomain)
	http.HandleFunc("/rapi/getdomainbyname", GetDomainByName(zones))
	http.HandleFunc("/rapi/createregulardomain", CreateRegularDomain(zones))
	// http.HandleFunc("/rapi/createreversedomain", CreateReverseDomain)
	http.HandleFunc("/rapi/deletedomain", DeleteDomain(zones))
	http.HandleFunc("/rapi/updatedomain2", UpdateDomain2(zones))

	http.HandleFunc("/rapi/getrecords", GetRecords(zones))
	http.HandleFunc("/rapi/getrecord", GetRecord(zones))
	http.HandleFunc("/rapi/createrecord", CreateRecord(zones))
	http.HandleFunc("/rapi/updaterecord2", UpdateRecord2(zones))
	http.HandleFunc("/rapi/deleterecord", DeleteRecord(zones))

	// if no cert, use http
	//   flagapicert
	//   flagapiprivatekey

	log.Println("Starting HTTP interface on", *flagapiport)
	log.Fatal(http.ListenAndServe(*flagapiport, nil))

	return true
}

func GetCurrentTime( writer http.ResponseWriter, req *http.Request) {
	currentTime := time.Now().Format(time.RFC3339Nano)

	responseString :=  fmt.Sprintf( "{ \"utctime\" : \"%s\", \"version\" : \"%s\" }", currentTime, VERSION)
	writer.Write([]byte(responseString))
}

// make a list of zones/domains 
func GetDomains(zones Zones) func( writer http.ResponseWriter, req *http.Request) {

	return func(writer http.ResponseWriter, req *http.Request) {

		//array
		for name, zone := range zones {

			//create new Domain for each zone
			d := Domain{}  
			// d.id = id
			d.Name = name
			d.Email = zone.Options.Contact
			// "type" // Int16,
			// "subnet_mask" // Int16,
			// "default_ns1" // String,
			// "default_ns2" // String
		}

		responseString :=  fmt.Sprintf( "{  }")
		writer.Write([]byte(responseString))
	}
}

// 	//return domain
// 	responseString :=  fmt.Sprintf( "{  }")
// 	writer.Write([]byte(responseString))
// }

func GetDomain(zones Zones) func( writer http.ResponseWriter, req *http.Request) {

	return func(writer http.ResponseWriter, req *http.Request) {
		id := req.FormValue("id")
		log.Println("GetDomain( %s", id)

		responseString :=  fmt.Sprintf( "{  }")
		writer.Write([]byte(responseString))		
	}
}

func GetDomainByName(zones Zones) func( writer http.ResponseWriter, req *http.Request) {

	return func(writer http.ResponseWriter, req *http.Request) {

		name := req.FormValue("name")
		log.Println("GetDomainByName( %s", name)

		// go through zones and fine one for 'name'
		//array
		for domainName, zone := range zones {

			// see if this is the domain we want
			if (strings.EqualFold(domainName, name)) {
				//return domain
				d := Domain{}  
				// d.id = id
				d.Name = name
				d.Email = zone.Options.Contact

				responseString :=  fmt.Sprintf( "{  }")
				writer.Write([]byte(responseString))		
			}
		}

		// domain might not be found
	}
}

func CreateRegularDomain(zones Zones) func( writer http.ResponseWriter, req *http.Request) {

	return func(writer http.ResponseWriter, req *http.Request) {

		name := req.FormValue("name")
		email := req.FormValue("email")
		apiAccess := req.FormValue("apiaccess")
		ns1 := req.FormValue("ns1")
		ns2 := req.FormValue("ns2")
		log.Println("CreateRegularDomain( %s, %s, %s, %s, %s)", name, email, apiAccess, ns1, ns2)

		responseString :=  fmt.Sprintf( "{  }")
		writer.Write([]byte(responseString))
	}
}

// func CreateReverseDomain( writer http.ResponseWriter, req *http.Request) {
// 	responseString :=  fmt.Sprintf( "{  }")
// 	writer.Write([]byte(responseString))
// }

func DeleteDomain(zones Zones) func( writer http.ResponseWriter, req *http.Request) {

	return func(writer http.ResponseWriter, req *http.Request) {
		id := req.FormValue("id")
		log.Println("DeleteDomain( %s)", id)

		responseString :=  fmt.Sprintf( "{  }")
		writer.Write([]byte(responseString))
	}
}

func UpdateDomain2(zones Zones) func( writer http.ResponseWriter, req *http.Request) {

	return func(writer http.ResponseWriter, req *http.Request) {
		id := req.FormValue("id")
		email := req.FormValue("email")
		apiAccess := req.FormValue("apiaccess")
		ns1 := req.FormValue("ns1")
		ns2 := req.FormValue("ns2")
		log.Println("UpdateDomain2( %s, %s, %s, %s, %s)", id, email, apiAccess, ns1, ns2)

		responseString :=  fmt.Sprintf( "{  }")
		writer.Write([]byte(responseString))
	}
}

func GetRecords(zones Zones) func( writer http.ResponseWriter, req *http.Request) {

	return func(writer http.ResponseWriter, req *http.Request) {
		//id= domainId
		//name= record's name
		responseString :=  fmt.Sprintf( "{  }")
		writer.Write([]byte(responseString))
	}
}

func GetRecord(zones Zones) func( writer http.ResponseWriter, req *http.Request) {

	return func(writer http.ResponseWriter, req *http.Request) {
		//id= domainId
		responseString :=  fmt.Sprintf( "{  }")
		writer.Write([]byte(responseString))
	}
}

func CreateRecord( zones Zones) func( writer http.ResponseWriter, req *http.Request) {

	return func(writer http.ResponseWriter, req *http.Request) {
		//required: id, name, content, type, priority
		//optional: active, ttl, geozone, geolock, geolat, geolong, udplimit
		responseString :=  fmt.Sprintf( "{  }")
		writer.Write([]byte(responseString))
	}
}

func UpdateRecord2( zones Zones) func( writer http.ResponseWriter, req *http.Request) {

	return func(writer http.ResponseWriter, req *http.Request) {
		//required: id, name, content, type, priority
		//optional: active, ttl, geozone, geolock, geolat, geolong, udplimit
		responseString :=  fmt.Sprintf( "{  }")
		writer.Write([]byte(responseString))
	}
}

func DeleteRecord( zones Zones) func( writer http.ResponseWriter, req *http.Request) {

	return func(writer http.ResponseWriter, req *http.Request) {
		//id
		responseString :=  fmt.Sprintf( "{  }")
		writer.Write([]byte(responseString))
	}
}
