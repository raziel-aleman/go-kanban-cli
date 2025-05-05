# Simple Kanban CLI

A simple terminal-based Kanban board application built with Go and the Bubble Tea TUI library. Organize your tasks in a To Do, In Progress, and Done column directly from your command line.

## Features

* Task management: Add, move, and delete tasks.
* Column navigation: Easily switch between To Do, In Progress, and Done columns.
* Persistence: Saves your board state to a file so your tasks are still there next time you run the app.
* Filtering: Quickly find tasks within a column.
* Clean TUI interface using Lipgloss.

## Installation

1.  **Prerequisites:** Make sure you have Go installed (`go version`). If not, follow the instructions on the [official Go website](https://golang.org/doc/install).
2.  **Clone the repository:**
    ```bash
    git clone https://github.com/raziel-aleman/go-kanban-cli.git
    cd go-kanban-cli
    ```

3.  **Run the application:**
    ```bash
    go run .
    ```
    The first time you run it, a storage file (details depend on your `ReadFromStorage`/`WriteToStorage` implementation) might be created.

## Usage

The application uses simple keybindings for interaction:

* **Left Arrow (`←`) / Right Arrow (`→`):** Change focus between the To Do, In Progress, and Done columns.
* **Up Arrow (`↑`) / Down Arrow (`↓`):** Navigate through the tasks within the currently focused column.
* **Enter (`⏎`):** Move the selected task to the next column (To Do -> In Progress -> Done -> To Do).
* **n:** Initiate adding a new task (this will likely switch to an input form, assuming you have implemented that part).
* **x:** Delete the selected task from the current column.
* **/:** Enable filtering for the currently focused list. Start typing to filter tasks.
* **Ctrl+C / q:** Quit the application. Your board state will be saved automatically on quitting.

## Project Structure

* `main.go`: Entry point of the cli.
* `model.go`: Contains the main application logic, the Bubble Tea model (`Kanban`), and the `Update`, `View`, and `Init` methods.
* `task.go`: Defines the `Task` struct and related methods.
* `form.go`: Defines the `Form` struct and related methods.
* `datastore.go`: Defines logic to handle to read from storage and write to storage.
* `storage.json`: Prersistent storage.
* `constants.go`: Defines keybindings like `QuitKeys`.

*(Adjust these assumed files based on your actual project structure)*

## Dependencies

This project relies on the following Go modules:

* `github.com/charmbracelet/bubbletea`
* `github.com/charmbracelet/bubbles/key`
* `github.com/charmbracelet/bubbles/list`
* `github.com/charmbracelet/lipgloss`

These dependencies will be automatically downloaded when you run `go run .` or build the project (`go build`).

## Contributing

(Add information on how others can contribute, if applicable)

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.