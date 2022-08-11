package main

import (
	"context"
	"fmt"

	containeranalysis "cloud.google.com/go/containeranalysis/apiv1"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	grafeaspb "google.golang.org/genproto/googleapis/grafeas/v1"
)

var ctx = context.Background()

func (gcpApi *gcpApi) Login() {
	if gcpApi.credentials {
		gcpApi.validateAndGetClient(gcpApi.credentialsPath)
	} else {
		gcpApi.fetchCredentials()
	}
}

func (gcpApi *gcpApi) fetchCredentials() {
	var path string
	fmt.Println("Enter the credentials path downloaded to the local after creating the service account with container analysis role: ")
	fmt.Scanln(&path)
	callback := gcpApi.validateAndGetClient(path)
	if !callback {
		gcpApi.fetchCredentials()
	} else {
		return
	}
}

func (gcpApi *gcpApi) validateAndGetClient(path string) bool {

	client, err := containeranalysis.NewClient(ctx, option.WithCredentialsFile(path))
	if err != nil {
		fmt.Println(err)
		return false
	}
	gcpApi.client = client
	gcpApi.credentials = true
	gcpApi.credentialsPath = path
	return true
}

func (gcpApi *gcpApi) getVulnerability() {
	//get user input for image
	gcpApi.fetchResourceUrlAndProject()

	// create a grafeaspb occurence request
	req := &grafeaspb.ListOccurrencesRequest{
		Parent: fmt.Sprintf("projects/%s", gcpApi.projectID),
		Filter: fmt.Sprintf(`resourceUrl=%q`, gcpApi.resourceUrl),
	}

	//get array of occurences from the iterator
	gcpApi.getArrayFromRequest(req)

}


func (gcpApi *gcpApi) getArrayFromRequest(req *grafeaspb.ListOccurrencesRequest) {
	//fetch the occurence iterator
	it := gcpApi.client.GetGrafeasClient().ListOccurrences(ctx, req)
	fmt.Println("Entering the for loop")
	var count int

	for {
		occ, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		count++
		gcpApi.occs = append(gcpApi.occs, occ)
	}

	fmt.Println(count)

}

func (gcpApi *gcpApi) fetchResourceUrlAndProject() {
	var projectID, resourceUrl string
	fmt.Println("Enter the Project id: ")
	fmt.Scanln(&projectID)
	fmt.Println("Enter the Image url: ")
	fmt.Scanln(&resourceUrl)
	// validate function for projectid and resource Url

	gcpApi.projectID = projectID
	gcpApi.resourceUrl = resourceUrl
}


