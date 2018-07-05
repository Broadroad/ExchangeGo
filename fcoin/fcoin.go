package fcoin
import(
	"gtihub.com/exchangego/common"
	"net/http"
)

const (
	DEPTH_API       = "market/depth/%s/%s"
	TRADE_URL       = "orders"
	GET_ACCOUNT_API = "accounts/balance"
	GET_ORDER_API   = "orders/%s"
	GET_UNFINISHED_ORDERS_API = "getUnfinishedOrdersIgnoreTradeType"
	PLACE_ORDER_API           = "order"
	WITHDRAW_API              = "withdraw"
	CANCELWITHDRAW_API        = "cancelWithdraw"
	SERVER_TIME               = "public/server-time"
)

type FCoinTicker struct {
	Ticker
	SellAmount,
	BuyAmount float64
}

type FCoin struct {
	httpClient *http.Client
	baseUrl,
	accessKey,
	secretKey string
	timeoffset int64
}

func NewFCoin(client *http.Client, apikey, secretkey string) *FCoin {
	fc := &FCoin{baseUrl: "https://api.fcoin.com/v2/", accessKey: apikey, secretKey: secretkey, httpClient: client}
	fc.setTimeOffset()
	return fc
}

// setTimeOffset get server timestamp, because server and client time can not exceed 30 seconds
func (ft *FCoin) setTimeOffset() error {
	respmap, err := HttpGet(ft.httpClient, ft.baseUrl+"public/server-time")
	if err != nil {
		return err
	}
	stime := int64(ToInt(respmap["data"]))
	st := time.Unix(stime/1000, 0)
	lt := time.Now()
	offset := st.Sub(lt).Seconds()
	ft.timeoffset = int64(offset)
	return nil
}

// GetTicker get ticker data
func (ft *FCoin) GetTicker(currencyPair CurrencyPair) (*Ticker, error) {
	respmap, err := HttpGet(ft.httpClient, ft.baseUrl+fmt.Sprintf("market/ticker/%s",
		strings.ToLower(currencyPair.ToSymbol(""))))

	if err != nil {
		return nil, err
	}

	////log.Println("ticker respmap:", respmap)
	if respmap["status"].(float64) != 0 {
		return nil, errors.New(respmap["err-msg"].(string))
	}

	//
	tick, ok := respmap["data"].(map[string]interface{})
	if !ok {
		return nil, API_ERR
	}

	tickmap, ok := tick["ticker"].([]interface{})
	if !ok {
		return nil, API_ERR
	}

	ticker := new(Ticker)
	ticker.Pair = currencyPair
	ticker.Date = uint64(time.Now().Nanosecond() / 1000)
	ticker.LastAmount = ToFloat64(tickmap[1])
	ticker.Last24Amount = ToFloat64(tickmap[6])
	ticker.Last = ToFloat64(tickmap[0])
	ticker.Vol = ToFloat64(tickmap[9])
	ticker.Low = ToFloat64(tickmap[8])
	ticker.High = ToFloat64(tickmap[7])
	ticker.Buy = ToFloat64(tickmap[2])
	ticker.Sell = ToFloat64(tickmap[4])
	ticker.SellAmount = ToFloat64(tickmap[5])
	ticker.BuyAmount = ToFloat64(tickmap[3])

	return ticker, nil
}

