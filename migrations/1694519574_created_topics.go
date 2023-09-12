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
			"id": "dmy5hvsggkxgkqd",
			"created": "2023-09-12 11:52:54.189Z",
			"updated": "2023-09-12 11:52:54.189Z",
			"name": "topics",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "xvp58fbw",
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
					"id": "nxkjoh3c",
					"name": "category",
					"type": "select",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"Kindheit & Jugend",
							"Familie",
							"Tägliches Leben",
							"Ausbildung & Beruf",
							"Liebe",
							"Lebensstil und Freizeit",
							"Reisen",
							"Lebensweisheiten",
							"Freundschaft",
							"Besondere Momente"
						]
					}
				}
			],
			"indexes": [
				"CREATE INDEX ` + "`" + `_dmy5hvsggkxgkqd_created_idx` + "`" + ` ON ` + "`" + `topic` + "`" + ` (` + "`" + `created` + "`" + `)"
			],
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

		collection, err := dao.FindCollectionByNameOrId("dmy5hvsggkxgkqd")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
