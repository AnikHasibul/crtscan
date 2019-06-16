package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/anikhasibul/queue"
)

// CRTLog holds a log
type CRTLog struct {
	IssuerCaID        int    `json:"issuer_ca_id"`
	IssuerName        string `json:"issuer_name"`
	MinCertID         int    `json:"min_cert_id"`
	MinEntryTimestamp string `json:"min_entry_timestamp"`
	NameValue         string `json:"name_value"`
	NotAfter          string `json:"not_after"`
	NotBefore         string `json:"not_before"`
}

// CRTLogs holds all log
type CRTLogs []*CRTLog

var q = queue.New(5)
var nameList = make(map[string]bool)

func main() {
	q.Add()
	if len(os.Args) != 2 {
		die(errors.New("Invalid argument given. Please read the instructions:" +
			"https://github.com/AnikHasibul/crtscan"))
	}
	scan(os.Args[1])
	q.Wait()
}
func scan(uri string) {
	defer q.Done()
	resp, err := http.Get(fmt.Sprintf("https://crt.sh/?q=%v&output=json", uri))
	if err != nil {
		die(err)
		return
	}
	c := new(CRTLogs)
	c.Decode(resp.Body)
	for _, v := range *c {
		v.Scan()
	}
}
func die(err error) {
	log.Fatal(err)
}

// Decode decodes the json
func (c *CRTLogs) Decode(d io.ReadCloser) {
	defer d.Close()
	err := json.NewDecoder(d).Decode(c)
	if err != nil {
		die(err)
		return
	}
}

// Scan takes action as per the instruction
func (c *CRTLog) Scan() {
	if c.IsAlreadyListed() {
		return
	}
	if c.IsWildcard() {
		return
	}
	c.PrintHost()
	q.Add()
	go scan(c.NameValue)

}

// PrintHost prints the hostname
func (c *CRTLog) PrintHost() {
	fmt.Println(c.NameValue)
}

// IsWildCard whether the name value is a wildcard
func (c *CRTLog) IsWildcard() bool {
	return c.NameValue[0] == '*'
}

// IsAlreadyListed whether the domain or wildcard is already discovered
func (c *CRTLog) IsAlreadyListed() bool {
	_, ok := nameList[c.NameValue]
	if !ok {
		nameList[c.NameValue] = true
	}
	return ok
}
