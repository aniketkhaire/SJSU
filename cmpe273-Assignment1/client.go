package main

import (
    "fmt"
    "log"
    "net/rpc/jsonrpc"
    "os"
    "strconv"
)

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

//request structure for viewing portfolio
type PortfolioRequest struct{
    TradeId int
}

//response structure from viewPortfolio
type PortfolioResponse struct{
    Stocks string
    CurrentMarketValue float64
    UnvestedAmount float64
}

func main() {
    if len(os.Args) == 3 {
        callBuyingStocks()              //no. of arguments = 3 (program name, stock name&pecentage, budget),  then buy stocks
    }else if len(os.Args) == 2{
        callViewPortfolio()             //no. of arguments = 2 (program name, trade id), then view portfolio for that id
    }else{
        fmt.Println("Usage: ", os.Args[0], "localhost:1234")
        log.Fatal(1)
    }
}


func callBuyingStocks(){
    var request Request
    
    request.StockSymbolAndPercentage = os.Args[1]
    request.Budget,_ = strconv.ParseFloat(os.Args[2], 64)

    client, err := jsonrpc.Dial("tcp", "localhost:1234")
    if err != nil {
        log.Fatal("dialing:", err)
    }
    
    var res Response

    err = client.Call("Arith.BuyStocks", request, &res)
    if err != nil {
        log.Fatal("arith error:", err)
    }

    fmt.Printf("Received Response... \nTradeId : %d", res.TradeId)
    fmt.Printf("\nStock : ")
    for i:=0; i<res.Count; i++{
        fmt.Printf("\"%s\" ; ", res.Stocks[i])        
    }
    fmt.Printf("\nUnvested Amt : %f\n", res.UnvestedAmount)
}

func callViewPortfolio(){
    var portfolioRequest PortfolioRequest
    
    i64,_ := strconv.ParseInt(os.Args[1], 10, 32)
    if(i64 > 0){
        portfolioRequest.TradeId = int(i64)
        client, err := jsonrpc.Dial("tcp", "localhost:1234")
        if err != nil {
            log.Fatal("dialing:", err)
        }
    
        var portfolioRes PortfolioResponse
        err = client.Call("Arith.Portfolio", portfolioRequest.TradeId - 1, &portfolioRes)
        if err != nil {
            log.Fatal("arith error:", err)
        }

        fmt.Printf("Received Response... \n")
        fmt.Printf("\nStock : ")
        fmt.Printf("\"%s\"", portfolioRes.Stocks)        
        fmt.Printf("\nCurrent Market Value : %f", portfolioRes.CurrentMarketValue)
        fmt.Printf("\nUnvested Amt : %f\n", portfolioRes.UnvestedAmount)
    }
}