package example

import (
	"entprom/ent"
	"entprom/ent/hook"
)

func main() {
	client, _ := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	// Add a hook only on user mutations.
	client.User.Use(exampleHook())

	// Add a hook only on update operations.
	client.Use(hook.On(exampleHook(), ent.OpUpdate|ent.OpUpdateOne))
}
