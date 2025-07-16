package mexcsdk

import (
	"fmt"

	"github.com/farbodghasemlu/mexc-sdk-go/request"
)

// SymbolOrderBookTicker represents symbol order book ticker
type SymbolOrderBookTicker struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
}

// SymbolPriceTicker represents symbol price ticker
type SymbolPriceTicker struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

// Ticker24hr represents 24hr ticker price change statistics
type Ticker24hr struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	BidPrice           string `json:"bidPrice"`
	BidQty             string `json:"bidQty"`
	AskPrice           string `json:"askPrice"`
	AskQty             string `json:"askQty"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	Count              string `json:"count"`
}

// CurrentAveragePrice represents current average price
type CurrentAveragePrice struct {
	Mins  int    `json:"mins"`
	Price string `json:"price"`
}

// KlineData represents kline/candlestick data
type KlineData []interface{}

// AggregateTrade represents compressed/aggregate trades
type AggregateTrade struct {
	A  interface{} `json:"a"` // Aggregate trade ID
	F  interface{} `json:"f"` // First trade ID
	L  interface{} `json:"l"` // Last trade ID
	P  string      `json:"p"` // Price
	Q  string      `json:"q"` // Quantity
	T  int64       `json:"T"` // Timestamp
	M  bool        `json:"m"` // Was the buyer the maker?
	M2 bool        `json:"M"` // Was the trade the best price match?
}

// RecentTrade represents recent trades
type RecentTrade struct {
	ID           interface{} `json:"id"`
	Price        string      `json:"price"`
	Qty          string      `json:"qty"`
	QuoteQty     string      `json:"quoteQty"`
	Time         int64       `json:"time"`
	IsBuyerMaker bool        `json:"isBuyerMaker"`
	IsBestMatch  bool        `json:"isBestMatch"`
}

// OrderBook represents order book data
type OrderBook struct {
	LastUpdateID int64      `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

// ExchangeInfo represents exchange information
type ExchangeInfo struct {
	Symbol                     string        `json:"symbol"`
	Status                     string        `json:"status"`
	BaseAsset                  string        `json:"baseAsset"`
	BaseAssetPrecision         int           `json:"baseAssetPrecision"`
	QuoteAsset                 string        `json:"quoteAsset"`
	QuotePrecision             int           `json:"quotePrecision"`
	QuoteAssetPrecision        int           `json:"quoteAssetPrecision"`
	BaseCommissionPrecision    int           `json:"baseCommissionPrecision"`
	QuoteCommissionPrecision   int           `json:"quoteCommissionPrecision"`
	OrderTypes                 []string      `json:"orderTypes"`
	QuoteOrderQtyMarketAllowed bool          `json:"quoteOrderQtyMarketAllowed"`
	IsSpotTradingAllowed       bool          `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed     bool          `json:"isMarginTradingAllowed"`
	QuoteAmountPrecision       string        `json:"quoteAmountPrecision"`
	BaseSizePrecision          string        `json:"baseSizePrecision"`
	Permissions                []string      `json:"permissions"`
	Filters                    []interface{} `json:"filters"`
	MaxQuoteAmount             string        `json:"maxQuoteAmount"`
	MakerCommission            string        `json:"makerCommission"`
	TakerCommission            string        `json:"takerCommission"`
	TradeSideType              string        `json:"tradeSideType"`
}

// ServerTime represents server time response
type ServerTime struct {
	ServerTime int64 `json:"serverTime"`
}

// SubAccountAsset represents sub-account asset
type SubAccountAsset struct {
	Balances []Balance `json:"balances"`
}

// MasterUniversalTransferHistory represents universal transfer history for master account
type MasterUniversalTransferHistory struct {
	TranID          string `json:"tranId"`
	FromAccount     string `json:"fromAccount"`
	ToAccount       string `json:"toAccount"`
	ClientTranID    string `json:"clientTranId"`
	Asset           string `json:"asset"`
	Amount          string `json:"amount"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
	FromSymbol      string `json:"fromSymbol"`
	ToSymbol        string `json:"toSymbol"`
	Status          string `json:"status"`
	Timestamp       int64  `json:"timestamp"`
}

// MasterUniversalTransferResponse represents universal transfer response for master account
type MasterUniversalTransferResponse struct {
	TranID int64 `json:"tranId"`
}

// DeleteSubAccountAPIKeyResponse represents response for deleting sub-account API key
type DeleteSubAccountAPIKeyResponse struct {
	SubAccount string `json:"subAccount"`
}

// SubAccountAPIKey represents sub-account API key information
type SubAccountAPIKey struct {
	Note        string `json:"note"`
	APIKey      string `json:"apiKey"`
	Permissions string `json:"permissions"`
	IP          string `json:"ip"`
	CreatTime   int64  `json:"creatTime"`
}

// SubAccountAPIKeyResponse represents response for sub-account API key query
type SubAccountAPIKeyResponse struct {
	SubAccount []SubAccountAPIKey `json:"subAccount"`
}

// CreateSubAccountAPIKeyResponse represents response for creating sub-account API key
type CreateSubAccountAPIKeyResponse struct {
	SubAccount  string `json:"subAccount"`
	Note        string `json:"note"`
	APIKey      string `json:"apiKey"`
	SecretKey   string `json:"secretKey"`
	Permissions string `json:"permissions"`
	IP          string `json:"ip"`
	CreatTime   int64  `json:"creatTime"`
}

// SubAccountList represents sub-account list
type SubAccountList struct {
	SubAccounts []SubAccountInfo `json:"subAccounts"`
}

type SubAccountInfo struct {
	SubAccount string `json:"subAccount"`
	IsFreeze   bool   `json:"isFreeze"`
	CreateTime int64  `json:"createTime"`
	UID        string `json:"uid"`
}

// CreateSubAccountResponse represents response for creating sub-account
type CreateSubAccountResponse struct {
	SubAccount string `json:"subAccount"`
	Note       string `json:"note"`
}

// Trade represents account trade information
type Trade struct {
	Symbol          string `json:"symbol"`
	ID              string `json:"id"`
	OrderID         string `json:"orderId"`
	OrderListID     int64  `json:"orderListId"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	QuoteQty        string `json:"quoteQty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            int64  `json:"time"`
	IsBuyer         bool   `json:"isBuyer"`
	IsMaker         bool   `json:"isMaker"`
	IsBestMatch     bool   `json:"isBestMatch"`
	IsSelfTrade     bool   `json:"isSelfTrade"`
	ClientOrderID   string `json:"clientOrderId"`
}

// AccountInfo represents account information
type AccountInfo struct {
	CanTrade    bool      `json:"canTrade"`
	CanWithdraw bool      `json:"canWithdraw"`
	CanDeposit  bool      `json:"canDeposit"`
	UpdateTime  *int64    `json:"updateTime"`
	AccountType string    `json:"accountType"`
	Balances    []Balance `json:"balances"`
	Permissions []string  `json:"permissions"`
}

type Balance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

// Order represents order information
type Order struct {
	Symbol              string `json:"symbol"`
	OrderID             int64  `json:"orderId"`
	OrderListID         int64  `json:"orderListId"`
	ClientOrderID       string `json:"clientOrderId"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
	StopPrice           string `json:"stopPrice"`
	IcebergQty          string `json:"icebergQty"`
	Time                int64  `json:"time"`
	UpdateTime          int64  `json:"updateTime"`
	IsWorking           bool   `json:"isWorking"`
	OrigQuoteOrderQty   string `json:"origQuoteOrderQty"`
}

// CancelOrderResponse represents cancel order response
type CancelOrderResponse struct {
	Symbol              string `json:"symbol"`
	OrigClientOrderID   string `json:"origClientOrderId"`
	OrderID             int64  `json:"orderId"`
	ClientOrderID       string `json:"clientOrderId"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
}

// BatchOrderResponse represents batch order response
type BatchOrderResponse struct {
	Success []BatchOrderSuccess `json:"success"`
	Error   []BatchOrderError   `json:"error"`
}

type BatchOrderSuccess struct {
	Symbol      string `json:"symbol"`
	OrderID     string `json:"orderId"`
	OrderListID int64  `json:"orderListId"`
}

type BatchOrderError struct {
	Symbol           string `json:"symbol,omitempty"`
	OrderID          string `json:"orderId,omitempty"`
	NewClientOrderID string `json:"newClientOrderId,omitempty"`
	OrderListID      int64  `json:"orderListId,omitempty"`
	Msg              string `json:"msg"`
	Code             int    `json:"code"`
}

// NewOrderResponse represents new order response
type NewOrderResponse struct {
	Symbol       string `json:"symbol"`
	OrderID      string `json:"orderId"`
	OrderListID  int64  `json:"orderListId"`
	Price        string `json:"price"`
	OrigQty      string `json:"origQty"`
	Type         string `json:"type"`
	Side         string `json:"side"`
	TransactTime int64  `json:"transactTime"`
}

// DefaultSymbolResponse represents default symbol response
type DefaultSymbolResponse struct {
	Code int      `json:"code"`
	Data []string `json:"data"`
	Msg  *string  `json:"msg"`
}

// KYCStatus represents KYC status
type KYCStatus struct {
	Status string `json:"status"`
}

// WithdrawResponse represents response for withdraw requests
type WithdrawResponse struct {
	ID string `json:"id"`
}

// InternalTransferHistory represents internal transfer history data
type InternalTransferHistory struct {
	Page         int                      `json:"page"`
	TotalRecords int                      `json:"totalRecords"`
	TotalPageNum int                      `json:"totalPageNum"`
	Data         []InternalTransferRecord `json:"data"`
}

type InternalTransferRecord struct {
	TranID        string `json:"tranId"`
	Asset         string `json:"asset"`
	Amount        string `json:"amount"`
	ToAccountType string `json:"toAccountType"`
	ToAccount     string `json:"toAccount"`
	FromAccount   string `json:"fromAccount"`
	Status        string `json:"status"`
	Timestamp     int64  `json:"timestamp"`
}

// InternalTransferResponse represents response for internal transfer
type InternalTransferResponse struct {
	TranID string `json:"tranId"`
}

// DustLog represents dust conversion log
type DustLog struct {
	TotalRecords int           `json:"totalRecords"`
	Page         int           `json:"page"`
	TotalPageNum int           `json:"totalPageNum"`
	Data         []DustLogItem `json:"data"`
}

type DustLogItem struct {
	TotalConvert   string          `json:"totalConvert"`
	TotalFee       string          `json:"totalFee"`
	ConvertTime    int64           `json:"convertTime"`
	ConvertDetails []ConvertDetail `json:"convertDetails"`
}

type ConvertDetail struct {
	ID      string `json:"id"`
	Convert string `json:"convert"`
	Fee     string `json:"fee"`
	Amount  string `json:"amount"`
	Time    int64  `json:"time"`
	Asset   string `json:"asset"`
}

// DustTransferResponse represents response for dust transfer
type DustTransferResponse struct {
	SuccessList  []string `json:"successList"`
	FailedList   []string `json:"failedList"`
	TotalConvert string   `json:"totalConvert"`
	ConvertFee   string   `json:"convertFee"`
}

// ConvertibleAsset represents assets that can be converted to MX
type ConvertibleAsset struct {
	ConvertMx   string `json:"convertMx"`
	ConvertUsdt string `json:"convertUsdt"`
	Balance     string `json:"balance"`
	Asset       string `json:"asset"`
	Code        string `json:"code"`
	Message     string `json:"message"`
}

// UniversalTransferHistory represents universal transfer history
type UniversalTransferHistory struct {
	TranID          string `json:"tranId"`
	ClientTranID    string `json:"clientTranId"`
	Asset           string `json:"asset"`
	Amount          string `json:"amount"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
	Symbol          string `json:"symbol"`
	Status          string `json:"status"`
	Timestamp       int64  `json:"timestamp"`
}

// UniversalTransferHistoryList represents list of universal transfer histories
type UniversalTransferHistoryList struct {
	Rows  []UniversalTransferHistory `json:"rows"`
	Total int                        `json:"total"`
}

// UniversalTransferResponse represents response for universal transfer
type UniversalTransferResponse struct {
	TranID string `json:"tranId"`
}

// WithdrawAddress represents withdraw address information
type WithdrawAddress struct {
	Data         []AddressInfo `json:"data"`
	TotalRecords int           `json:"totalRecords"`
	Page         int           `json:"page"`
	TotalPageNum int           `json:"totalPageNum"`
}

type AddressInfo struct {
	Coin       string `json:"coin"`
	Network    string `json:"network"`
	Address    string `json:"address"`
	AddressTag string `json:"addressTag"`
	Memo       string `json:"memo"`
}

// DepositAddress represents deposit address information
type DepositAddress struct {
	Coin    string `json:"coin"`
	Network string `json:"network"`
	Address string `json:"address"`
	Memo    string `json:"memo"`
}

// GeneratedDepositAddress represents generated deposit address
type GeneratedDepositAddress struct {
	Coin    string `json:"coin"`
	Network string `json:"network"`
	Address string `json:"address"`
	Tag     string `json:"tag"`
	Memo    string `json:"memo"`
}

// WithdrawHistory represents withdraw history
type WithdrawHistory struct {
	ID             string `json:"id"`
	TxID           string `json:"txId"`
	Coin           string `json:"coin"`
	Network        string `json:"network"`
	Address        string `json:"address"`
	Amount         string `json:"amount"`
	TransferType   int    `json:"transferType"`
	Status         int    `json:"status"`
	TransactionFee string `json:"transactionFee"`
	ConfirmNo      string `json:"confirmNo"`
	ApplyTime      int64  `json:"applyTime"`
	Remark         string `json:"remark"`
	Memo           string `json:"memo"`
	TransHash      string `json:"transHash"`
	UpdateTime     int64  `json:"updateTime"`
	CoinID         string `json:"coinId"`
	VcoinID        string `json:"vcoinId"`
}

// DepositHistory represents deposit history
type DepositHistory struct {
	Amount        string `json:"amount"`
	Coin          string `json:"coin"`
	Network       string `json:"network"`
	Status        int    `json:"status"`
	Address       string `json:"address"`
	TxID          string `json:"txId"`
	InsertTime    int64  `json:"insertTime"`
	UnlockConfirm string `json:"unlockConfirm"`
	ConfirmTimes  string `json:"confirmTimes"`
	Memo          string `json:"memo"`
}

// CancelWithdrawResponse represents response for cancel withdraw
type CancelWithdrawResponse struct {
	ID string `json:"id"`
}

// CurrencyInfo represents currency information
type CurrencyInfo struct {
	Coin        string        `json:"coin"`
	Name        string        `json:"Name"`
	NetworkList []NetworkInfo `json:"networkList"`
}

type NetworkInfo struct {
	Coin                    string  `json:"coin"`
	DepositDesc             *string `json:"depositDesc"`
	DepositEnable           bool    `json:"depositEnable"`
	MinConfirm              int     `json:"minConfirm"`
	Name                    string  `json:"Name"`
	Network                 string  `json:"network"`
	WithdrawEnable          bool    `json:"withdrawEnable"`
	WithdrawFee             string  `json:"withdrawFee"`
	WithdrawIntegerMultiple *string `json:"withdrawIntegerMultiple"`
	WithdrawMax             string  `json:"withdrawMax"`
	WithdrawMin             string  `json:"withdrawMin"`
	SameAddress             bool    `json:"sameAddress"`
	Contract                string  `json:"contract"`
	WithdrawTips            string  `json:"withdrawTips"`
	DepositTips             string  `json:"depositTips"`
	NetWork                 string  `json:"netWork"`
}

// MXDeductStatus represents MX deduct status
type MXDeductStatus struct {
	Data struct {
		MXDeductEnable bool `json:"mxDeductEnable"`
	} `json:"data"`
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Timestamp int64  `json:"timestamp"`
}
type Spot interface {
	Ping(jsonParams string) interface{}
	Time(jsonParams string) *ServerTime
	ApiSymbol(jsonParams string) *DefaultSymbolResponse
	ExchangeInfo(jsonParams string) *ExchangeInfo
	Depth(jsonParams string) *OrderBook
	Trades(jsonParams string) []RecentTrade
	AggTrades(jsonParams string) []AggregateTrade
	Kline(jsonParams string) []KlineData
	AvgPrice(jsonParams string) *CurrentAveragePrice
	Ticker24hr(jsonParams string) []Ticker24hr
	Price(jsonParams string) []SymbolPriceTicker
	BookTicker(jsonParams string) []SymbolOrderBookTicker

	CreateSub(jsonParams string) *CreateSubAccountResponse
	QuerySub(jsonParams string) *SubAccountList
	CreateSubApikey(jsonParams string) *CreateSubAccountAPIKeyResponse
	QuerySubApikey(jsonParams string) *SubAccountAPIKeyResponse
	DeleteSubApikey(jsonParams string) *DeleteSubAccountAPIKeyResponse
	UniTransfer(jsonParams string) *MasterUniversalTransferResponse
	QueryUniTransfer(jsonParams string) *MasterUniversalTransferHistory

	SelfSymbols(jsonParams string) *DefaultSymbolResponse
	TestOrder(jsonParams string) interface{}
	PlaceOrder(jsonParams string) *NewOrderResponse
	BatchOrder(jsonParams string) *BatchOrderResponse
	CancelOrder(jsonParams string) *CancelOrderResponse
	CancelAllOrders(jsonParams string) []CancelOrderResponse
	QueryOrder(jsonParams string) *Order
	OpenOrder(jsonParams string) []Order
	AllOrders(jsonParams string) []Order
	SpotAccountInfo(jsonParams string) *AccountInfo
	SpotmyTrade(jsonParams string) []Trade
	MxDeduct(jsonParams string) *MXDeductStatus
	QueryMxDeduct(jsonParams string) *MXDeductStatus

	QueryCurrencyInfo(jsonParams string) []CurrencyInfo
	Withdraw(jsonParams string) *WithdrawResponse
	CancelWithdraw(jsonParams string) *CancelWithdrawResponse
	DepositHistory(jsonParams string) []DepositHistory
	WithdrawHistory(jsonParams string) []WithdrawHistory
	GenDepositAddress(jsonParams string) []GeneratedDepositAddress
	DepositAddress(jsonParams string) []DepositAddress
	WithdrawAddress(jsonParams string) *WithdrawAddress
	Transfer(jsonParams string) *InternalTransferResponse
	TransferHistory(jsonParams string) *InternalTransferHistory
	TransferHistoryById(jsonParams string) *UniversalTransferHistory
	ConvertList(jsonParams string) []ConvertibleAsset
	Convert(jsonParams string) *DustTransferResponse
	ConvertHistory(jsonParams string) *DustLog
	ETFInfo(jsonParams string) interface{} // No type provided in docs
	InternalTransfer(jsonParams string) *InternalTransferResponse
	InternalTransferHistory(jsonParams string) *InternalTransferHistory

	CreateListenKey(jsonParams string) interface{} // No type provided in docs
	KeepListenKey(jsonParams string) interface{}   // No type provided in docs
	CloseListenKey(jsonParams string) interface{}  // No type provided in docs

	RebateHistory(jsonParams string) interface{}             // No type provided in docs
	RebateDetail(jsonParams string) interface{}              // No type provided in docs
	SelfRecordsDetail(jsonParams string) interface{}         // No type provided in docs
	ReferCode(jsonParams string) interface{}                 // No type provided in docs
	AffiliateCommission(jsonParams string) interface{}       // No type provided in docs
	AffiliateWithdraw(jsonParams string) interface{}         // No type provided in docs
	AffiliateCommissionDetail(jsonParams string) interface{} // No type provided in docs
	AffiliateReferral(jsonParams string) interface{}         // No type provided in docs
	Subaffiliates(jsonParams string) interface{}             // No type provided in docs
}

type spotClient struct {
	reqService *request.RequestService
}

func NewSpot(apiKey, apiSecret, baseURL string) Spot {
	return &spotClient{
		reqService: request.NewRequestService(apiKey, apiSecret, baseURL),
	}
}

func (c *spotClient) Ping(jsonParams string) interface{} {
	caseUrl := "/ping"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response, err := c.reqService.PublicGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) Time(jsonParams string) *ServerTime {
	caseUrl := "/time"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response ServerTime
	_, err := c.reqService.PublicGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) ApiSymbol(jsonParams string) *DefaultSymbolResponse {
	caseUrl := "/defaultSymbols"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response DefaultSymbolResponse
	_, err := c.reqService.PublicGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) ExchangeInfo(jsonParams string) *ExchangeInfo {
	caseUrl := "/exchangeInfo"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response ExchangeInfo
	_, err := c.reqService.PublicGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) Depth(jsonParams string) *OrderBook {
	caseUrl := "/depth"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response OrderBook
	_, err := c.reqService.PublicGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) Trades(jsonParams string) []RecentTrade {
	caseUrl := "/trades"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response []RecentTrade
	_, err := c.reqService.PublicGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) AggTrades(jsonParams string) []AggregateTrade {
	caseUrl := "/aggTrades"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response []AggregateTrade
	_, err := c.reqService.PublicGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) Kline(jsonParams string) []KlineData {
	caseUrl := "/klines"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response []KlineData
	_, err := c.reqService.PublicGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) AvgPrice(jsonParams string) *CurrentAveragePrice {
	caseUrl := "/avgPrice"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response CurrentAveragePrice
	_, err := c.reqService.PublicGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) Ticker24hr(jsonParams string) []Ticker24hr {
	caseUrl := "/ticker/24hr"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response []Ticker24hr
	_, err := c.reqService.PublicGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) Price(jsonParams string) []SymbolPriceTicker {
	caseUrl := "/ticker/price"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response []SymbolPriceTicker
	_, err := c.reqService.PublicGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) BookTicker(jsonParams string) []SymbolOrderBookTicker {
	caseUrl := "/ticker/bookTicker"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response []SymbolOrderBookTicker
	_, err := c.reqService.PublicGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) CreateSub(jsonParams string) *CreateSubAccountResponse {
	caseUrl := "/sub-account/virtualSubAccount"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response CreateSubAccountResponse
	_, err := c.reqService.PrivatePost(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) QuerySub(jsonParams string) *SubAccountList {
	caseUrl := "/sub-account/list"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response SubAccountList
	_, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) CreateSubApikey(jsonParams string) *CreateSubAccountAPIKeyResponse {
	caseUrl := "/sub-account/apiKey"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response CreateSubAccountAPIKeyResponse
	_, err := c.reqService.PrivatePost(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) QuerySubApikey(jsonParams string) *SubAccountAPIKeyResponse {
	caseUrl := "/sub-account/apiKey"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response SubAccountAPIKeyResponse
	_, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) DeleteSubApikey(jsonParams string) *DeleteSubAccountAPIKeyResponse {
	caseUrl := "/sub-account/apiKey"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response DeleteSubAccountAPIKeyResponse
	_, err := c.reqService.PrivateDelete(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) UniTransfer(jsonParams string) *MasterUniversalTransferResponse {
	caseUrl := "/capital/sub-account/universalTransfer"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response MasterUniversalTransferResponse
	_, err := c.reqService.PrivatePost(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) QueryUniTransfer(jsonParams string) *MasterUniversalTransferHistory {
	caseUrl := "/capital/sub-account/universalTransfer"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response MasterUniversalTransferHistory
	_, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) SelfSymbols(jsonParams string) *DefaultSymbolResponse {
	caseUrl := "/selfSymbols"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response DefaultSymbolResponse
	_, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) TestOrder(jsonParams string) interface{} {
	caseUrl := "/order/test"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response, err := c.reqService.PrivatePost(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) PlaceOrder(jsonParams string) *NewOrderResponse {
	caseUrl := "/order"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response NewOrderResponse
	_, err := c.reqService.PrivatePost(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) BatchOrder(jsonParams string) *BatchOrderResponse {
	caseUrl := "/batchOrders"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response BatchOrderResponse
	_, err := c.reqService.PrivatePost(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) CancelOrder(jsonParams string) *CancelOrderResponse {
	caseUrl := "/order"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response CancelOrderResponse
	_, err := c.reqService.PrivateDelete(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) CancelAllOrders(jsonParams string) []CancelOrderResponse {
	caseUrl := "/openOrders"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response []CancelOrderResponse
	_, err := c.reqService.PrivateDelete(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) QueryOrder(jsonParams string) *Order {
	caseUrl := "/order"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response Order
	_, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) OpenOrder(jsonParams string) []Order {
	caseUrl := "/openOrders"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response []Order
	_, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) AllOrders(jsonParams string) []Order {
	caseUrl := "/allOrders"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response []Order
	_, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) SpotAccountInfo(jsonParams string) *AccountInfo {
	caseUrl := "/account"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response AccountInfo
	_, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) SpotmyTrade(jsonParams string) []Trade {
	caseUrl := "/myTrades"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response []Trade
	_, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) MxDeduct(jsonParams string) *MXDeductStatus {
	caseUrl := "/mxDeduct/enable"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response MXDeductStatus
	_, err := c.reqService.PrivatePost(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) QueryMxDeduct(jsonParams string) *MXDeductStatus {
	caseUrl := "/mxDeduct/enable"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response MXDeductStatus
	_, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) QueryCurrencyInfo(jsonParams string) []CurrencyInfo {
	caseUrl := "/capital/config/getall"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response []CurrencyInfo
	_, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) Withdraw(jsonParams string) *WithdrawResponse {
	caseUrl := "/capital/withdraw/apply"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response WithdrawResponse
	_, err := c.reqService.PrivatePost(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) CancelWithdraw(jsonParams string) *CancelWithdrawResponse {
	caseUrl := "/capital/withdraw"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response CancelWithdrawResponse
	_, err := c.reqService.PrivateDelete(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) DepositHistory(jsonParams string) []DepositHistory {
	caseUrl := "/capital/deposit/hisrec"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response []DepositHistory
	_, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) WithdrawHistory(jsonParams string) []WithdrawHistory {
	caseUrl := "/capital/withdraw/historyl"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response []WithdrawHistory
	_, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) GenDepositAddress(jsonParams string) []GeneratedDepositAddress {
	caseUrl := "/capital/deposit/address"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response []GeneratedDepositAddress
	_, err := c.reqService.PrivatePost(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) DepositAddress(jsonParams string) []DepositAddress {
	caseUrl := "/capital/deposit/address"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response []DepositAddress
	_, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) WithdrawAddress(jsonParams string) *WithdrawAddress {
	caseUrl := "/capital/withdraw/address"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response WithdrawAddress
	_, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) Transfer(jsonParams string) *InternalTransferResponse {
	caseUrl := "/capital/transfer"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response InternalTransferResponse
	_, err := c.reqService.PrivatePost(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) TransferHistory(jsonParams string) *InternalTransferHistory {
	caseUrl := "/capital/transfer"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response InternalTransferHistory
	_, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) TransferHistoryById(jsonParams string) *UniversalTransferHistory {
	caseUrl := "/capital/transfer/tranId"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response UniversalTransferHistory
	_, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) ConvertList(jsonParams string) []ConvertibleAsset {
	caseUrl := "/capital/convert/list"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response []ConvertibleAsset
	_, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) Convert(jsonParams string) *DustTransferResponse {
	caseUrl := "/capital/convert"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response DustTransferResponse
	_, err := c.reqService.PrivatePost(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) ConvertHistory(jsonParams string) *DustLog {
	caseUrl := "/capital/convert"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response DustLog
	_, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) ETFInfo(jsonParams string) interface{} {
	caseUrl := "/etf/info"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) InternalTransfer(jsonParams string) *InternalTransferResponse {
	caseUrl := "/capital/transfer/internal"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response InternalTransferResponse
	_, err := c.reqService.PrivatePost(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) InternalTransferHistory(jsonParams string) *InternalTransferHistory {
	caseUrl := "/capital/transfer/internal"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	var response InternalTransferHistory
	_, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return &response
}

func (c *spotClient) CreateListenKey(jsonParams string) interface{} {
	caseUrl := "/userDataStream"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response, err := c.reqService.PrivatePost(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) KeepListenKey(jsonParams string) interface{} {
	caseUrl := "/userDataStream"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response, err := c.reqService.PrivatePut(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) CloseListenKey(jsonParams string) interface{} {
	caseUrl := "/userDataStream"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response, err := c.reqService.PrivateDelete(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) RebateHistory(jsonParams string) interface{} {
	caseUrl := "/rebate/taxQuery"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) RebateDetail(jsonParams string) interface{} {
	caseUrl := "/rebate/detail"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) SelfRecordsDetail(jsonParams string) interface{} {
	caseUrl := "/rebate/detail/kickback"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) ReferCode(jsonParams string) interface{} {
	caseUrl := "/rebate/referCode"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) AffiliateCommission(jsonParams string) interface{} {
	caseUrl := "/rebate/affiliate/commission"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) AffiliateWithdraw(jsonParams string) interface{} {
	caseUrl := "/rebate/affiliate/withdraw"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) AffiliateCommissionDetail(jsonParams string) interface{} {
	caseUrl := "/rebate/affiliate/commission/detail"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) AffiliateReferral(jsonParams string) interface{} {
	caseUrl := "/rebate/affiliate/referral"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}

func (c *spotClient) Subaffiliates(jsonParams string) interface{} {
	caseUrl := "/rebate/affiliate/subaffiliates"
	requestUrl := "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response, err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		return nil
	}
	return response
}
