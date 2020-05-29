package main

import (
	"github.com/brianvoe/gofakeit"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

// RandResourceURI generates a random resource URI
func RandResourceURI() string {
	var uri string
	num := gofakeit.Number(1, 4)
	for i := 0; i < num; i++ {
		uri += "/" + url.QueryEscape(gofakeit.BS())
	}
	uri = strings.ToLower(uri)
	return uri
}

// RandAuthUserID generates a random auth user id
func RandAuthUserID() string {
	candidates := []string{"-", strings.ToLower(gofakeit.Username())}
	return candidates[rand.Intn(2)]
}

// RandHTTPVersion returns a random http version
func RandHTTPVersion() string {
	versions := []string{"HTTP/1.0", "HTTP/1.1", "HTTP/2.0"}
	return versions[rand.Intn(3)]
}

// RandPackageName returns a random package name
func RandPackageName() string {

	// PackageNames consists of Citrus package names
	/*var PackageNames = map[string][]string{
		"citrus": {"infrastructure.common", "utils.FacterInfo"},
	}*/

	PackageNames := []string{
		"infrastructure.common",
		"utils.FacterInfo",
	}
	rand.Seed(time.Now().UnixNano())
	return PackageNames[rand.Intn(2)]
}
