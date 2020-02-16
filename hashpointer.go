package assignment01bca

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	transaction   string
	prevBlock     *Block
	prevBlockHash string
}

func getHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func InsertBlock(transaction string, chainHead *Block) *Block {
	var newHead *Block
	if chainHead == nil {
		newHead = &Block{transaction: transaction, prevBlock: chainHead, prevBlockHash: ""}
		newHead = &Block{transaction: "", prevBlock: newHead, prevBlockHash: getHash(transaction)}
	} else {
		chainHead.transaction = transaction
		newHead = &Block{transaction: "", prevBlock: chainHead, prevBlockHash: getHash(transaction)}
	}
	return newHead
}

func listBlocks_(chainHead *Block) {
	if chainHead == nil {
		return
	}
	listBlocks_(chainHead.prevBlock)
	fmt.Println(chainHead.transaction)
}

func ListBlocks(chainHead *Block) {
	listBlocks_(chainHead.prevBlock)
	fmt.Println()
}
func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {
	tmpHead := chainHead
	for tmpHead != nil {
		if tmpHead.transaction == oldTrans {
			tmpHead.transaction = newTrans
			break
		}
		tmpHead = tmpHead.prevBlock
	}
}

func VerifyChain(chainHead *Block) {
	tmpHead := chainHead
	for tmpHead != nil {
		if getHash(tmpHead.prevBlock.transaction) != tmpHead.prevBlockHash {
			fmt.Println("Chain has been tempered: This transaction didn't belonged in the chain " + tmpHead.prevBlock.transaction)
			break
		}
		tmpHead = tmpHead.prevBlock
	}
}
