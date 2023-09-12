package main

import (
	"fmt"

	_ "example.com/migrations"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func main() {
	app := pocketbase.New()

	// create the minimal setup
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// create a test user 
		userCollection, _ := app.Dao().FindCollectionByNameOrId("users")
		user := models.NewRecord(userCollection)
		user.SetUsername("testuser")
		user.Set("email", "testuser@example.com")
		user.Set("password", "password123")
		err := app.Dao().SaveRecord(user)
		if err != nil{
			return err
		}

		// create the example book
		bookCollection, _ := app.Dao().FindCollectionByNameOrId("books")
		book := models.NewRecord(bookCollection)
		book.Set("title", "test book")
		book.Set("writer", user.GetId())
		book.Set("startDate", "2006-01-02 15:04:05")
		err = app.Dao().SaveRecord(book)
		if err != nil{
			return err
		}

		// create the example conversation
		conversationCollection, _ := app.Dao().FindCollectionByNameOrId("conversations")
		conversation := models.NewRecord(conversationCollection)
		conversation.Set("book", book.GetId())
		conversation.Set("topic", "Test Topic")
		conversation.Set("topicId", "testId")
		conversation.Set("topicCategory", "testCategory")
		err = app.Dao().SaveRecord(conversation)
		if err != nil{
			return err
		}

		// creat example message 
		messageCollection, _ := app.Dao().FindCollectionByNameOrId("messages")
		message := models.NewRecord(messageCollection)
		message.Set("conversation", conversation.GetId())
		message.Set("text", "This is a test message")
		message.Set("type", "text")
		err = app.Dao().SaveRecord(message)
		if err != nil{
			return err
		}

		return nil
	})


	// hooks
	// app.OnModelAfterCreate().Add(hooks.OnModelAfterCreate(app))
	app.OnModelBeforeDelete().Add(func(e *core.ModelEvent) error {
		switch e.Model.TableName() {
		case "messages":
			// subtract the word count from the total number of words written for this conversation (for ai or for user)
			newMessage, err := app.Dao().FindRecordById("messages", e.Model.GetId())
			if err != nil {
				return err
			}
			conversation, err := app.Dao().FindRecordById("conversations", newMessage.GetString("conversation"))
			if err != nil {
				return err
			}
			book, err := app.Dao().FindRecordById("books", conversation.GetString("book"))
			if err != nil {
				return err
			}
			countFieldName := "aiWordCount"
			if newMessage.GetBool("isUserMessage") {
				countFieldName = "userWordCount"
			}

			// subtract the word count from the total number of words written by the ai
			newConversationWordCount := conversation.GetInt(countFieldName) - newMessage.GetInt("wordCount")
			if newConversationWordCount < 0 {
				newConversationWordCount = 0
			}
			conversation.Set(countFieldName, newConversationWordCount)

			// same for the book
			newBookWordCount := book.GetInt(countFieldName) - newMessage.GetInt("wordCount")
			if newBookWordCount < 0 {
				newBookWordCount = 0
			}
			book.Set(countFieldName, newBookWordCount)

			// save the book
			if err := app.Dao().SaveRecord(book); err != nil {
				return err
			}

			// save the conversation
			if err := app.Dao().SaveRecord(conversation); err != nil {
				return err
			}
		}
		return nil
	})

	if err := app.Start(); err != nil {
		fmt.Println(err)
	}
}
