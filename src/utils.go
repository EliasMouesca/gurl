package main

import (
    "fmt"
    "net"
    "errors"
    "os"
    "syscall"
    "net/http"
)

func die(str string) {
    fmt.Fprintln(os.Stderr, str)
    os.Exit(1)
}

func handleRequestErrors(request *http.Request, err error) {
    host := request.URL.Host

    if (errors.Is(err, syscall.ECONNREFUSED)) {
        die("Connection refused at '"+host+"'")
    } 

    var dnsError *net.DNSError
    if errors.As(err, &dnsError) {
        die("Could not resolve '"+host+"'")
    }

    die(err.Error())

}

