package main

import "log"

type Verifier struct {
	deaman *Deaman
}

func StartVerifier(deaman *Deaman) {
	verifier := Verifier{deaman: deaman}
	go verifier.run()
}

// run verifier
func (v *Verifier) run() {
	log.Printf("Verifier is running ...")
	for {
		select {
		case trx := <-v.deaman.NetInTransactionChan:
			v.newTransaction(trx)
		case blk := <-v.deaman.NetInBlockChan:
			v.newBlock(blk)
		}
	}
}

// verifyBlock check block validity
func (v *Verifier) verifyBlock(b *Block) bool {
	// TODO
	// verify its transactions
	// verify its sign
	// verify POW
	return false
}

// verifyTransaction check transaction validity
func (v *Verifier) verifyTransaction(t *Transaction) bool {
	// verify sign -> check duplication -> check feasibility -> check duplication & add to mining set
	//TODO
	return false
}

// newTransaction handle new transaction
func (v *Verifier) newTransaction(t *Transaction) {
	//TODO
}

// newBlock handle new block
func (v *Verifier) newBlock(b *Block) {
	//TODO
}
