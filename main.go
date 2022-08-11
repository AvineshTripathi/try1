package main

// import (
// 	"context"
// 	//"encoding/json"
// 	"fmt"

// 	containeranalysis "cloud.google.com/go/containeranalysis/apiv1"
// 	"google.golang.org/api/iterator"
// 	"google.golang.org/api/option"
// 	grafeaspb "google.golang.org/genproto/googleapis/grafeas/v1"
	//"google.golang.org/genproto/googleapis/devtools/containeranalysis/v1beta1/grafeas"
//)

// getOccurrencesForImage retrieves all the Occurrences associated with a specified image.
// Here, all Occurrences are simply printed and counted.
func main() {

	var g gcpApi
	g.Login()
	g.getVulnerability()

}




