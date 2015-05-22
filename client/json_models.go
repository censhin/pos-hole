package client

type Account struct {
    AccountId  int
    AccountKey int
    Balance    float64
}

type AccountBalance struct {
    Accounts []Account
}

type Starbase struct {
    ItemId          string
    TypeId          string
    LocationId      string
    MoonId          string
    State           string
    StateTimestamp  string
    OnlineTimestamp string
    StandingOwnerId string
}

type StarbaseList struct {
    Starbases []Starbase
}

