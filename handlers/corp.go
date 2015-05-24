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

func writeResponse(resp interface{}, w http.ResponseWriter) {
    body := make(map[string]interface{})
    body["resp"] = resp

    b, err := json.Marshal(body)
    if err != nil {
        log.Panic(err)
    }

    w.Write(b)
}

func corpAccountBalanceHandler(w http.ResponseWriter, r *http.Request) {
    charId := r.URL.Query()

    resp, err := eveClient.CorpAccountBalance(charId["characterID"][0])
    if err != nil {
        log.Panic(err)
    }

    writeResponse(resp, w)
}

func corpAssetListHandler(w http.ResponseWriter, r *http.Request) {
    charId := r.URL.Query()

    resp, err := eveClient.CorpAssetList(charId["characterID"][0])
    if err != nil {
        log.Panic(err)
    }

    writeResponse(resp, w)
}

func corpStarbaseDetailHandler(w http.ResponseWriter, r *http.Request) {
    itemId := r.URL.Query()

    resp, err := eveClient.StarbaseDetail(itemId["itemID"][0])
    if err != nil {
        log.Panic(err)
    }

    writeResponse(resp, w)
}

func corpStarbaseListHandler(w http.ResponseWriter, r *http.Request) {
    resp, err := eveClient.StarbaseList()
    if err != nil {
        log.Panic(err)
    }

    writeResponse(resp, w)
}


