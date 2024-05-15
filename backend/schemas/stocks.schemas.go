package schemas

type SearchStocksSchema struct {
	BestMatches []StockSchema `json:"bestMatches"`
}

type StockSchema struct {
	Symbol      string `json:"1. symbol"`
	Name        string `json:"2. name"`
	Type        string `json:"3. type"`
	Region      string `json:"4. region"`
	MarketOpen  string `json:"5. marketOpen"`
	MarketClose string `json:"6. marketClose"`
	Timezone    string `json:"7. timezone"`
	Currency    string `json:"8. currency"`
	MatchScore  string `json:"9. matchScore"`
}

type TimeData struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}

type StockMetaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	Inteval       string `json:"4. Interval"`
	OutputSize    string `json:"5. Output Size"`
	TimeZone      string `json:"6. Time Zone"`
}

type GetStocksResponseSchema struct {
	MetaData      StockMetaData       `json:"Meta Data"`
	TimeData1min  map[string]TimeData `json:"Time Series (1min)"`
	TimeData5min  map[string]TimeData `json:"Time Series (5min)"`
	TimeData15min map[string]TimeData `json:"Time Series (15min)"`
	TimeData30min map[string]TimeData `json:"Time Series (30min)"`
	TimeData60min map[string]TimeData `json:"Time Series (60min)"`
}
