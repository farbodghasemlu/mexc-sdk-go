package mexcsdk

import (
	"github.com/farbodghasemlu/mexc-sdk-go/request"
	"fmt"
)
type Spot interface {
    
    Ping(jsonParams string) interface{}
    Time(jsonParams string) interface{}
    ApiSymbol(jsonParams string) interface{}
    ExchangeInfo(jsonParams string) interface{}
    Depth(jsonParams string) interface{}
    Trades(jsonParams string) interface{}
    AggTrades(jsonParams string) interface{}
    Kline(jsonParams string) interface{}
    AvgPrice(jsonParams string) interface{}
    Ticker24hr(jsonParams string) interface{}
    Price(jsonParams string) interface{}
    BookTicker(jsonParams string) interface{}
    
    
    CreateSub(jsonParams string) interface{}
    QuerySub(jsonParams string) interface{}
    CreateSubApikey(jsonParams string) interface{}
    QuerySubApikey(jsonParams string) interface{}
    DeleteSubApikey(jsonParams string) interface{}
    UniTransfer(jsonParams string) interface{}
    QueryUniTransfer(jsonParams string) interface{}
    
    
    SelfSymbols(jsonParams string) interface{}
    TestOrder(jsonParams string) interface{}
    PlaceOrder(jsonParams string) interface{}
    BatchOrder(jsonParams string) interface{}
    CancelOrder(jsonParams string) interface{}
    CancelAllOrders(jsonParams string) interface{}
    QueryOrder(jsonParams string) interface{}
    OpenOrder(jsonParams string) interface{}
    AllOrders(jsonParams string) interface{}
    SpotAccountInfo(jsonParams string) interface{}
    SpotmyTrade(jsonParams string) interface{}
    MxDeduct(jsonParams string) interface{}
    QueryMxDeduct(jsonParams string) interface{}
    
    
    QueryCurrencyInfo(jsonParams string) interface{}
    Withdraw(jsonParams string) interface{}
    CancelWithdraw(jsonParams string) interface{}
    DepositHistory(jsonParams string) interface{}
    WithdrawHistory(jsonParams string) interface{}
    GenDepositAddress(jsonParams string) interface{}
    DepositAddress(jsonParams string) interface{}
    WithdrawAddress(jsonParams string) interface{}
    Transfer(jsonParams string) interface{}
    TransferHistory(jsonParams string) interface{}
    TransferHistoryById(jsonParams string) interface{}
    ConvertList(jsonParams string) interface{}
    Convert(jsonParams string) interface{}
    ConvertHistory(jsonParams string) interface{}
    ETFInfo(jsonParams string) interface{}
    InternalTransfer(jsonParams string) interface{}
    InternalTransferHistory(jsonParams string) interface{}
    
    
    CreateListenKey(jsonParams string) interface{}
    KeepListenKey(jsonParams string) interface{}
    CloseListenKey(jsonParams string) interface{}
    
    
    RebateHistory(jsonParams string) interface{}
    RebateDetail(jsonParams string) interface{}
    SelfRecordsDetail(jsonParams string) interface{}
    ReferCode(jsonParams string) interface{}
    AffiliateCommission(jsonParams string) interface{}
    AffiliateWithdraw(jsonParams string) interface{}
    AffiliateCommissionDetail(jsonParams string) interface{}
    AffiliateReferral(jsonParams string) interface{}
    Subaffiliates(jsonParams string) interface{}
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
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PublicGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) Time(jsonParams string) interface{} {
	caseUrl := "/time"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PublicGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) ApiSymbol(jsonParams string) interface{} {
	caseUrl := "/defaultSymbols"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PublicGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) ExchangeInfo(jsonParams string) interface{} {
	caseUrl := "/exchangeInfo"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PublicGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) Depth(jsonParams string) interface{} {
	caseUrl := "/depth"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PublicGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) Trades(jsonParams string) interface{} {
	caseUrl := "/trades"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PublicGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) AggTrades(jsonParams string) interface{} {
	caseUrl := "/aggTrades"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PublicGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) Kline(jsonParams string) interface{} {
	caseUrl := "/klines"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PublicGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) AvgPrice(jsonParams string) interface{} {
	caseUrl := "/avgPrice"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PublicGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) Ticker24hr(jsonParams string) interface{} {
	caseUrl := "/ticker/24hr"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PublicGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) Price(jsonParams string) interface{} {
	caseUrl := "/ticker/price"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PublicGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) BookTicker(jsonParams string) interface{} {
	caseUrl := "/ticker/bookTicker"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PublicGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}




func (c *spotClient) CreateSub(jsonParams string) interface{} {
	caseUrl := "/sub-account/virtualSubAccount"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivatePost(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) QuerySub(jsonParams string) interface{} {
	caseUrl := "/sub-account/list"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) CreateSubApikey(jsonParams string) interface{} {
	caseUrl := "/sub-account/apiKey"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivatePost(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) QuerySubApikey(jsonParams string) interface{} {
	caseUrl := "/sub-account/apiKey"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) DeleteSubApikey(jsonParams string) interface{} {
	caseUrl := "/sub-account/apiKey"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateDelete(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) UniTransfer(jsonParams string) interface{} {
	caseUrl := "/capital/sub-account/universalTransfer"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivatePost(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) QueryUniTransfer(jsonParams string) interface{} {
	caseUrl := "/capital/sub-account/universalTransfer"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}




func (c *spotClient) SelfSymbols(jsonParams string) interface{} {
	caseUrl := "/selfSymbols"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) TestOrder(jsonParams string) interface{} {
	caseUrl := "/order/test"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivatePost(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) PlaceOrder(jsonParams string) interface{} {
	caseUrl := "/order"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivatePost(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) BatchOrder(jsonParams string) interface{} {
	caseUrl := "/batchOrders"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivatePost(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) CancelOrder(jsonParams string) interface{} {
	caseUrl := "/order"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateDelete(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) CancelAllOrders(jsonParams string) interface{} {
	caseUrl := "/openOrders"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateDelete(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) QueryOrder(jsonParams string) interface{} {
	caseUrl := "/order"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) OpenOrder(jsonParams string) interface{} {
	caseUrl := "/openOrders"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) AllOrders(jsonParams string) interface{} {
	caseUrl := "/allOrders"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) SpotAccountInfo(jsonParams string) interface{} {
	caseUrl := "/account"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) SpotmyTrade(jsonParams string) interface{} {
	caseUrl := "/myTrades"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) MxDeduct(jsonParams string) interface{} {
	caseUrl := "/mxDeduct/enable"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivatePost(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) QueryMxDeduct(jsonParams string) interface{} {
	caseUrl := "/mxDeduct/enable"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}




func (c *spotClient) QueryCurrencyInfo(jsonParams string) interface{} {
	caseUrl := "/capital/config/getall"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) Withdraw(jsonParams string) interface{} {
	caseUrl := "/capital/withdraw/apply"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivatePost(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) CancelWithdraw(jsonParams string) interface{} {
	caseUrl := "/capital/withdraw"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateDelete(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) DepositHistory(jsonParams string) interface{} {
	caseUrl := "/capital/deposit/hisrec"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) WithdrawHistory(jsonParams string) interface{} {
	caseUrl := "/capital/withdraw/historyl"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) GenDepositAddress(jsonParams string) interface{} {
	caseUrl := "/capital/deposit/address"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivatePost(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) DepositAddress(jsonParams string) interface{} {
	caseUrl := "/capital/deposit/address"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) WithdrawAddress(jsonParams string) interface{} {
	caseUrl := "/capital/withdraw/address"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) Transfer(jsonParams string) interface{} {
	caseUrl := "/capital/transfer"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivatePost(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) TransferHistory(jsonParams string) interface{} {
	caseUrl := "/capital/transfer"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) TransferHistoryById(jsonParams string) interface{} {
	caseUrl := "/capital/transfer/tranId"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) ConvertList(jsonParams string) interface{} {
	caseUrl := "/capital/convert/list"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) Convert(jsonParams string) interface{} {
	caseUrl := "/capital/convert"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivatePost(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) ConvertHistory(jsonParams string) interface{} {
	caseUrl := "/capital/convert"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) ETFInfo(jsonParams string) interface{} {
	caseUrl := "/etf/info"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) InternalTransfer(jsonParams string) interface{} {
	caseUrl := "/capital/transfer/internal"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivatePost(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) InternalTransferHistory(jsonParams string) interface{} {
	caseUrl := "/capital/transfer/internal"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}




func (c *spotClient) CreateListenKey(jsonParams string) interface{} {
	caseUrl := "/userDataStream"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivatePost(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) KeepListenKey(jsonParams string) interface{} {
	caseUrl := "/userDataStream"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivatePut(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) CloseListenKey(jsonParams string) interface{} {
	caseUrl := "/userDataStream"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateDelete(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}




func (c *spotClient) RebateHistory(jsonParams string) interface{} {
	caseUrl := "/rebate/taxQuery"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) RebateDetail(jsonParams string) interface{} {
	caseUrl := "/rebate/detail"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) SelfRecordsDetail(jsonParams string) interface{} {
	caseUrl := "/rebate/detail/kickback"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) ReferCode(jsonParams string) interface{} {
	caseUrl := "/rebate/referCode"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) AffiliateCommission(jsonParams string) interface{} {
	caseUrl := "/rebate/affiliate/commission"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) AffiliateWithdraw(jsonParams string) interface{} {
	caseUrl := "/rebate/affiliate/withdraw"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) AffiliateCommissionDetail(jsonParams string) interface{} {
	caseUrl := "/rebate/affiliate/commission/detail"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) AffiliateReferral(jsonParams string) interface{} {
	caseUrl := "/rebate/affiliate/referral"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}


func (c *spotClient) Subaffiliates(jsonParams string) interface{} {
	caseUrl := "/rebate/affiliate/subaffiliates"
	requestUrl :=  "https://api.mexc.com/api/v3" + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response , err := c.reqService.PrivateGet(caseUrl, jsonParams)
	if err != nil {
		// Handle error appropriately
		return nil
	}
	return response
}
