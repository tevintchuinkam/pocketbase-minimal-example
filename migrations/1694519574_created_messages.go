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
			"id": "hk8g3e54fhnaogv",
			"created": "2023-09-12 11:52:54.189Z",
			"updated": "2023-09-12 11:52:54.189Z",
			"name": "messages",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "8j4mkxaj",
					"name": "conversation",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "036j7vu1h2wvd1o",
						"cascadeDelete": true,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": []
					}
				},
				{
					"system": false,
					"id": "yo1ealwv",
					"name": "text",
					"type": "text",
					"required": false,
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
					"id": "1gv3whoe",
					"name": "isUserMessage",
					"type": "bool",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {}
				},
				{
					"system": false,
					"id": "8u1mkxep",
					"name": "image",
					"type": "file",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"maxSize": 5242880,
						"mimeTypes": [
							"image/vnd.mozilla.apng",
							"image/jpeg"
						],
						"thumbs": [],
						"protected": false
					}
				},
				{
					"system": false,
					"id": "5gqkjt4u",
					"name": "type",
					"type": "select",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"values": [
							"text",
							"audio",
							"image"
						]
					}
				},
				{
					"system": false,
					"id": "mtqocorc",
					"name": "wordCount",
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
					"id": "ci4tlqif",
					"name": "audio",
					"type": "file",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"maxSelect": 1,
						"maxSize": 5242880,
						"mimeTypes": [],
						"thumbs": [],
						"protected": false
					}
				}
			],
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_conversation_messages` + "`" + ` ON ` + "`" + `messages` + "`" + ` (` + "`" + `conversation` + "`" + `)"
			],
			"listRule": "@request.auth.id ?= conversation.book.writer",
			"viewRule": "@request.auth.id ?= conversation.book.writer",
			"createRule": "@request.auth.id ?= conversation.book.writer",
			"updateRule": "@request.auth.id ?= conversation.book.writer",
			"deleteRule": "@request.auth.id ?= conversation.book.writer",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("hk8g3e54fhnaogv")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
