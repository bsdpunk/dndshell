package quotes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type Quote struct {
	Index  int    `json:"index"`
	Quote  string `json:"quote"`
	Source string `json:"source"`
}
type Quotes struct {
	Quotes []Quote
}

func (q *Quotes) RandQ() {

	seed := rand.NewSource(time.Now().UnixNano())
	rq := rand.New(seed)
	rqIndex := rq.Intn(len(q.Quotes))
	selected := q.Quotes[rqIndex]
	fmt.Println(selected.Quote, " - ", selected.Source)
}

func (q *Quotes) Load() {
	jsonFile, err := os.Open("./json/quotes.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	if err := json.Unmarshal(byteValue, &q.Quotes); err != nil {
		fmt.Println(err)
	}

	return
}
