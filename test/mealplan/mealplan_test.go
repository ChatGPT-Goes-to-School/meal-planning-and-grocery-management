package mealplan_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/models"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Mealplan", func() {
	var (
		server   *httptest.Server
		endpoint string
	)

	BeforeEach(func() {
		// Set up any preconditions
		// Start the test server
		server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Handle the API routes
		}))
		endpoint = server.URL + "/api/meal-plans"
	})

	AfterEach(func() {
		// Clean up any resources
		// Stop the test server
		server.Close()
	})

	It("should create a new meal plan", func() {
		// Set up test data and request payload
		payload := []byte(`{
		"Name": "Hi",
		"Description": "Stuff happens",
		"Duration": 4,
		"Meals": "1,2,3,4,5",
		"Username": "franky"
	}`)

		// Make a POST request to the endpoint
		resp, err := http.Post(endpoint, "application/json", bytes.NewReader(payload))
		Expect(err).ToNot(HaveOccurred())
		fmt.Println("Response Status Code:", resp.StatusCode)

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		Expect(err).ToNot(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(http.StatusCreated))

		// Validate the response body
		var createdMealPlan models.MealPlan
		err = json.Unmarshal(body, &createdMealPlan)
		Expect(err).ToNot(HaveOccurred())
		Expect(createdMealPlan.Name).To(Equal("Hi"))
		Expect(createdMealPlan.Description).To(Equal("Stuff happens"))
	})

})
