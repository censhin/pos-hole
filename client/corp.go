package client

import (
    "fmt"
    "log"
)

func (client *Client) CorpAccountBalance(characterId string) (*AccountBalance, error) {
    resource := "/corp/AccountBalance.xml.aspx"
    //params := fmt.Sprintf("&characterID=%v", characterId)
    output := AccountBalance{}

    err := ApiV2Request(client, resource, "", &output)
    if err != nil {
        log.Panic(err)
        return nil, err
    }

    return &output, nil
}

func (client *Client) StarbaseDetail(characterId string, itemId string) (*Starbase, error) {
    resource := "/corp/StarbaseDetail.xml.aspx"
    params := fmt.Sprintf("&characterId=%v&itemID=%v",characterId, itemId)
    output := Starbase{}

    err := ApiV2Request(client, resource, params, &output)
    if err != nil {
        log.Panic(err)
        return nil, err
    }

    return &output, nil
}

func (client *Client) StarbaseList() (*StarbaseList, error) {
    resource := "/corp/StarbaseList.xml.aspx"
    output := StarbaseList{}

    err := ApiV2Request(client, resource, "", &output)
    if err != nil {
        log.Panic(err)
        return nil, err
    }

    return &output, nil
}

