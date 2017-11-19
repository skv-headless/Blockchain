package main

import "fmt"
import "time"
import "crypto/sha256"
import "encoding/hex"
import "encoding/json"

type Transaction struct {
	sender    string
	recipient string
	amount    int
}

type Block struct {
	index         int
	timestamp     int64
	previous_hash string
	proof         int
	transactions  []Transaction
}

type Blockchain struct {
	chain                []Block
	current_transactions []Transaction
}

func (r *Blockchain) New_block(proof int, prev_hash string) {
	//TODO optional parameter
	var previous_hash = ""

	if prev_hash == "1" {
		previous_hash = prev_hash
	} else {
		previous_hash = Hash(r.chain[len(r.chain)-1])
	}

	block := Block{
		index:         len(r.chain),
		timestamp:     time.Now().UnixNano(),
		transactions:  r.current_transactions,
		proof:         proof,
		previous_hash: previous_hash,
	}

	r.current_transactions = []Transaction{}
	r.chain = append(r.chain, block)
}

func (r *Blockchain) New_transaction(sender string, recipient string, amount int) int {
	transaction := Transaction{sender, recipient, amount}
	r.current_transactions = append(r.current_transactions, transaction)
	return r.Last_block().index + 1
}

func (r *Blockchain) Init() {
	r.New_block(100, "1")
}

func (r Blockchain) Last_block() Block {
	return r.chain[len(r.chain)-1]
}

func (r Blockchain) Proof_of_work(last_proof int) int {
	proof := 0
	for {
		if ValidProof(last_proof, proof) {
			break
		}
		proof = proof + 1
	}
	return proof
}

func ValidProof(last_proof int, proof int) bool {
	guess := fmt.Sprintf("%d%d", last_proof, proof)

	sum := sha256.Sum256([]byte(guess))
	guess_hash := hex.EncodeToString(sum[:])

	if guess_hash[0:4] == "0000" {
		return true
	}

	return false
}

func Hash(block Block) string {
	b, err := json.Marshal(block)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return ""
	}

	sum := sha256.Sum256([]byte(b))
	return hex.EncodeToString(sum[:])
}

func main() {
	blockchain := Blockchain{}
	blockchain.Init()
	blockchain.New_transaction("1", "2", 1)

	proof := blockchain.Proof_of_work(100)

	fmt.Printf("%+v\n", proof)
	fmt.Printf("%+v\n", blockchain)
	fmt.Printf("%s\n", Hash(blockchain.chain[0]))
}
