# Globetrotter Backend

Globetrotter is a backend service built using Golang and MongoDB to support a travel-themed quiz game. It handles user registration, game sessions, and friend challenges.

## Features
- User Authentication (Login/SignUp)
- User Registration with Unique Username
- Single-Player Mode: Answer destination-related questions based on clues
- Multiplayer Mode: Challenge a friend
- Real-time Score Tracking for every user
- Emoji based Feedback and Fun Fact given for every answer

## Tech Stack
- **Backend:** Golang (Gin Framework)
- **Database:** MongoDB
- **Authentication:** JWT
- **Testing:** Postman

## Setup Instructions

### Prerequisites
Ensure you have the following installed:
- Go (v1.23.2+)
- MongoDB
- Postman (for testing)
- MongoDB Compass (for visualising data)

### Clone the Repository
```sh
git clone https://github.com/yuvraj-2503/globetrotter.git
cd globetrotter
```

### Install Dependencies
```sh
go mod tidy
```

### Setup Environment Variables
Replace the content inside `.env` file in the directory `./globetrotter/config/` with:
```env
SERVER_PORT=8080
MONGO_CONNECTION_STRING=mongodb://localhost:27017
DATABASE=globetrotter
DB_USER=
DB_PASSWORD=
SECRET_KEY=<your_secret_key>
INVITE_BASE_URL=http://localhost:8080/api/v1/globetrotter/users
```

### Run the Server
```sh
go run main.go
```

---

## API Documentation

### Authentication

#### Register User
```http
POST /api/v1/globetrotter/signup
```
**Request Body:**
```json
{
  "email": "john_doe@abc.com",
  "password" : "Pass@123"
}
```
**Response:**
```json
{
  "token": "jwt_token_here"
}
```

#### Login User
```http
POST /api/v1/globetrotter/login
```
**Request Body:**
```json
{
  "email": "john_doe@abc.com",
  "password" : "Pass@123"
}
```
**Response:**
```json
{
  "token": "jwt_token_here"
}
```

### User Management

#### Register User
```http
GET /api/v1/globetrotter/users?username={username}
Authorization: Bearer <token>
```
**Response:**
```http 
Status Code: 200
```

#### Get By Username
```http
GET /api/v1/globetrotter/users?username={username}
Authorization: Bearer <token>
```
**Response:**
```json 
{
    "username": "yuvraj-singh123",
    "score": 20
}
```

#### Get User's Score
```http
GET /api/v1/globetrotter/users/myScore
Authorization: Bearer <token>
```
**Response:**
```json 
{
    "username": "yuvraj-singh123",
    "score": 20
}
```

### Game Session

#### Get Random Question
```http
GET /api/v1/globetrotter/question
Authorization: Bearer <token>
```
**Response:**
```json 
{
    "questionId": "67d6508ba6876329fa88af0b",
    "clue": "Home to Table Mountain and stunning coastal scenery.",
    "options": [
        "Johannesburg",
        "Durban",
        "Cape Town",
        "Pretoria"
    ]
}
```

#### Submit Answer
```http
POST /api/v1/globetrotter/answer
Authorization: Bearer <token>
```
**Request Body:**
```json
{
  "questionId" : "67d6508ba6876329fa88af18",
  "answer" : "Bangalore"
}
```
**Response:**
```json
{
  "feedback": "🎉",
  "funFact": "Bangalore has a pleasant climate year-round due to its elevation."
}
```

### Challenge a Friend

#### Generate Invite Link
```http
GET /api/v1/globetrotter/invite?invitee={username}
Authorization: Bearer <token>
```
**Response:**
```json
{
  "inviteLink": "http://localhost:8081/api/v1/globetrotter/users?username=sushil"
}
```

#### Get Inviter Score
```http
GET /api/v1/globetrotter/inviter/score
Authorization: Bearer <token>
```
**Response:**
```json
{
  "inviter": "yuvraj-singh",
  "score": 20
}
```

## License
This project is open-source under the MIT License.

