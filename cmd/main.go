package main

import (
	"fmt"
)

const address = "TFdZsjFTqT5APEcWmGo4oeMJ84BGSZgJnH"
const privateKey = "5e408d58b36aaaef69a91dacd444a4bfc285d0e61fdb147bcb501aed8d5f959d"

func main() {

	fmt.Println(GetAccountTRC10(address))
}
