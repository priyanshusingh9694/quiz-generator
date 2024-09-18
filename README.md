# **Multiple Choice Quiz Generator using Generative Language API**

This project is a console-based application written in Go, which generates multiple-choice questions based on a user-specified technology using the Generative Language API. It allows the user to play a quiz-like game, answer the questions, and get a final score.

## **Features**

- Generate multiple-choice questions based on a chosen technology.
- The questions are formatted with four options and an answer provided in the API response.
- Interactive gameplay with feedback on correct and incorrect answers.
- Displays the final score after completing the quiz.

## **Prerequisites**

Before running the application, ensure you have:

- [Go](https://golang.org/doc/install) installed on your machine.
- An API key for the Generative Language API from Google.

## **Getting Started**

### **Clone the repository**

```bash
git clone https://github.com/priyanshusingh9694/quiz-generator.git
cd quiz-generator
go run main.go
```

The application will prompt you to enter your favorite technology and generate a set of questions based on that input.

```bash
What's your favorite technology? React Native

Question 1:
Which of the following is NOT a feature of React Native?
(A) Cross-platform development
(B) Native performance
(C) Hot reloading
(D) Declarative UI
Your answer (A/B/C/D): B
Correct!

Question 2:
...

Game Over! Your final score is 4 out of 5.
```
