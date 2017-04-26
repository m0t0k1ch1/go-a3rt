# go-a3rt

[![GoDoc](https://godoc.org/github.com/m0t0k1ch1/go-a3rt?status.svg)](https://godoc.org/github.com/m0t0k1ch1/go-a3rt) [![wercker status](https://app.wercker.com/status/b7779cb2f08a91c25314a364bb9ad4ad/s/master "wercker status")](https://app.wercker.com/project/byKey/b7779cb2f08a91c25314a364bb9ad4ad)

an [A3RT](https://a3rt.recruit-tech.co.jp) API client for golang

## Examples

### Talk API

``` go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	a3rt "github.com/m0t0k1ch1/go-a3rt"
)

func main() {
	client := a3rt.NewClient()
	client.SetApiKey("Your API key")

	res, err := client.SmallTalk(context.Background(), "ご機嫌いかが？")
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
```

``` json
{
  "perplexity": 1.1609984387848817,
  "reply": "そんなに元気じゃありません"
}
```

## Progress

- [ ] Listing API
- [ ] Image Influence API
- [ ] Text Classification API
- [ ] Text Suggest API
- [ ] Proofreading API
- [x] Talk API
