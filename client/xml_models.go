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

type StarbaseRow struct {
    TypeId   string `xml:"typeID,attr"`
    Quantity string `xml:"quantity,attr"`
}

type StarbaseOnStatusDrop struct {
    Enabled  string `xml:"enabled"`
    Standing string `xml:"standing"`
}

type StarbaseCombatSettings struct {
    UseStandingsFrom string               `xml:"useStandingsFrom>ownerID,attr"`
    OnStandingDrop   string               `xml:"onStandingDrop>standing,attr"`
    OnStatusDrop     StarbaseOnStatusDrop `xml:"onStatusDrop"`
    OnAggression     string               `xml:"onAggression>enabled,attr"`
    OnCorporationWar string               `xml:"onCorporationWar>enabled,attr"`
}

type StarbaseGeneralSettings struct {
    UsageFlags       int `xml:"usageFlags"`
    deployFlags      int `xml:"deployFlags"`
    AllowCorpMembers int `xml:"allowCorporationMembers"`
    AllowAllyMembers int `xml:"allowAllianceMembers"`
}

type StarbaseDetailResult struct {
    State           int                     `xml:"state"`
    StateTimestamp  eveTime                 `xml:"stateTimestamp"`
    OnlineTimestamp eveTime                 `xml:"onlineTimestamp"`
    GeneralSettings StarbaseGeneralSettings `xml:"generalSettings"`
    CombatSettings  StarbaseCombatSettings  `xml:"combatSettings"`
    Rowset          []StarbaseRow           `xml:"rowset>row"`
}

type StarbaseDetail struct {
    APIResult
    Result []StarbaseDetailResult `xml:"result"`
}

type Starbase struct {
    ItemId          string `xml:"itemID,attr"`
    TypeId          string `xml:"typeID,attr"`
    LocationId      string `xml:"locationID,attr"`
    MoonId          string `xml:"moonID,attr"`
    State           string `xml:"state,attr"`
    StateTimestamp  string `xml:"stateTimestamp,attr"`
    OnlineTimestamp string `xml:"onlineTimestamp,attr"`
    StandingOwnerId string `xml:"standingOwnerID,attr"`
}

type StarbaseList struct {
    APIResult
    Starbases []Starbase `xml:"result>rowset>row"`
}

