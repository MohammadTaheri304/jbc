package main

import (
	"github.com/theckman/go-securerandom"
	"log"
	"math/rand"
	"strings"
)

type Miner struct {
	deaman          *Deaman
	randomGenerator *rand.Rand
}

func StartMiner(deaman *Deaman) {
	ri64, err := securerandom.Int64()
	if err != nil {
		log.Fatalf("Error in starting miner, %+v\n", err)
	}

	ra := rand.New(rand.NewSource(ri64))
	miner := Miner{
		deaman:          deaman,
		randomGenerator: ra,
	}
	go miner.run()
}

func (m *Miner) run() {
	log.Printf("Miner is running ...")
	var current *Block
	for {
		select {
		case newBlock := <-m.deaman.MinerBlockChan:
			current = newBlock
		default:
			if current != nil {
				m.mine(current)
			}
		}
	}
}

func (m *Miner) mine(block *Block) {
	block.Miner = m.deaman.MinerPubKey
	powSize := powSize(block)
	log.Printf("POW size is %+v\n", powSize)
	block.Nonce = m.randomGenerator.Int63()
	bHash := block.Hash()
	if !checkPow(bHash, powSize) {
		log.Println("POW Failed!")
		return
	}

	log.Println("POW was successful!")
	block.SignIt(m.deaman.MinerPrivKey)
	m.broadcastBlock(block)
}

func powSize(block *Block) int {
	matchSize := findMatchSize(block.PrevHash, block.Miner)
	transactionFactor := len(block.Transactions) / 256
	return 36 - matchSize - transactionFactor
}

func findMatchSize(preHash, minerPubKey string) int {
	s1 := strings.ToLower(preHash)
	s2 := strings.ToLower(minerPubKey)
	matchSize := 0
	for i := 0; i < 16; i++ {
		if s1[i] == s2[i] {
			matchSize++
		} else {
			break
		}
	}
	return matchSize
}

func checkPow(hash []byte, size int) bool {
	for i := 0; i < size; i++ {
		if hash[i] != 0 {
			return false
		}
	}

	return true
}

func (m *Miner) broadcastBlock(block *Block) {
	m.deaman.NetOutBlockChan <- block
}
