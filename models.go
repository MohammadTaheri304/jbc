package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

//Transaction encapsulate a transaction's data
// From is source account public key
// To is destination account public key
// Sign is signed transaction hash via source account private key
type Transaction struct {
	UUID      string `json:"uuid"`
	From      string `json:"from"`
	To        string `json:"to`
	Value     int64  `json:"value"`
	Tag       string `json:"tag"`
	Timestamp int64  `json:"timestamp"`
	Sign      string `json:"sign"`
}

//Hash compute transaction's hash
func (t *Transaction) Hash() []byte {
	s := t.UUID + t.From + t.To + strconv.FormatInt(t.Value, 10) + t.Tag + strconv.FormatInt(t.Timestamp, 10)
	return Hash(s)
}

// Verify verify transaction sign
func (t *Transaction) Verify(pub string) bool {
	hash := t.Hash()
	publicKey := StringToPublicKey(pub)

	err := VerifySign(hash, t.Sign, publicKey)
	if err == nil {
		return true
	}
	return false
}

//CreateTransaction create a signed transaction
func CreateTransaction(fromPub, fromPriv, toPub string, value int64, tag string) *Transaction {
	trx := &Transaction{
		UUID:      uuid.New().String(),
		From:      fromPub,
		To:        toPub,
		Value:     value,
		Tag:       tag,
		Timestamp: time.Now().UnixNano(),
	}

	hash := trx.Hash()
	privateKey := StringToPrivateKey(fromPriv)
	signeds, _ := Sign(hash, privateKey)
	trx.Sign = signeds
	return trx
}

//Block represent a block od transactions in the blockchain
type Block struct {
	UUID         string         `json:"uuid"`
	Depth        int64          `json:"depth"`
	Timestamp    int64          `json:"timestamp"`
	PrevHash     string         `json:"prevhash"`
	Miner        string         `json:"miner"`
	Transactions []*Transaction `json:"transactions"`
	Nonce        int64          `json:"nonce"`
	Sign         string         `json:"sign"`
}

//Hash compute block's hash
func (b *Block) Hash() []byte {
	sb := &strings.Builder{}
	sb.WriteString(b.UUID)
	sb.WriteString(strconv.FormatInt(b.Depth, 10))
	sb.WriteString(strconv.FormatInt(b.Timestamp, 10))
	sb.WriteString(b.PrevHash)
	sb.WriteString(b.Miner)
	for _, t := range b.Transactions {
		sb.WriteString(Encode64(t.Hash()))
	}

	sb.WriteString(strconv.FormatInt(b.Nonce, 10))

	return Hash(sb.String())
}

//SignIt the block with the given key
func (b *Block) SignIt(priv string) {
	hash := b.Hash()
	privateKey := StringToPrivateKey(priv)
	signeds, _ := Sign(hash, privateKey)
	b.Sign = signeds
}

// VerifyIt verify Block sign
func (b *Block) VerifyIt(pub string) bool {
	hash := b.Hash()
	publicKey := StringToPublicKey(pub)

	err := VerifySign(hash, b.Sign, publicKey)
	if err == nil {
		return true
	}
	return false
}

//CreateUnsignedBlock create an unsigned block
func CreateUnsignedBlock(depth int64, prevHash, minerPubKey string, transactions []*Transaction) *Block {
	return &Block{
		UUID:         uuid.New().String(),
		Depth:        depth,
		Timestamp:    time.Now().Unix(),
		PrevHash:     prevHash,
		Miner:        minerPubKey,
		Transactions: transactions,
	}
}
