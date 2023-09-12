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
			"id": "rfasyz2qq4q0kn1",
			"created": "2023-09-12 11:52:54.188Z",
			"updated": "2023-09-12 11:52:54.188Z",
			"name": "settings",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "kt4hqhll",
					"name": "pricePerPackage",
					"type": "number",
					"required": true,
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
					"id": "bxkb7v23",
					"name": "pricePerBook",
					"type": "number",
					"required": true,
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
					"id": "a9xnk3pc",
					"name": "status",
					"type": "select",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"active",
							"inactive",
							"temporaryDefault"
						]
					}
				}
			],
			"indexes": [],
			"listRule": null,
			"viewRule": null,
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

		collection, err := dao.FindCollectionByNameOrId("rfasyz2qq4q0kn1")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
