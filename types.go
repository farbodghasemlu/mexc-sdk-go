package mexcsdk

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
	A  interface{} `json:"a"` // Aggregate trade Id
	F  interface{} `json:"f"` // First trade Id
	L  interface{} `json:"l"` // Last trade Id
	P  string      `json:"p"` // Price
	Q  string      `json:"q"` // Quantity
	T  int64       `json:"T"` // Timestamp
	M  bool        `json:"m"` // Was the buyer the maker?
	M2 bool        `json:"M"` // Was the trade the best price match?
}

// RecentTrade represents recent trades
type RecentTrade struct {
	Id           interface{} `json:"id"`
	Price        string      `json:"price"`
	Qty          string      `json:"qty"`
	QuoteQty     string      `json:"quoteQty"`
	Time         int64       `json:"time"`
	IsBuyerMaker bool        `json:"isBuyerMaker"`
	IsBestMatch  bool        `json:"isBestMatch"`
}

// OrderBook represents order book data
type OrderBook struct {
	LastUpdateId int64      `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

// ExchangeInfo represents exchange information
type ExchangeInfo struct {
	Timezone        string        `json:"timezone"`
	ServerTime      int64         `json:"serverTime"`
	RateLimits      []interface{} `json:"rateLimits"`
	ExchangeFilters []interface{} `json:"exchangeFilters"`
	Symbols         []Symbol      `json:"symbols"`
}

type Symbol struct {
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
	IsSpotTradingAllowed       bool          `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed     bool          `json:"isMarginTradingAllowed"`
	QuoteAmountPrecision       string        `json:"quoteAmountPrecision"`
	BaseSizePrecision          string        `json:"baseSizePrecision"`
	Permissions                []string      `json:"permissions"`
	Filters                    []interface{} `json:"filters"`
	MaxQuoteAmount             string        `json:"maxQuoteAmount"`
	MakerCommission            string        `json:"makerCommission"`
	TakerCommission            string        `json:"takerCommission"`
	QuoteAmountPrecisionMarket string        `json:"quoteAmountPrecisionMarket"`
	MaxQuoteAmountMarket       string        `json:"maxQuoteAmountMarket"`
	FullName                   string        `json:"fullName"`
	TradeSideType              int           `json:"tradeSideType"`
	St                         bool          `json:"st"`
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
	TranId          string `json:"tranId"`
	FromAccount     string `json:"fromAccount"`
	ToAccount       string `json:"toAccount"`
	ClientTranId    string `json:"clientTranId"`
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
	TranId int64 `json:"tranId"`
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
	Uid        string `json:"uid"`
}

// CreateSubAccountResponse represents response for creating sub-account
type CreateSubAccountResponse struct {
	SubAccount string `json:"subAccount"`
	Note       string `json:"note"`
}

// Trade represents account trade information
type Trade struct {
	Symbol          string `json:"symbol"`
	Id              string `json:"id"`
	OrderId         string `json:"orderId"`
	OrderListId     int64  `json:"orderListId"`
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
	ClientOrderId   string `json:"clientOrderId"`
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
	OrderId             int64  `json:"orderId"`
	OrderListId         int64  `json:"orderListId"`
	ClientOrderId       string `json:"clientOrderId"`
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
	OrigClientOrderId   string `json:"origClientOrderId"`
	OrderId             int64  `json:"orderId"`
	ClientOrderId       string `json:"clientOrderId"`
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
	OrderId     string `json:"orderId"`
	OrderListId int64  `json:"orderListId"`
}

type BatchOrderError struct {
	Symbol           string `json:"symbol,omitempty"`
	OrderId          string `json:"orderId,omitempty"`
	NewClientOrderId string `json:"newClientOrderId,omitempty"`
	OrderListId      int64  `json:"orderListId,omitempty"`
	Msg              string `json:"msg"`
	Code             int    `json:"code"`
}

// NewOrderResponse represents new order response
type NewOrderResponse struct {
	Symbol       string `json:"symbol"`
	OrderId      string `json:"orderId"`
	OrderListId  int64  `json:"orderListId"`
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
	Id string `json:"id"`
}

// InternalTransferHistory represents internal transfer history data
type InternalTransferHistory struct {
	Page         int                      `json:"page"`
	TotalRecords int                      `json:"totalRecords"`
	TotalPageNum int                      `json:"totalPageNum"`
	Data         []InternalTransferRecord `json:"data"`
}

type InternalTransferRecord struct {
	TranId        string `json:"tranId"`
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
	TranId string `json:"tranId"`
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
	Id      string `json:"id"`
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
	TranId          string `json:"tranId"`
	ClientTranId    string `json:"clientTranId"`
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
	TranId string `json:"tranId"`
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
	Id             string `json:"id"`
	TxId           string `json:"txId"`
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
	CoinId         string `json:"coinId"`
	VcoinId        string `json:"vcoinId"`
}

// DepositHistory represents deposit history
type DepositHistory struct {
	Amount        string `json:"amount"`
	Coin          string `json:"coin"`
	Network       string `json:"network"`
	Status        int    `json:"status"`
	Address       string `json:"address"`
	TxId          string `json:"txId"`
	InsertTime    int64  `json:"insertTime"`
	UnlockConfirm string `json:"unlockConfirm"`
	ConfirmTimes  string `json:"confirmTimes"`
	Memo          string `json:"memo"`
}

// CancelWithdrawResponse represents response for cancel withdraw
type CancelWithdrawResponse struct {
	Id string `json:"id"`
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
