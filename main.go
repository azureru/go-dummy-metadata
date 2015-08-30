package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
    "strings"
    "sort"
)

var MetaDataMap = map[string]interface{}{
    "ami-id"                : "ami-5cdummy",
    "ami-launch-index"      : "0",
    "ami-manifest-path"     : "(unknown)",
    "block-device-mapping/" : "ami\nephemeral0\nroot",
    "hostname"              : "ip-10-10-10-101.ap-dummy-1.compute.internal",
    "instance-action"       : "none",
    "instance-id"           : "i-ccdummyfu",
    "instance-type"         : "t2.small",
    "kernel-id"             : "aki-aadummy",
    "local-hostname"        : "ip-10-10-10-101.ap-dummy-1.compute.internal",
    "local-ipv4"            : "10.10.10.101",
    "mac"                   : "12:31:FF:FF:FF:CC",
    "metrics/"              : "",
    "network/"              : "",
    "placement/"            : "",
    "profile"               : "default-paravirtual",
    "public-hostname"       : "ec2-44-144-144-144.ap-dummy-1.compute.amazonaws.com",
    "public-ipv4"           : "44.144.144.144",
    "public-keys/"          : "",
    "reservation-id"        : "r-8008135",
    "security-groups"       : "sc-group-default",
}

// Keys to get `keys` of map and put it in array
func Keys(m map[string]interface{}) (keys []string) {
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}

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
    lines := Keys(MetaDataMap)
    sort.Strings(lines)
    fmt.Fprint(w, strings.Join(lines, "\n"))
}

// MetaDataContent serve some meta-data values
func MetaDataContent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    id := ps.ByName("id")
    value := MetaDataMap[id]
    fmt.Fprintf(w, "%s\n", value)
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