package main

func main() {
	var (
		m1    = movie{"m1", 10.5}
		m2    = movie{"m2", 20}
		b1    = book{"b1", 12.23, nil}
		b2    = book{"b2", 43, "118281600"}
		b3    = book{"b3", 43, 733622400}
		store list
	)
	store = append(store, &m1, &m2, &b1, &b2, &b3)
	//	movies := []*movie{&m1, &m2}
	//	fmt.Printf("\n %v", movies)
	list := list(store)
	list.print()
	list.discount(.1)

}
