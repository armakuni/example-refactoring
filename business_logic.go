package main

func Search(filename string) ([]Item, Paginated, error) {
	items := make([]Item, 0)
	pagination, err := NewFileReader(filename).Paginate(func(index int, name string) {
		items = append(items, Item{index + 1, name})
	})
	return items, pagination, err
}
