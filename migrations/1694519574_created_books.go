package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "3rcgppnumn009x9",
			"created": "2023-09-12 11:52:54.188Z",
			"updated": "2023-09-12 11:52:54.188Z",
			"name": "books",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "4gjshohi",
					"name": "title",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "wwu4f3ir",
					"name": "writer",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": true,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": []
					}
				},
				{
					"system": false,
					"id": "japtxllv",
					"name": "asker",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": []
					}
				},
				{
					"system": false,
					"id": "cklnbbky",
					"name": "startDate",
					"type": "date",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": "",
						"max": ""
					}
				},
				{
					"system": false,
					"id": "9gjwlh1g",
					"name": "order",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "p2jgfsc737eviik",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": []
					}
				},
				{
					"system": false,
					"id": "l3xwsx2z",
					"name": "firstQuestionSent",
					"type": "bool",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {}
				},
				{
					"system": false,
					"id": "cqc49pwf",
					"name": "aiWordCount",
					"type": "number",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"noDecimal": false
					}
				},
				{
					"system": false,
					"id": "uokpb2e7",
					"name": "userWordCount",
					"type": "number",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"noDecimal": false
					}
				}
			],
			"indexes": [
				"CREATE INDEX ` + "`" + `_3rcgppnumn009x9_created_idx` + "`" + ` ON ` + "`" + `books` + "`" + ` (` + "`" + `created` + "`" + `)"
			],
			"listRule": "asker.id = @request.auth.id || writer.id = @request.auth.id",
			"viewRule": "asker.id = @request.auth.id || writer.id = @request.auth.id",
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("3rcgppnumn009x9")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
