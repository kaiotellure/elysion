package table

type TableItem struct {
	ID string
}

var TABLES = make(map[string][]*TableItem)

// removed items leaves an nil pointer artifact, when looping throught it, remember to check for nil values
func ListTable(key string) []*TableItem {
	current, ok := TABLES[key]
	if !ok {
		TABLES[key] = make([]*TableItem, 0)
	}
	return current
}

func AddItem(key string, id string) {
	ListTable(key) // make sure exists
	TABLES[key] = append(TABLES[key], &TableItem{id})
}

func RemoveItem(key string, id string) {
	for i, v := range TABLES[key] {
		if v != nil && v.ID == id {
			TABLES[key][i] = nil
		}
	}
}

func ContainsItem(key string, id string) bool {
	for _, v := range TABLES[key] {
		if v != nil && v.ID == id {
			return true
		}
	}
	return false
}
