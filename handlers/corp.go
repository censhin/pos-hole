package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/censhin/pos-hole/client"
)

var eveClient *client.Client = client.NewClient()

type Response struct {
    Resp struct{}
}

func corpAccountBalanceHandler(w http.ResponseWriter, r *http.Request) {
    charId := r.URL.Query()

    resp, err := eveClient.CorpAccountBalance(charId["characterID"][0])
    if err != nil {
        log.Panic(err)
    }

    body := make(map[string]interface{})
    body["resp"] = resp

    b, err := json.Marshal(body)
    if err != nil {
        log.Panic(err)
    }

    w.Write(b)
}

