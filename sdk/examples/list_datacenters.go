package main

import (
	"fmt"
	"time"

	"../ovirtsdk4"
)

func main() {
	inputRawURL := "https://10.1.111.229/ovirt-engine/api"

	conn, err := ovirtsdk4.NewConnectionBuilder().
		URL(inputRawURL).
		Username("admin@internal").
		Password("qwer1234").
		Insecure(true).
		Compress(true).
		Timeout(time.Second * 10).
		Build()
	if err != nil {
		fmt.Printf("Make connection failed, reason: %s\n", err.Error())
	}
	defer conn.Close()

	// Get the reference to the "datacenters" service:
	datacentersService := conn.SystemService().DataCentersService()

	// Use the "list" method of the "datacenters" service to list all the datacenters of the system:
	datacentersResponse, err := datacentersService.List().Send()
	if err != nil {
		fmt.Printf("Failed to get datacenter list, reason: %v\n", err)
		return
	}

	// Print the datacenter names and identifiers:
	if datacenters, ok := datacentersResponse.DataCenters(); ok {
		for _, dc := range datacenters.Slice() {
			fmt.Printf("Datacenter: ")
			if dcName, ok := dc.Name(); ok {
				fmt.Printf(" name: %v", dcName)
			}
			if dcId, ok := dc.Id(); ok {
				fmt.Printf(" id: %v", dcId)
			}
			fmt.Printf("  Supported versions are: ")
			if svs, ok := dc.SupportedVersions(); ok {
				for _, sv := range svs.Slice() {
					if major, ok := sv.Major(); ok {
						fmt.Printf(" Major: %v", major)
					}
					if minor, ok := sv.Minor(); ok {
						fmt.Printf(" Minor: %v", minor)
					}
				}
			}
			fmt.Println("")
		}
	}
}
