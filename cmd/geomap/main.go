package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"

	"github.com/cydave/geomap/internal"
)

func main() {
	ls, err := internal.NewLookupService(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer ls.Close()

	items := make([]internal.LookupResult, 0)
	scn := bufio.NewScanner(os.Stdin)
	for scn.Scan() {
		line := scn.Text()
		line = strings.TrimSpace(line)
		ip := net.ParseIP(line)
		if ip == nil {
			continue
		}
		if ok := ip.To4(); ok == nil {
			continue
		}

		res, err := ls.Lookup(ip)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, *res)
	}

	if err := scn.Err(); err != nil {
		log.Fatal(err)
	}

	internal.RenderResults(items)
}
