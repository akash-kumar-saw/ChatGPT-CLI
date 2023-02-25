# ChatGPT CLI

The ChatGPT CLI is a command-line interface tool that allows you to communicate with the ChatGPT language model. You can use this tool to generate text responses to prompts, have conversations with the model, and explore its capabilities.

## Features
- Easy-to-use CLI interface
- Ability to generate text responses to prompts
- Ability to have conversations with the model
- Support for various parameters to customize the output
- Integration with the OpenAI API for accessing the ChatGPT model

## Installation
- Clone the repository from GitHub:
```bash
git clone https://github.com/akash-kumar-saw/chatgpt-cli
```
- Install the required dependencies:
```bash
cd chatgpt-cli
go mod tidy
```
- Build the CLI tool:
```bash
go build .
```
- Run the CLI tool:
```bash
chatgpt-cli
```

## Configuration

The application requires an OpenAI API key to function. You can obtain an API key by following the instructions in the [OpenAI API documentation](https://platform.openai.com/docs/api-reference/introduction).
Once you have an API key, you can set it using the following commands.
```bash
echo "API_KEY="<OpenAI API Key>"" >> .env
```

## Usage

- Generating text responses:
```bash
***** ChatGPT - Press 'Q' to Exit *****
You : What is the meaning of life?
```

- Having a conversation with the model:
```bash
***** ChatGPT - Press 'Q' to Exit *****
You : Hi, how are you?

I'm fine, thanks. How can I help you?

You : q
***** Thanks for Using ChatGPT *****
```

## Contributing

If you want to contribute to this project, you can submit pull requests, report bugs, or suggest new features.

## Credits

This project was developed by Akash Kumar Saw and is based on the OpenAI API.
