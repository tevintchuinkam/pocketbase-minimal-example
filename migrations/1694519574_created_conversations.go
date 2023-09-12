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
			"id": "036j7vu1h2wvd1o",
			"created": "2023-09-12 11:52:54.189Z",
			"updated": "2023-09-12 11:52:54.189Z",
			"name": "conversations",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "qw8ccqtm",
					"name": "done",
					"type": "bool",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {}
				},
				{
					"system": false,
					"id": "nfzkl1sb",
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
				},
				{
					"system": false,
					"id": "qaqjv7ej",
					"name": "book",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "3rcgppnumn009x9",
						"cascadeDelete": true,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": []
					}
				},
				{
					"system": false,
					"id": "cebv0xhf",
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
					"id": "5yalsrgg",
					"name": "wasOpened",
					"type": "bool",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {}
				},
				{
					"system": false,
					"id": "zqtimyzm",
					"name": "topic",
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
					"id": "irhptzd5",
					"name": "topicId",
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
					"id": "j06avows",
					"name": "topicCategory",
					"type": "text",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				}
			],
			"indexes": [
				"CREATE INDEX ` + "`" + `_036j7vu1h2wvd1o_created_idx` + "`" + ` ON ` + "`" + `conversations` + "`" + ` (` + "`" + `created` + "`" + `)"
			],
			"listRule": "",
			"viewRule": "",
			"createRule": null,
			"updateRule": "",
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

		collection, err := dao.FindCollectionByNameOrId("036j7vu1h2wvd1o")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
