package main

import (
	containeranalysis "cloud.google.com/go/containeranalysis/apiv1"
	grafeaspb "google.golang.org/genproto/googleapis/grafeas/v1"
)

type gcpApi struct {
	clientId        string
	credentialsPath string
	client          *containeranalysis.Client
	resourceUrl     string
	projectID       string
	occs            []*grafeaspb.Occurrence
	credentials     bool
}
