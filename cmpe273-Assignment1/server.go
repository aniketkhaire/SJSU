package main

import (
    "fmt"
    "net/http"
    "net"
    "net/rpc"
    "net/rpc/jsonrpc"
    "os"
    "strings"
    "strconv"
    "encoding/json"
    "io/ioutil"
    "math"
)

//structure to store data fetched from yahoo api
type Stock struct {
    List struct {
        Resources []struct {
            Resource struct {
                Fields struct {
                    Name    string `json:"name"`
                    Price   string `json:"price"`
                    Symbol  string `json:"symbol"`
                    Ts      string `json:"ts"`
                    Type    string `json:"type"`
                    UTCTime string `json:"utctime"`
                    Volume  string `json:"volume"`
                } `json:"fields"`
            } `json:"resource"`
        } `json:"resources"`
    } `json:"list"`
}

//request structure for buying stocks
type Request struct {
    StockSymbolAndPercentage string
    Budget float64
}

//response structure from buying stocks
type Response struct {
    TradeId int
    Stocks [10]string
    UnvestedAmount float64
    Count int
}

//response structure from viewPortfolio
type PortfolioResponse struct{
    Stocks string
    CurrentMarketValue float64
    UnvestedAmount float64
}

type Arith int              
//variable to keep track of TradeId (starts with 0 by default)
var newId int 
//slice to store transactions
var SliceResponse []Response

//Remote procedure definition 
func (t *Arith) BuyStocks(request Request, res *Response) error {
    
    //split string by ","
    stockNames := strings.Split(request.StockSymbolAndPercentage,",")

    //increment TradeId
    newId++
    res.TradeId = newId
    
    //split string by ":"
    for i := range(stockNames){

        StockSymbol := strings.Split(stockNames[i], ":")        
        res.Stocks[i] = StockSymbol[0]

        //extract the percentage
        tempPercentage := StockSymbol[1]
        percentage,_ := strconv.ParseFloat(tempPercentage[:len(tempPercentage) -1], 64)

        //calculate the amount
        amount := request.Budget * (percentage/100)
        
        //call the yahoo finance API
        url := fmt.Sprintf("http://finance.yahoo.com/webservice/v1/symbols/%s/quote?format=json",res.Stocks[i])
        urlRes,err := http.Get(url)
        if err != nil{
            panic(err)
        }
        defer urlRes.Body.Close()

        body,err := ioutil.ReadAll(urlRes.Body)
        if err != nil{
            panic(err)
        }
        
        var stock Stock

        err = json.Unmarshal(body, &stock)
        if err != nil{
            panic(err)
        }
    
        //extract the cost
        stockPrice,_ := strconv.ParseFloat(stock.List.Resources[0].Resource.Fields.Price, 64) 
    
        //calculate the shares
        noOfShares := math.Floor(amount/stockPrice)
    
        //store invested amount (temporarily)
        res.UnvestedAmount += stockPrice * noOfShares
    
        //append the number of shares in stocks alongwith cost
        res.Stocks[i] = res.Stocks[i]+":"+strconv.FormatFloat(noOfShares,'f',0,64)+":$"+strconv.FormatFloat(stockPrice,'f',6,64)
    
        //increment count
        res.Count++
    }
    //calcuate uninvested amount (Budget - Invested amount)
    res.UnvestedAmount = request.Budget - res.UnvestedAmount
    
    //store data in slice
    SliceResponse = append(SliceResponse, *res)
    return nil
}

//Remote procedure definition 
func (t *Arith) Portfolio(index int, res *PortfolioResponse) error {
    
    if(index <= len(SliceResponse)){
        //fetch all the data in local variables
        request := SliceResponse[index]
        tempCount := request.Count
    
        for i:=0; i<tempCount; i++{
            tempData := strings.Split(request.Stocks[i], ":")               // symbol : noOfShares : $stockPrice
        
            //call the yahoo finance API
            url := fmt.Sprintf("http://finance.yahoo.com/webservice/v1/symbols/%s/quote?format=json",tempData[0])
            urlRes,err := http.Get(url)
            if err != nil{
                panic(err)
            }
            defer urlRes.Body.Close()

            body,err := ioutil.ReadAll(urlRes.Body)
            if err != nil{
                panic(err)
            }
      
            var stock Stock

            err = json.Unmarshal(body, &stock)
            if err != nil{
                panic(err)
            }
    
            //extract the cost
            tempPrice,_ := strconv.ParseFloat(stock.List.Resources[0].Resource.Fields.Price, 64) 
        
            //concatenate stock name and no. of shares
            res.Stocks = res.Stocks+tempData[0]+":"+tempData[1]
        
            //remove "$" from previous price string
            prevPricestr := tempData[2]
            prevPricestr = prevPricestr[1: len(prevPricestr)]
            //convert string to float
            prevPrice,_ := strconv.ParseFloat(prevPricestr, 64)
        
            //check whether profit or loss
            if tempPrice > prevPrice{
                res.Stocks +=":+$"+strconv.FormatFloat(tempPrice,'f', 6, 64)+" ; "
            }else if tempPrice < prevPrice{
                res.Stocks +=":-$"+strconv.FormatFloat(tempPrice,'f', 6, 64)+" ; "
            }else{
                res.Stocks +=":$"+strconv.FormatFloat(tempPrice,'f', 6, 64)+" ; "
            }

            //store no. of shares
            shareCount,_ := strconv.ParseFloat(tempData[1], 64)
            //calculate current market value
            tempCurrentMarketValue :=  shareCount* tempPrice
            res.CurrentMarketValue += tempCurrentMarketValue
        }
        //return unInvested amount
        res.UnvestedAmount = request.UnvestedAmount
    }
    return nil
}

func main() {

    arith := new(Arith)
    rpc.Register(arith)

    tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:1234")
    checkError(err)

    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)

    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        jsonrpc.ServeConn(conn)
    }
}

func checkError(err error) {
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)
    }
}