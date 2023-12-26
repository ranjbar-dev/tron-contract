package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	// privateKey := "5e408d58b36aaaef69a91dacd444a4bfc285d0e61fdb147bcb501aed8d5f959d"
	address := "TFdZsjFTqT5APEcWmGo4oeMJ84BGSZgJnH"

	getAccountInfo(address)
}

func getAccountInfo(address string) {

	url := "https://api.shasta.trongrid.io/wallet/getaccount"

	payload := strings.NewReader("{\"address\":\"TZ4UXDV5ZhNW7fb2AMSbgfAEZ7hWsnYS2g\",\"visible\":true}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

}
