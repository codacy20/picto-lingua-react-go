# Picto Lingua - Image-Based Language Learning

Picto Lingua is a language learning application that helps users learn vocabulary words in context through images. The application presents users with images representing different themes, and then provides vocabulary words related to those themes.

## Features

- Theme-based learning with visual context
- Multiple language support (English and Dutch)
- Flashcard game mode for vocabulary practice
- Image selection from Unsplash API
- Vocabulary generated through OpenAI
- Session-based progress tracking

## Tech Stack

- **Frontend**: React
- **Backend**: Golang
- **APIs**: Unsplash, OpenAI

## Getting Started

### Prerequisites

- Go 1.16+
- Node.js 14+
- npm or yarn
- Unsplash API key
- OpenAI API key

### Backend Setup

1. Clone the repository
   ```
   git clone https://github.com/yourusername/picto-lingua-react-go.git
   cd picto-lingua-react-go
   ```

2. Set up environment variables
   ```
   cp .env.example .env
   ```
   Edit the `.env` file to add your Unsplash and OpenAI API keys.

3. Run the backend
   ```
   cd backend
   go run main.go
   ```
   The server will start on port 8080 (or the port specified in your .env file)

### Frontend Setup

1. Install dependencies
   ```
   cd frontend
   npm install
   ```

2. Start the development server
   ```
   npm start
   ```
   The React app will be available at http://localhost:3000

## API Endpoints

### Images

- `GET /api/images?theme=<theme>` - Get images for a specific theme

### Vocabulary

- `GET /api/vocabulary?theme=<theme>&count=<count>&language=<language>` - Get vocabulary words for a specific theme and language
  - `language` parameter can be "english" (default) or "dutch"

### Sessions

- `POST /api/session` - Create or update a session
- `GET /api/session?session_id=<session_id>` - Get a session by ID

### Themes

- `GET /api/themes` - Get all available themes

## Project Structure

```
picto-lingua-react-go/
├── backend/                      # Go backend
│   ├── api/
│   │   ├── handlers/             # API endpoint handlers
│   │   ├── models/               # Data models
│   │   └── services/             # Service layer (Unsplash, OpenAI, etc.)
│   ├── config/                   # Configuration management
│   ├── utils/                    # Utility functions
│   ├── main.go                   # Application entry point
│   ├── go.mod                    # Go module definition
│   └── go.sum                    # Go module checksums
├── frontend/                     # React frontend
│   ├── public/                   # Static files
│   └── src/
│       ├── components/           # React components
│       ├── services/             # API service layer
│       └── App.js                # Main application component
├── .env                          # Environment variables
├── .gitignore                    # Git ignore rules
└── README.md                     # Project documentation
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgements

- [Unsplash](https://unsplash.com/) for providing beautiful, free images
- [OpenAI](https://openai.com/) for powering the vocabulary generation # picto-lingua-react-go
