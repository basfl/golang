package main

import (
	"fmt"
	"strconv"
	"strings"
)

type result struct {
	domain string
	visits int
}
type parser struct {
	sum     map[string]result
	domains []string
	total   int
	lines   int
}

func newParser() parser {
	return parser{sum: make(map[string]result)}
}

func parse(p *parser, line string) (parsed result, err error) {
	// Parse the fields
	// var (
	// 	parsed result
	// 	err    error
	// )

	fields := strings.Fields(line)
	if len(fields) != 2 {
		err = fmt.Errorf("wrong input: %v (line #%d)", fields, p.lines)
		//	return parsed, err
		return
	}

	parsed.domain = fields[0]

	// Sum the total visits per domain
	parsed.visits, err = strconv.Atoi(fields[1])
	if parsed.visits < 0 || err != nil {
		err = fmt.Errorf("wrong input: %q (line #%d)", fields[1], p.lines)
		//	return parsed, err
		return
	}
	//return parsed, err
	return
}

func update(p *parser, parsed result) {

	// Collect the unique domains
	if _, ok := p.sum[parsed.domain]; !ok {
		p.domains = append(p.domains, parsed.domain)
	}

	// Keep track of total and per domain visits
	p.total += parsed.visits

	r := result{
		domain: parsed.domain,
		visits: parsed.visits + p.sum[parsed.domain].visits,
	}
	p.sum[parsed.domain] = r

}
