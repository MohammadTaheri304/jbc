package main

import (
	"encoding/json"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

//ContentTypeAPPLICATIONـJSON constant value for json header
const ContentTypeAPPLICATIONـJSON = "application/json; application/json; charset=utf-8"

//NetInHandler encapsulate Net-In-handler state and it's dependencies
type NetInHandler struct {
	deaman *Deaman
}

//StartNetInHandler creates an instance of Net-In-handler and run it on a goroutine
func StartNetInHandler(deaman *Deaman) {
	netInHandler := NetInHandler{deaman: deaman}
	go netInHandler.run()
}

func (ni *NetInHandler) run() {
	log.Printf("NetInHandler is running ...")
	router := router.New()
	router.GET("/echo", ni.echo)
	router.POST("/transaction", ni.newTransaction)
	router.POST("/block", ni.newBlock)
	go fasthttp.ListenAndServe("0.0.0.0:8090", router.Handler)
}

func (ni *NetInHandler) newTransaction(context *fasthttp.RequestCtx) {
	trx := &Transaction{}
	err := json.Unmarshal(context.Request.Body(), trx)
	if err != nil {
		log.Fatalf("Error in newTransaction unmarshal, %+v\n", err)
		context.Error("Bad.Data", 403)
		return
	}

	ni.deaman.NetInTransactionChan <- trx
	context.SuccessString(ContentTypeAPPLICATIONـJSON, "{\"ok\": true}")
}

func (ni *NetInHandler) newBlock(context *fasthttp.RequestCtx) {
	blk := &Block{}
	err := json.Unmarshal(context.Request.Body(), blk)
	if err != nil {
		log.Fatalf("Error in newBlock unmarshal, %+v\n", err)
		context.Error("Bad.Data", 403)
		return
	}

	ni.deaman.NetInBlockChan <- blk
	context.SuccessString(ContentTypeAPPLICATIONـJSON, "{\"ok\": true}")
}

func (ni *NetInHandler) echo(context *fasthttp.RequestCtx) {
	context.SuccessString(ContentTypeAPPLICATIONـJSON, "{\"ok\": true}")
}
