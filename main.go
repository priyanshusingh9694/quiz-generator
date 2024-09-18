package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var apiKey = os.Getenv("GEMINI_API_KEY")
const apiURL = "https://generativelanguage.googleapis.com/v1beta/models/gemini-pro:generateContent"

type GenerateTextResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func generateQuestions(technology string, numQuestions int) ([]string, error) {
	questions := []string{}
	for i := 0; i < numQuestions; i++ {
		// Construct the prompt based on user input
		prompt := fmt.Sprintf(`Generate a multiple-choice question with 4 options about %s with answer. all in caps The question should be formatted as follows: 
**Question:** <question_text>
(A) <option_A>
(B) <option_B>
(C) <option_C>
(D) <option_D>
**Answer:** (<correct_option>)`, technology)

		// Define the request payload
		payload := fmt.Sprintf(`{
			"contents": [
				{
					"parts": [
						{
							"text": "%s"
						}
					]
				}
			]
		}`, prompt)

		// Prepare the HTTP request
		reqBody := strings.NewReader(payload)
		req, err := http.NewRequest("POST", fmt.Sprintf("%s?key=%s", apiURL, apiKey), reqBody)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}
		req.Header.Add("Content-Type", "application/json")

		// Send the request
		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("failed to send request: %w", err)
		}
		defer res.Body.Close()

		// Read the response
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}

		// Parse the response
		var response GenerateTextResponse
		if err := json.Unmarshal(body, &response); err != nil {
			return nil, fmt.Errorf("failed to unmarshal response: %w", err)
		}

		if len(response.Candidates) > 0 {
			questionText := response.Candidates[0].Content.Parts[0].Text
			questions = append(questions, questionText)
		}
	}

	return questions, nil
}

func parseQuestion(questionText string) (string, string, string, string, string, string) {
	//fmt.Println(questionText)
	lines := strings.Split(questionText, "\n")

	// Extract the question and options
	question := strings.TrimSpace(lines[0])
	optionA := strings.TrimSpace(strings.TrimPrefix(lines[1], "(A)"))
	optionB := strings.TrimSpace(strings.TrimPrefix(lines[2], "(B)"))
	optionC := strings.TrimSpace(strings.TrimPrefix(lines[3], "(C)"))
	optionD := strings.TrimSpace(strings.TrimPrefix(lines[4], "(D)"))

	// Extract the answer
	answer := extractAnswerOption(lines[5])

	return question, optionA, optionB, optionC, optionD, answer
}

func extractAnswerOption(answerLine string) string {
	// Find the position of the opening and closing parentheses
	start := strings.Index(answerLine, "(")
	end := strings.Index(answerLine, ")")

	// Ensure the parentheses are found and extract the content inside them
	if start != -1 && end != -1 && end > start {
		return strings.TrimSpace(answerLine[start+1 : end]) // Extract and return the letter inside the parentheses
	}

	// Return an empty string if no valid answer was found
	return ""
}

func playGame(questions []string) {
	reader := bufio.NewReader(os.Stdin)
	score := 0
	numQuestions := len(questions)

	for i, question := range questions {
		// Parse the question
		fmt.Printf("\nQuestion %d:\n", i+1)
		questionText, optionA, optionB, optionC, optionD, correctAnswer := parseQuestion(question)

		// Print the question and options
		fmt.Println(questionText)
		fmt.Printf("(A) %s\n(B) %s\n(C) %s\n(D) %s\n", optionA, optionB, optionC, optionD)

		// Get user input
		fmt.Print("Your answer (A/B/C/D): ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		// Check if the user's answer is correct
		if strings.EqualFold(answer, correctAnswer) {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Println("Incorrect!")
			fmt.Printf("The correct answer is: %s\n", correctAnswer)
		}

		fmt.Printf("Your current score: %d out of %d\n", score, i+1)
		fmt.Println("---------------------------")
	}

	// Game finished, show final score
	fmt.Printf("\nGame Over! Your final score is %d out of %d\n", score, numQuestions)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Ask the user for their favorite technology
	fmt.Print("What's your favorite technology? ")
	technology, _ := reader.ReadString('\n')
	technology = strings.TrimSpace(technology)

	// Generate and print the questions
	numQuestions := 5
	questions, err := generateQuestions(technology, numQuestions)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Play the game with the generated questions
	playGame(questions)
}
