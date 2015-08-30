package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
    "strings"
)

// Index the metadata API root
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    lines := []string{"1.0","2014-11-05","2007-01-19"}
    fmt.Fprint(w, strings.Join(lines, "\n"))
}

// VersionRoot serve root
func VersionRoot(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    lines := []string{"dynamic","metadata"}
    fmt.Fprint(w, strings.Join(lines, "\n"))
}

// MetaDataRoot list whole meta-data keys
func MetaDataRoot(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    lines := []string{
        "ami-id",
        "ami-launch-index",
        "ami-manifest-path",
        "block-device-mapping/",
        "hostname",
        "instance-action",
        "instance-id",
        "instance-type",
        "kernel-id",
        "local-hostname",
        "local-ipv4",
        "mac",
        "metrics/",
        "network/",
        "placement/",
        "profile",
        "public-hostname",
        "public-ipv4",
        "public-keys/",
        "reservation-id",
        "security-groups",
    }
    fmt.Fprint(w, strings.Join(lines, "\n"))
}

// MetaDataContent serve some meta-data values
func MetaDataContent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    version := ps.ByName("version")
    id := ps.ByName("id");
    fmt.Fprintf(w, "hello, %s %s!\n", version, id)
}

// UserDataContent serve user-data file to simulate AWS user-data request
func UserDataContent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    http.ServeFile(w, r, "./files/user-data")
}

func main() {
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/:version/", VersionRoot)
    router.GET("/:version/meta-data/", MetaDataRoot)
    router.GET("/:version/meta-data/:id", MetaDataContent)
    router.GET("/:version/user-data", UserDataContent)

    log.Fatal(http.ListenAndServe(":8080", router))
}