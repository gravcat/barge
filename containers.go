package main

import (
	"encoding/json"
	"strconv"
)

// Container is a Docker container
type Container struct {
	ID    string
	Image string
	State string
}

func (p Portainer) getContainersForEndpoint(endpoint Endpoint) []Container {
	output := p.fetch("endpoints/" + strconv.Itoa(endpoint.ID) + "/docker/containers/json")

	containers := make([]Container, 0)

	json.Unmarshal([]byte(output), &containers)

	return containers
}

func (p Portainer) populateContainersForEndpoints(endpoints []Endpoint) []Endpoint {
	newEndpoints := []Endpoint{}
	var endpoint Endpoint

	for _, e := range endpoints {
		endpoint = e
		endpoint.Containers = p.getContainersForEndpoint(e)

		newEndpoints = append(newEndpoints, endpoint)
	}

	return newEndpoints
}
