package db

var todoListDB = make(map[string]interface{})

func FindAll() []interface{} {
	items := make([]interface{}, 0, len(todoListDB))
	for _, v := range todoListDB {
		items = append(items, v)
	}

	return items
}

func Save(key string, item interface{}) {
	todoListDB[key] = item
}

func Remove(key string) {
	delete(todoListDB, key)
}
