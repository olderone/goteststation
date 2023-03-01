package clickhouse

import (
	"fmt"
)

func (c *Ch)GetUsers() error {
	defer func() {
		_ = c.db.Close()
	}()

	var items []User
	if err := c.db.Select(&items, "select * from max.user limit 10"); err != nil {
		fmt.Printf("clickhouse connect err %s", err.Error())
		return err
	}
	for _, item := range items {
		fmt.Printf("id: %d, name: %s, age: %d\n", item.Id, item.Name, item.Age)
	}
	return nil
}