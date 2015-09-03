package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strings"
)

// Index the metadata API root
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	lines := []string{"1.0", "2014-11-05", "2007-01-19"}
	fmt.Fprint(w, strings.Join(lines, "\n"))
}

// VersionRoot serve root
func VersionRoot(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	lines := []string{"dynamic", "meta-data"}
	fmt.Fprint(w, strings.Join(lines, "\n"))
}

// MetaDataContent serve some meta-data values
func MetaDataContent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	http.ServeFile(w, r, "./files/meta-data/"+id)
}

// UserDataContent serve user-data file to simulate AWS user-data request
func UserDataContent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.ServeFile(w, r, "./files/user-data")
}

func main() {
	router := httprouter.New()

	router.GET("/", Index)
	router.GET("/:version/", VersionRoot)
	router.GET("/:version/user-data", UserDataContent)
	router.GET("/:version/meta-data/:id", MetaDataContent)
	router.GET("/:version/meta-data/:id/", MetaDataContent)

	log.Fatal(http.ListenAndServe(":8080", router))
}
