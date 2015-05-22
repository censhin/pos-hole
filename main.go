package main

import (
    "fmt"
    "log"

    "github.com/censhin/pos-hole/client"
)

func main() {
    baseUrl := "https://api.eveonline.com"
    keyId := "4394542"
    vCode := "rIgfd5ofkJJm7m9AzNI4Pq5FLQaibOUPM3n9J4teCiMBAidgiFOhzaEFaq32MY5p"

    client := client.NewClient(baseUrl, keyId, vCode)
    resp, err := client.CorpAccountBalance("1853330359")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(resp)
}

