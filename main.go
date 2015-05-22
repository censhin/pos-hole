package main

import (
    "fmt"
    "log"

    "github.com/censhin/pos-hole/client"
)

func main() {
    baseUrl := "https://api.eveonline.com"
    keyId := ""
    vCode := ""

    client := client.NewClient(baseUrl, keyId, vCode)
    resp, err := client.CorpAccountBalance("")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(resp)
}

