package main

import "log"

type NetInHandler struct {
	deaman *Deaman
}

func StartNetInHandler(deaman *Deaman) {
	netInHandler := NetInHandler{deaman: deaman}
	go netInHandler.run()
}

func (ni *NetInHandler) run() {
	log.Printf("NetInHandler is running ...")
	//TODO
}

func (ni *NetInHandler) NewTransaction(t *Transaction) {
	//TODO
}

func (ni *NetInHandler) NewBlock(b *Block) {
	//TODO
}
