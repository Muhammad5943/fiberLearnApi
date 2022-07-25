package noteHandler

import (
	"fmt"

	"github.com/Muhammad5943/fiberLearnApi/database"
	"github.com/Muhammad5943/fiberLearnApi/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetNotes(c *fiber.Ctx) error {
	db := database.DB
	var notes []model.Note

	//find All notes in the database
	db.Find(&notes)

	// if no note is present return error
	if len(notes) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "No notes present",
			"data":    nil,
		})
	}

	// else return notes
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Notes Found",
		"data":    notes,
	})
}

func CreateNote(c *fiber.Ctx) error {
	db := database.DB
	note := new(model.Note)

	// Store the body in the note and return error if encountered
	err := c.BodyParser(note)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Review your input",
			"data":    err,
		})
	}

	// Add uuid to note
	note.ID = uuid.New()
	// Create the Note and return error if encountered
	err = db.Create(&note).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not create note",
			"data":    err,
		})
	}

	// Return the created note
	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "Created Note",
		"data":    note,
	})
}

func GetNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	// Read the param noteId
	id := c.Params("noteId")

	// Find the note with the given Id
	db.Find(&note, "id = ?", id)

	// If no such note present return an error
	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "No note present",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Notes Found",
		"data":    note,
	})
}

func UpdateNote(c *fiber.Ctx) error {
	type updateNote struct {
		Title    string `json:"title"`
		SubTitle string `json:"sub_title"`
		Text     string `json:"Text"`
	}
	db := database.DB
	var note model.Note

	// Read the param noteId
	id := c.Params("noteId")

	// Find the note with the given Id
	db.Find(&note, "id = ?", id)

	// If no such note present return an error
	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "No note present",
			"data":    nil,
		})
	}

	// Store the body containing the updated data and return error if encountered
	var updateNoteData updateNote
	err := c.BodyParser(&updateNoteData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Review your input",
			"data":    err,
		})
	}

	fmt.Println("note ", note)
	fmt.Println("updateNoteData ", updateNoteData)

	// Edit the note
	note.Title = updateNoteData.Title
	note.SubTitle = updateNoteData.SubTitle
	note.Text = updateNoteData.Text

	// Save the Changes
	db.Save(&note)

	// Return the updated note
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Notes Found",
		"data":    note,
	})
}

func DeleteHandler(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	// Read the param noteId
	id := c.Params("noteId")

	// Find the note with params noteId
	db.Find(&note, "id = ?", id)

	// If no such note present return an error
	if note.ID == uuid.Nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Review your input",
			"data":    nil,
		})
	}

	// Delete the note and return error if encountered
	err := db.Delete(&note, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to delete note",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Deleted Note",
	})
}
