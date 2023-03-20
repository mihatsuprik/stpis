package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Radser2001/products_api/database"
	"github.com/Radser2001/products_api/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateResponseApprove(t *testing.T) {
	// Define an approveModel for input to the function
	approveModel := models.Approve{
		ID:        1,
		Requestid: 123,
		Docid:     456,
	}

	// Call the function and get the output
	output := CreateResponseApprove(approveModel)

	// Check that the output matches the expected values
	if output.ID != approveModel.ID {
		t.Errorf("Output ID does not match expected value. Got %d, expected %d", output.ID, approveModel.ID)
	}
	if output.Requestid != approveModel.Requestid {
		t.Errorf("Output Requestid does not match expected value. Got %d, expected %d", output.Requestid, approveModel.Requestid)
	}
	if output.Docid != approveModel.Docid {
		t.Errorf("Output Docid does not match expected value. Got %d, expected %d", output.Docid, approveModel.Docid)
	}
}
func TestApproveDocument(t *testing.T) {
	// Create a new Fiber context for the test
	app := fiber.New()
	req := httptest.NewRequest(http.MethodPost, "/approve", strings.NewReader(`{"id":1,"requestid":123,"docid":456}`))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	// Check that the response status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	}

	// Parse the response body and check that it matches the expected values
	var output models.Approve
	if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	expectedOutput := models.Approve{ID: 1, Requestid: 123, Docid: 456}
	if output.ID != expectedOutput.ID || output.Requestid != expectedOutput.Requestid || output.Docid != expectedOutput.Docid {
		t.Errorf("Output does not match expected values. Got %+v, expected %+v", output, expectedOutput)
	}
}

func TestDeleteApproveDocument(t *testing.T) {
	// First, create an approve object to use for the test
	approve := models.Approve{ID: 1, Requestid: 123, Docid: 456}
	if err := database.Database.Db.Create(&approve).Error; err != nil {
		t.Fatalf("Error creating approve object: %v", err)
	}

	// Create a new Fiber context for the test and set the ID parameter to the ID of the approve object we just created
	app := fiber.New()
	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/approve/%d", approve.ID), nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	// Check that the response status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	}

	// Check that the approve object was deleted from the database
	var deletedApprove models.Approve
	if err := database.Database.Db.First(&deletedApprove, approve.ID).Error; err == nil {
		t.Errorf("Expected approve object to be deleted, but it still exists in the database")
	}
}

func TestCreateResponseDocument(t *testing.T) {
	// Create a test document object to use for the test
	documentModel := models.Documents{ID: 1, Name: "Test Document", Content: "This is a test document", Creater: "Test User"}

	// Call the CreateResponseDocument function with the test document object
	output := CreateResponseDocument(documentModel)

	// Check that the output matches the expected values
	expectedOutput := Documents{ID: 1, Name: "Test Document", Content: "This is a test document", Creater: "Test User"}
	if output.ID != expectedOutput.ID || output.Name != expectedOutput.Name || output.Content != expectedOutput.Content || output.Creater != expectedOutput.Creater {
		t.Errorf("Output does not match expected values. Got %+v, expected %+v", output, expectedOutput)
	}
}

func TestCreateDocument(t *testing.T) {
	// Create a test document object to use for the test
	document := models.Documents{Name: "Test Document", Content: "This is a test document", Creater: "Test User"}

	// Encode the document object as JSON and create a new request with the encoded data
	body, err := json.Marshal(document)
	if err != nil {
		t.Fatalf("Error encoding document object as JSON: %v", err)
	}
	req := httptest.NewRequest(http.MethodPost, "/documents", bytes.NewReader(body))

	// Create a new Fiber context and send the request to the CreateDocument function
	app := fiber.New()
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	// Check that the response status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	}

	// Decode the response body and check that it matches the expected output
	var output Documents
	if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
		t.Fatalf("Error decoding response body: %v", err)
	}
	expectedOutput := Documents{Name: "Test Document", Content: "This is a test document", Creater: "Test User"}
	if output.Name != expectedOutput.Name || output.Content != expectedOutput.Content || output.Creater != expectedOutput.Creater {
		t.Errorf("Output does not match expected values. Got %+v, expected %+v", output, expectedOutput)
	}

	// Check that the document object was added to the database
	var dbDocument models.Documents
	if err := database.Database.Db.First(&dbDocument, output.ID).Error; err != nil {
		t.Errorf("Expected document object to be added to database, but it was not found. Error: %v", err)
	}
}

func TestCreateMessage(t *testing.T) {
	// Create a test message object to use for the test
	message := models.Message{Sender: "Test Sender", Recipient: "Test Recipient", Subject: "Test Subject", Body: "This is a test message body."}

	// Encode the message object as JSON and create a new request with the encoded data
	body, err := json.Marshal(message)
	if err != nil {
		t.Fatalf("Error encoding message object as JSON: %v", err)
	}
	req := httptest.NewRequest(http.MethodPost, "/messages", bytes.NewReader(body))

	// Create a new Fiber context and send the request to the CreateMessage function
	app := fiber.New()
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	// Check that the response status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	}

	// Decode the response body and check that it matches the expected output
	var output Message
	if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
		t.Fatalf("Error decoding response body: %v", err)
	}
	expectedOutput := Message{Sender: "Test Sender", Recipient: "Test Recipient", Subject: "Test Subject", Body: "This is a test message body."}
	if output.Sender != expectedOutput.Sender || output.Recipient != expectedOutput.Recipient || output.Subject != expectedOutput.Subject || output.Body != expectedOutput.Body {
		t.Errorf("Output does not match expected values. Got %+v, expected %+v", output, expectedOutput)
	}

	// Check that the message object was added to the database
	var dbMessage models.Message
	if err := database.Database.Db.First(&dbMessage, output.ID).Error; err != nil {
		t.Errorf("Expected message object to be added to database, but it was not found. Error: %v", err)
	}
}
