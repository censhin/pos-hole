package client

type APIResult struct {
    Version     int       `xml:"version,attr"`
    CurrentTime eveTime   `xml:"currentTime"`
    Error       *APIError `xml:"error,omitempty"`
    CachedUntil eveTime   `xml:"cachedUntil"`
}

type APIError struct {
    Code    int    `xml:"code,attr"`
    Message string `xml:",chardata"`
}

type Account struct {
    AccountId  int     `xml:"accountID,attr"`
    AccountKey int     `xml:"accountKey,attr"`
    Balance    float64 `xml:"balance,attr"`
}

type AccountBalance struct {
    APIResult
    Result []Account `xml:"result>rowset>row"`
}

type Starbase struct {
    ItemId          string   `xml:"itemID,attr"`
    TypeId          string   `xml:"typeID,attr"`
    LocationId      string   `xml:"locationID,attr"`
    MoonId          string   `xml:"moonID,attr"`
    State           string   `xml:"state,attr"`
    StateTimestamp  string   `xml:"stateTimestamp,attr"`
    OnlineTimestamp string   `xml:"onlineTimestamp,attr"`
    StandingOwnerId string   `xml:"standingOwnerID,attr"`
}

type StarbaseList struct {
    APIResult
    Starbases []Starbase `xml:"result>rowset>row"`
}

