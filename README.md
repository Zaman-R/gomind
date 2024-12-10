# GoMind - Task Management CLI

GoMind is a CLI tool for managing tasks with due dates and reminders. You can add tasks, list them, and receive notifications when tasks are due.

## Features
- Add new tasks with descriptions and due dates.
- List all tasks with their IDs, descriptions, and due dates.
- Receive reminders when tasks are due.
- **Concurrency**: Uses Go's concurrency model (goroutines and channels) to send reminders asynchronously, ensuring that reminders for tasks are triggered without blocking the main application flow.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/gomind.git
   cd gomind
2. Build the project:
   ```bash
   go build -o gomind

To ADD:
  1. gomind add
      ```bash
      gomind add --description "Your Task Description" --due "YYYY-MM-DD HH:MM:SS"
  2. gomind list
      ```bash
      gomind list

### Future Improvements
- *Add support for marking tasks as completed.*
- *Implement task deletion functionality.*
- *Enhance notification system with different levels of reminders.*


### Key Sections:
- **Features**: Lists the functionalities provided by the project.
- **Installation**: Provides instructions on how to build and run the project.
- **Commands**: Describes the available commands (`add`, `list`), including usage and flags.
- **Storage**: Explains how tasks are stored in the `tasks.json` file.
- **Future Improvements**: Mentions any planned features or improvements.
- **Contributing**: Encourages others to contribute and provides guidance on submitting pull requests.
- **License**: Adds a note about the project license.

Feel free to adjust the content as needed to match your project specifics!
