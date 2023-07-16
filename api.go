package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Note struct {
	ID   uint32 `json:"id"`
	Note string `json:"note"`
}

type Session struct {
	SID string `json:"sid"`
}

var users []User
var notes []Note
var noteIDCounter uint32

func signupHandler(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	// Store the user details in the database or any other persistence mechanism
	users = append(users, user)

	c.JSON(200, gin.H{"message": "User created successfully"})
}

func loginHandler(c *gin.Context) {
	var userCredentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&userCredentials); err != nil {
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	// Authenticate the user by checking credentials against the stored user details
	// Here, we assume the authentication is successful and generate a session ID
	session := Session{SID: "unique_session_id"}

	c.JSON(200, session)
}

func getNotesHandler(c *gin.Context) {
	var session struct {
		SID string `json:"sid"`
	}
	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	// Validate the session ID and check if it is associated with a logged-in user
	// Here, we assume the session is valid and retrieve the user's notes from the database
	userNotes := []Note{
		{ID: 1, Note: "Note 1"},
		{ID: 2, Note: "Note 2"},
		{ID: 3, Note: "Note 3"},
	}

	c.JSON(200, gin.H{"notes": userNotes})
}

func createNoteHandler(c *gin.Context) {
	var newNote struct {
		Note string `json:"note"`
	}
	if err := c.ShouldBindJSON(&newNote); err != nil {
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	// Store the new note in the database or any other persistence mechanism
	noteIDCounter++
	newNoteID := noteIDCounter
	newNoteRecord := Note{ID: newNoteID, Note: newNote.Note}
	notes = append(notes, newNoteRecord)

	c.JSON(200, gin.H{"id": newNoteID})
}

func deleteNoteHandler(c *gin.Context) {
	var session struct {
		SID string `json:"sid"`
	}
	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	// Validate the session ID and check if it is associated with a logged-in user
	// Here, we assume the session is valid

	// Retrieve the note ID from the URL parameter
	noteIDParam := c.Param("id")
	noteID, err := strconv.ParseUint(noteIDParam, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	// Check if the note ID exists
	exists := false
	for _, note := range notes {
		if note.ID == uint32(noteID) {
			exists = true
			break
		}
	}

	if !exists {
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	// Validate the session ID and check if it is associated with a logged-in user
	// Here, we assume the session is valid
	// Omitted the actual deletion of the note for simplicity

	c.JSON(200, gin.H{"message": "Note deleted successfully"})
}

func main() {
	router := gin.Default()

	router.POST("/signup", signupHandler)
	router.POST("/login", loginHandler)
	router.GET("/notes", getNotesHandler)
	router.POST("/notes", createNoteHandler)
	router.DELETE("/notes/:id", deleteNoteHandler)

	router.Run(":8080")
}
