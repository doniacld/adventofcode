package bags

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
	regSourceBag   = regexp.MustCompile("^([a-z]+ [a-z]+) bags contain (.*).")
	regContentBags = regexp.MustCompile("([0-9]+) ([a-z]+ [a-z]+) bags?")
)

// ParseLine takes a line from the input and stores all the bags with their associated
// source bags and the number of each one
func (r Bags) ParseLine(s string) {
	// extract the source bag
	rule := regSourceBag.FindStringSubmatch(s)

	// extract the bag content
	if rule[2] == "no other bags" {
		return
	}
	content := regContentBags.FindAllStringSubmatch(rule[2], -1)

	// store all the bags
	for _, l := range content {
		// extract the number of bags
		nb, err := strconv.Atoi(l[1])
		if err != nil {
			log.Fatal(err)
		}
		// extract the bag name
		leave := strings.Join(l[2:], "")
		// insert the content bags as key and the source bag as value
		r.insertParentBag(leave, bag{nb, rule[1]})
		r.insertChildBag(rule[1], bag{nb, leave})
	}
}

// insertParentBag insert a parent bag into the map
func (r Bags) insertParentBag(name string, parent bag) {
	v, ok := r[name]
	if !ok {
		r[name] = newBagFamily()
	}
	v.parents = append(v.parents, parent)
	r[name] = v
}

// insertChildBag inserts a child bag into the map
func (r Bags) insertChildBag(name string, child bag) {
	v, ok := r[name]
	if !ok {
		r[name] = newBagFamily()
	}
	v.children = append(v.children, child)
	r[name] = v
}
