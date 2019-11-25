package main

import "log"

func main() {
	log.Println("Starting JiBiTCoin ...")
	StartDeaman()

	// wait
	<-make(chan bool)
}

type Deaman struct {
	MinerPubKey           string
	MinerPrivKey          string
	NetInTransactionChan  chan *Transaction
	NetInBlockChan        chan *Block
	NetOutTransactionChan chan *Transaction
	NetOutBlockChan       chan *Block
	MinerBlockChan        chan *Block
}

func StartDeaman() {
	d := &Deaman{
		NetInTransactionChan:  make(chan *Transaction, 1024),
		NetInBlockChan:        make(chan *Block, 32),
		NetOutTransactionChan: make(chan *Transaction, 1024),
		NetOutBlockChan:       make(chan *Block, 32),
		MinerBlockChan:        make(chan *Block, 32),
	}

	StartMiner(d)
	StartVerifier(d)
	StartNetInHandler(d)
	StartNetOutHandler(d)
}
