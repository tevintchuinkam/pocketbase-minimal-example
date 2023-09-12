package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("id = @request.auth.id || id ?= @request.auth.books.asker.id")

		collection.ViewRule = types.Pointer("id = @request.auth.id || id ?= @request.auth.books.asker.id")

		json.Unmarshal([]byte(`[
			"CREATE INDEX ` + "`" + `__pb_users_auth__created_idx` + "`" + ` ON ` + "`" + `users` + "`" + ` (` + "`" + `created` + "`" + `)",
			"CREATE INDEX ` + "`" + `idx_unique_au1t9bz7` + "`" + ` ON ` + "`" + `users` + "`" + ` (` + "`" + `books` + "`" + `)"
		]`), &collection.Indexes)

		// remove
		collection.Schema.RemoveField("users_name")

		// add
		new_books := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "au1t9bz7",
			"name": "books",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "3rcgppnumn009x9",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": null,
				"displayFields": []
			}
		}`), new_books)
		collection.Schema.AddField(new_books)

		// add
		new_oneTimePassword := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "iz0ykc8m",
			"name": "oneTimePassword",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_oneTimePassword)
		collection.Schema.AddField(new_oneTimePassword)

		// add
		new_overWrittenOTP := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "txx1jubh",
			"name": "overWrittenOTP",
			"type": "bool",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), new_overWrittenOTP)
		collection.Schema.AddField(new_overWrittenOTP)

		// add
		new_firstName := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "1gbvy29p",
			"name": "firstName",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_firstName)
		collection.Schema.AddField(new_firstName)

		// add
		new_lastName := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "b0rabw7q",
			"name": "lastName",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_lastName)
		collection.Schema.AddField(new_lastName)

		// add
		new_stripeCustomerID := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "q4gwigyb",
			"name": "stripeCustomerID",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), new_stripeCustomerID)
		collection.Schema.AddField(new_stripeCustomerID)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("id = @request.auth.id")

		collection.ViewRule = types.Pointer("id = @request.auth.id")

		json.Unmarshal([]byte(`[]`), &collection.Indexes)

		// add
		del_name := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "users_name",
			"name": "name",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), del_name)
		collection.Schema.AddField(del_name)

		// remove
		collection.Schema.RemoveField("au1t9bz7")

		// remove
		collection.Schema.RemoveField("iz0ykc8m")

		// remove
		collection.Schema.RemoveField("txx1jubh")

		// remove
		collection.Schema.RemoveField("1gbvy29p")

		// remove
		collection.Schema.RemoveField("b0rabw7q")

		// remove
		collection.Schema.RemoveField("q4gwigyb")

		return dao.SaveCollection(collection)
	})
}
