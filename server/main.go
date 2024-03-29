package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	i "performance/server/initialize"
	"sort"
	"time"
)

// Response is a global var that is used to prevent the go compiler from optimizing my examples
var Response string

type response struct {
	uuid map[int]string `json:"uuid_id"`
}

func main() {

	// Get everything setup for the examples
	i.Init()
	sort.Strings(i.Match)
	<-time.After(time.Second * 1)

	// Examples for explicit program profiling
	//defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()

	//Response = buildResponse()
	//Response = buildResponseV2()
	//Response = buildResponseV3()

	// Example for running getting pprof profile on running server
	log.Println("Starting server")
	http.HandleFunc("/match", requestHandler)
	log.Fatal(http.ListenAndServe(":8086", nil))
}

// requestHandler is a trivial handler that always responds 200
func requestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", buildResponse())
}

// buildResponse uses nested loops to build a response object for the request handler.
// Disclaimer: I'm intentionally making this ridiculous.
// real    0m1.653s
// user    0m0.494s
// sys     0m0.018s
func buildResponse() string {

	// uuid will contain the response initialize that we'll marshal and return as a string
	uuid := make(map[int]string)
	// Loop for the size of the sync.Map
	for f := 0; f < i.Size; f++ {
		// Range over the []string uuids
		for _, value := range i.Match {
			// Hash into the sync.Map using the outer loop's var as its index
			v, ok := i.Lookup.Load(f)
			if ok {
				// Compare the outer loop's retrieved value against each uuid value of the []string
				if v == value {
					// Once found, add the value to uuid response map
					uuid[f] = value
				}
			}
		}
	}

	rs, _ := json.Marshal(uuid)
	return string(rs)
}

// buildResponseV2 is a little better, but not by much
// real    0m1.646s
// user    0m0.464s
// sys     0m0.019s
func buildResponseV2() string {

	// uuid will contain the response initialize that we'll marshal and return as a string
	uuid := make(map[int]string)
	// Loop for the size of the sync.Map
	for f := 0; f < i.Size; f++ {
		// Range over the []UUID
		for _, value := range i.Match {
			// Hash into the sync.Map using the outer loop's var as its index
			v, ok := i.Lookup.Load(f)
			if ok {
				// Compare the outer loop's retrieved value against each value of the []UUID
				if v == value {
					// Once found, add the value to uuid response map
					uuid[f] = value
					//break
				}
			}
		}
	}

	rs, _ := json.Marshal(uuid)
	return string(rs)
}

// buildResponseV3 is improved a bit more
// real    0m1.250s
// user    0m0.038s
// sys     0m0.018s
func buildResponseV3() string {

	// uuid will contain the response initialize that we'll marshal and return as a string
	uuid := make(map[int]string)

	for f := 0; f < i.Size; f++ {
		v, ok := i.Lookup.Load(f)
		if ok {
			s := convertString(v)
			if contains(i.Match, s) {
				// Once found, add the value to uuid response map
				uuid[f] = s
			}
		}
	}

	rs, _ := json.Marshal(uuid)
	return string(rs)
}

// contains is func I found online that uses sort.SearchStrings on a previously sorted []string
func contains(s []string, val string) bool {
	i := sort.SearchStrings(s, val)
	return i < len(s) && s[i] == val
}

// convertString is just a convenient way for me to convert sync.Map values that I know are strings into strings
func convertString(i interface{}) string {
	switch i.(type) {
	case string:
		return fmt.Sprintf("%v", i)
	default:
		return ""
	}
}
