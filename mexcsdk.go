package mexcsdk

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"

	"github.com/farbodghasemlu/mexc-sdk-go/request"
)

type Spot interface {
	Ping() error
	Time() (*ServerTime, error)
	ExchangeInfo(params map[string]string) ([]ExchangeInfo, error)
	Depth(params map[string]string) (*OrderBook, error)
	Trades(params map[string]string) ([]RecentTrade, error)
	AggTrades(params map[string]string) ([]AggregateTrade, error)
	Kline(params map[string]string) ([]KlineData, error)
	AvgPrice(params map[string]string) (*CurrentAveragePrice, error)
	Ticker24hr(params map[string]string) ([]Ticker24hr, error)
	Price(params map[string]string) ([]SymbolPriceTicker, error)
	BookTicker(params map[string]string) ([]SymbolOrderBookTicker, error)
	CreateSubAccount(params map[string]string) (*CreateSubAccountResponse, error)
	QuerySubAccounts(params map[string]string) (*SubAccountList, error)
	CreateSubAccountAPIKey(params map[string]string) (*CreateSubAccountAPIKeyResponse, error)
	QuerySubAccountAPIKeys(params map[string]string) (*SubAccountAPIKeyResponse, error)
	DeleteSubAccountAPIKey(params map[string]string) (*DeleteSubAccountAPIKeyResponse, error)
	UniversalTransfer(params map[string]string) (*MasterUniversalTransferResponse, error)
	QueryUniversalTransfer(params map[string]string) (*MasterUniversalTransferHistory, error)
	NewOrder(params map[string]string) (*NewOrderResponse, error)
	TestNewOrder(params map[string]string) error
	CancelOrder(params map[string]string) (*CancelOrderResponse, error)
	CancelAllOpenOrders(params map[string]string) ([]CancelOrderResponse, error)
	QueryOrder(params map[string]string) (*Order, error)
	CurrentOpenOrders(params map[string]string) ([]Order, error)
	AllOrders(params map[string]string) ([]Order, error)
	AccountInfo() (*AccountInfo, error)
	MyTrades(params map[string]string) ([]Trade, error)
	QueryCurrencyInfo(params map[string]string) ([]CurrencyInfo, error)
	Withdraw(params map[string]string) (*WithdrawResponse, error)
	DepositHistory(params map[string]string) ([]DepositHistory, error)
	WithdrawHistory(params map[string]string) ([]WithdrawHistory, error)
	DepositAddress(params map[string]string) ([]DepositAddress, error)
	// AccountStatus(params map[string]string) (*AccountStatusResponse, error)
	// APITradingStatus(params map[string]string) (*APITradingStatusResponse, error)
	CreateListenKey() (string, error)
	KeepaliveListenKey(listenKey string) error
	CloseListenKey(listenKey string) error
}

type spotClient struct {
	reqService *request.RequestService
	logger     *log.Logger
	baseURL    string
	apiKey     string
	apiSecret  string
}

func NewSpot(apiKey, apiSecret, baseURL string, logger *log.Logger) Spot {
	if logger == nil {
		logger = log.Default()
	}
	// if !strings.HasSuffix(baseURL, "/") {
	// 	baseURL += "/"
	// }
	return &spotClient{
		reqService: request.NewRequestService(apiKey, apiSecret, baseURL),
		logger:     logger,
		baseURL:    baseURL,
		apiKey:     apiKey,
		apiSecret:  apiSecret,
	}
}

func (c *spotClient) paramsToJSON(params map[string]string) string {
	if params == nil {
		return ""
	}
	jsonData, _ := json.Marshal(params)
	return string(jsonData)
}

func (c *spotClient) handleResponse(response interface{}, target interface{}) error {
	// Type assert to *resty.Response
	resp, ok := response.(*resty.Response)
	if !ok {
		return fmt.Errorf("invalid response type: %T", response)
	}

	// Check for HTTP errors
	if resp.IsError() {
		return fmt.Errorf("API error: %s", resp.Status())
	}

	// Unmarshal the JSON body
	if err := json.Unmarshal(resp.Body(), target); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return nil
}

func (c *spotClient) Ping() error {
	_, err := c.reqService.PublicGet("/ping", "")
	if err != nil {
		c.logger.Printf("Ping failed: %v", err)
	}
	return err
}

func (c *spotClient) Time() (*ServerTime, error) {
	response, err := c.reqService.PublicGet("/time", "")
	if err != nil {
		c.logger.Printf("Failed to get server time: %v", err)
		return nil, err
	}
	var serverTime ServerTime
	if err := c.handleResponse(response, &serverTime); err != nil {
		return nil, err
	}
	return &serverTime, nil
}

func (c *spotClient) ExchangeInfo(params map[string]string) ([]ExchangeInfo, error) {
	response, err := c.reqService.PublicGet("/exchangeInfo", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to get exchange info: %v", err)
		return nil, err
	}
	var info []ExchangeInfo
	if err := c.handleResponse(response, &info); err != nil {
		return nil, err
	}
	return info, nil
}

func (c *spotClient) Depth(params map[string]string) (*OrderBook, error) {
	response, err := c.reqService.PublicGet("/depth", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to get order book depth: %v", err)
		return nil, err
	}
	var orderBook OrderBook
	if err := c.handleResponse(response, &orderBook); err != nil {
		return nil, err
	}
	return &orderBook, nil
}

func (c *spotClient) Trades(params map[string]string) ([]RecentTrade, error) {
	response, err := c.reqService.PublicGet("/trades", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to get recent trades: %v", err)
		return nil, err
	}
	var trades []RecentTrade
	if err := c.handleResponse(response, &trades); err != nil {
		return nil, err
	}
	return trades, nil
}

func (c *spotClient) AggTrades(params map[string]string) ([]AggregateTrade, error) {
	response, err := c.reqService.PublicGet("/aggTrades", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to get aggregate trades: %v", err)
		return nil, err
	}
	var trades []AggregateTrade
	if err := c.handleResponse(response, &trades); err != nil {
		return nil, err
	}
	return trades, nil
}

func (c *spotClient) Kline(params map[string]string) ([]KlineData, error) {
	response, err := c.reqService.PublicGet("/klines", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to get kline data: %v", err)
		return nil, err
	}
	var klines []KlineData
	if err := c.handleResponse(response, &klines); err != nil {
		return nil, err
	}
	return klines, nil
}

func (c *spotClient) AvgPrice(params map[string]string) (*CurrentAveragePrice, error) {
	response, err := c.reqService.PublicGet("/avgPrice", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to get average price: %v", err)
		return nil, err
	}
	var avgPrice CurrentAveragePrice
	if err := c.handleResponse(response, &avgPrice); err != nil {
		return nil, err
	}
	return &avgPrice, nil
}

func (c *spotClient) Ticker24hr(params map[string]string) ([]Ticker24hr, error) {
	response, err := c.reqService.PublicGet("/ticker/24hr", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to get 24hr ticker: %v", err)
		return nil, err
	}
	var tickers []Ticker24hr
	if err := c.handleResponse(response, &tickers); err != nil {
		return nil, err
	}
	return tickers, nil
}

func (c *spotClient) Price(params map[string]string) ([]SymbolPriceTicker, error) {
	response, err := c.reqService.PublicGet("/ticker/price", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to get price ticker: %v", err)
		return nil, err
	}
	var tickers []SymbolPriceTicker
	if err := c.handleResponse(response, &tickers); err != nil {
		return nil, err
	}
	return tickers, nil
}

func (c *spotClient) BookTicker(params map[string]string) ([]SymbolOrderBookTicker, error) {
	response, err := c.reqService.PublicGet("/ticker/bookTicker", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to get book ticker: %v", err)
		return nil, err
	}
	var tickers []SymbolOrderBookTicker
	if err := c.handleResponse(response, &tickers); err != nil {
		return nil, err
	}
	return tickers, nil
}

func (c *spotClient) CreateSubAccount(params map[string]string) (*CreateSubAccountResponse, error) {
	response, err := c.reqService.PrivatePost("/sub-account/virtualSubAccount", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to create sub-account: %v", err)
		return nil, err
	}
	var result CreateSubAccountResponse
	if err := c.handleResponse(response, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *spotClient) QuerySubAccounts(params map[string]string) (*SubAccountList, error) {
	response, err := c.reqService.PrivateGet("/sub-account/list", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to query sub-accounts: %v", err)
		return nil, err
	}
	var result SubAccountList
	if err := c.handleResponse(response, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *spotClient) CreateSubAccountAPIKey(params map[string]string) (*CreateSubAccountAPIKeyResponse, error) {
	response, err := c.reqService.PrivatePost("/sub-account/apiKey", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to create sub-account API key: %v", err)
		return nil, err
	}
	var result CreateSubAccountAPIKeyResponse
	if err := c.handleResponse(response, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *spotClient) QuerySubAccountAPIKeys(params map[string]string) (*SubAccountAPIKeyResponse, error) {
	response, err := c.reqService.PrivateGet("/sub-account/apiKey", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to query sub-account API keys: %v", err)
		return nil, err
	}
	var result SubAccountAPIKeyResponse
	if err := c.handleResponse(response, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *spotClient) DeleteSubAccountAPIKey(params map[string]string) (*DeleteSubAccountAPIKeyResponse, error) {
	response, err := c.reqService.PrivateDelete("/sub-account/apiKey", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to delete sub-account API key: %v", err)
		return nil, err
	}
	var result DeleteSubAccountAPIKeyResponse
	if err := c.handleResponse(response, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *spotClient) UniversalTransfer(params map[string]string) (*MasterUniversalTransferResponse, error) {
	response, err := c.reqService.PrivatePost("/capital/sub-account/universalTransfer", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to perform universal transfer: %v", err)
		return nil, err
	}
	var result MasterUniversalTransferResponse
	if err := c.handleResponse(response, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *spotClient) QueryUniversalTransfer(params map[string]string) (*MasterUniversalTransferHistory, error) {
	response, err := c.reqService.PrivateGet("/capital/sub-account/universalTransfer", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to query universal transfer: %v", err)
		return nil, err
	}
	var result MasterUniversalTransferHistory
	if err := c.handleResponse(response, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *spotClient) NewOrder(params map[string]string) (*NewOrderResponse, error) {
	response, err := c.reqService.PrivatePost("/order", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to create new order: %v", err)
		return nil, err
	}
	var result NewOrderResponse
	if err := c.handleResponse(response, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *spotClient) TestNewOrder(params map[string]string) error {
	_, err := c.reqService.PrivatePost("/order/test", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to test new order: %v", err)
	}
	return err
}

func (c *spotClient) CancelOrder(params map[string]string) (*CancelOrderResponse, error) {
	response, err := c.reqService.PrivateDelete("/order", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to cancel order: %v", err)
		return nil, err
	}
	var result CancelOrderResponse
	if err := c.handleResponse(response, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *spotClient) CancelAllOpenOrders(params map[string]string) ([]CancelOrderResponse, error) {
	response, err := c.reqService.PrivateDelete("/openOrders", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to cancel all open orders: %v", err)
		return nil, err
	}
	var result []CancelOrderResponse
	if err := c.handleResponse(response, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *spotClient) QueryOrder(params map[string]string) (*Order, error) {
	response, err := c.reqService.PrivateGet("/order", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to query order: %v", err)
		return nil, err
	}
	var result Order
	if err := c.handleResponse(response, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *spotClient) CurrentOpenOrders(params map[string]string) ([]Order, error) {
	response, err := c.reqService.PrivateGet("/openOrders", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to get current open orders: %v", err)
		return nil, err
	}
	var result []Order
	if err := c.handleResponse(response, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *spotClient) AllOrders(params map[string]string) ([]Order, error) {
	response, err := c.reqService.PrivateGet("/allOrders", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to get all orders: %v", err)
		return nil, err
	}
	var result []Order
	if err := c.handleResponse(response, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *spotClient) AccountInfo() (*AccountInfo, error) {
	response, err := c.reqService.PrivateGet("/account", "")
	if err != nil {
		c.logger.Printf("Failed to get account info: %v", err)
		return nil, err
	}
	var result AccountInfo
	if err := c.handleResponse(response, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *spotClient) MyTrades(params map[string]string) ([]Trade, error) {
	response, err := c.reqService.PrivateGet("/myTrades", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to get my trades: %v", err)
		return nil, err
	}
	var result []Trade
	if err := c.handleResponse(response, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *spotClient) QueryCurrencyInfo(params map[string]string) ([]CurrencyInfo, error) {
	response, err := c.reqService.PrivateGet("/capital/config/getall", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to query currency info: %v", err)
		return nil, err
	}
	var result []CurrencyInfo
	if err := c.handleResponse(response, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *spotClient) Withdraw(params map[string]string) (*WithdrawResponse, error) {
	response, err := c.reqService.PrivatePost("/capital/withdraw/apply", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to withdraw: %v", err)
		return nil, err
	}
	var result WithdrawResponse
	if err := c.handleResponse(response, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *spotClient) DepositHistory(params map[string]string) ([]DepositHistory, error) {
	response, err := c.reqService.PrivateGet("/capital/deposit/hisrec", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to get deposit history: %v", err)
		return nil, err
	}
	var result []DepositHistory
	if err := c.handleResponse(response, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *spotClient) WithdrawHistory(params map[string]string) ([]WithdrawHistory, error) {
	response, err := c.reqService.PrivateGet("/capital/withdraw/history", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to get withdraw history: %v", err)
		return nil, err
	}
	var result []WithdrawHistory
	if err := c.handleResponse(response, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *spotClient) DepositAddress(params map[string]string) ([]DepositAddress, error) {
	response, err := c.reqService.PrivateGet("/capital/deposit/address", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to get deposit address: %v", err)
		return nil, err
	}
	var result []DepositAddress
	if err := c.handleResponse(response, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// func (c *spotClient) AccountStatus(params map[string]string) (*AccountStatusResponse, error) {
// 	response, err := c.reqService.PrivateGet("/account/status", c.paramsToJSON(params))
// 	if err != nil {
// 		c.logger.Printf("Failed to get account status: %v", err)
// 		return nil, err
// 	}
// 	var result AccountStatusResponse
// 	if err := c.handleResponse(response, &result); err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }

// func (c *spotClient) APITradingStatus(params map[string]string) (*APITradingStatusResponse, error) {
// 	response, err := c.reqService.PrivateGet("/account/apiTradingStatus", c.paramsToJSON(params))
// 	if err != nil {
// 		c.logger.Printf("Failed to get API trading status: %v", err)
// 		return nil, err
// 	}
// 	var result APITradingStatusResponse
// 	if err := c.handleResponse(response, &result); err != nil {
// 		return nil, err
// 	}
// 	return &result, nil
// }

func (c *spotClient) CreateListenKey() (string, error) {
	response, err := c.reqService.PrivatePost("/userDataStream", "")
	if err != nil {
		c.logger.Printf("Failed to create listen key: %v", err)
		return "", err
	}
	var result struct {
		ListenKey string `json:"listenKey"`
	}
	if err := c.handleResponse(response, &result); err != nil {
		return "", err
	}
	return result.ListenKey, nil
}

func (c *spotClient) KeepaliveListenKey(listenKey string) error {
	params := map[string]string{"listenKey": listenKey}
	_, err := c.reqService.PrivatePut("/userDataStream", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to keepalive listen key: %v", err)
	}
	return err
}

func (c *spotClient) CloseListenKey(listenKey string) error {
	params := map[string]string{"listenKey": listenKey}
	_, err := c.reqService.PrivateDelete("/userDataStream", c.paramsToJSON(params))
	if err != nil {
		c.logger.Printf("Failed to close listen key: %v", err)
	}
	return err
}
