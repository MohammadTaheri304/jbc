package main

import "log"

type NetOutHandler struct {
	deaman *Deaman
}

// StartNetOutHandler start net-out handler
func StartNetOutHandler(deaman *Deaman) {
	netOutHandler := NetOutHandler{deaman: deaman}
	go netOutHandler.run()
}

func (no *NetOutHandler) run() {
	log.Printf("NetOutHandler is running ...")
	for {
		select {
		case t := <-no.deaman.NetOutTransactionChan:
			no.broadcastTransaction(t)
		case b := <-no.deaman.NetOutBlockChan:
			no.broadcastBlock(b)
		}
	}
}

func (no *NetOutHandler) broadcastTransaction(t *Transaction) {
	//TODO
}

func (no *NetOutHandler) broadcastBlock(b *Block) {
	//TODO
}
