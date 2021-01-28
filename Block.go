package Block

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
	"unsafe"
)

type Block struct {
	hash         string
	merkeltree   string
	timeStamp    int64
	nonce        int
	genesis      bool
	Transactions []*transaction
}

func (b *Block) hashOfString(DataStr string) string {

	return fmt.Sprintf("%x", sha256.Sum256([]byte(DataStr)))
}

func (b *Block) DeriveHash() {

	info := strconv.Itoa(int(b.timeStamp)) + strconv.Itoa(b.nonce) + b.merkeltree
	hash := b.hashOfString(info)
	b.hash = hash
}

func InitBlock() *Block {

	b := new(Block)
	b.timeStamp = int64(time.Nanosecond) * time.Now().UnixNano()
	b.DeriveHash()
	b.genesis = false

	return b

}

func getMerkleTree(data []*transaction) string {

	//count := unsafe.Sizeof(transaction)

	return ""
}

func main() {
	var T []transaction
	//T.transactionId = "1"
	//T.blockNumber = 2
	//	T.info = "asmaa"

	fmt.Println("this is the length of a transaction", unsafe.Sizeof(T))

}
