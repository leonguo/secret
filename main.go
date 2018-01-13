package coinrua

import (
	"math/big"
	"fmt"
	"crypto/sha256"
)

func main() {

	data1 := []byte("i like dd")
	data2 := []byte("i like ddd22241f341")
	targetBits := 24
	target := big.NewInt(1)
	target.Lsh(target,uint(256-targetBits))
	fmt.Printf("%x\n", sha256.Sum256(data1))
	fmt.Printf("%64x\n", target)
	fmt.Printf("%x\n", sha256.Sum256(data2))
}