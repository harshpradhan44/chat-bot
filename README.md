# Chat Bot

This is a simple chat bot written in Golang that responds to different intents via HTTP endpoints. The bot supports three primary intents: `chat`, `image_upload`, and `image_download`.

## Endpoints

### `GET /bot`

This endpoint handles various intents based on the query parameters provided:

- `intent`: Specify the intent of the interaction (`chat`, `image_upload`, or `image_download`).
- `input`: The user input or message for the bot.

#### Examples

1. **Chat Intent**:
   ```
   GET /bot?intent=chat&input=how are you?
   ```
   This triggers a conversation with the chat bot. Replace `how are you?` with your desired message.

2. **Image Upload Intent**:
   ```
   GET /bot?intent=image_upload
   ```
   For uploading an image, send a POST request to this endpoint with the image under the `file` key within form-data.

3. **Image Download Intent**:
   ```
   GET /bot?intent=image_download
   ```
   This intent fetches an image from the bot.

### `POST /bot` (For Image Upload)

For uploading images, use this endpoint with a POST request. Include the image file within the `file` key in form-data.

## Getting Started

### Prerequisites

- Go installed on your system.
- Dependencies specified in the project's `go.mod` file.

### Installation

Clone this repository to your local machine:

```bash
git clone https://github.com/your-repo-name.git
```

### Running the Application

1. Navigate to the project directory:

   ```bash
   cd golang-basic-chat-bot
   ```

2. Build and run the application:

   ```bash
   go build -o chat-bot
   ./chat-bot
   ```

This will start the server at `localhost:8080`.

## Contributing

Feel free to contribute to this project by creating issues or submitting pull requests.
