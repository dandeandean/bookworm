# BookWorm 📚

**BookWorm** is a command-line tool designed for those who live in the terminal. It helps you manage your bookmarks directly from the command line, making it easy to save, organize, and access your favorite websites without ever leaving your workflow. It's perfect for developers, sysadmins, and anyone who prefers a keyboard-first approach.

***

## Key Features

* **Seamless Integration:** Save, list, and open your bookmarks directly from the terminal.
* **Intuitive Tagging:** Organize your bookmarks with custom tags for easy searching and categorization.
* **Quick Access:** Open bookmarks instantly in your default web browser.
* **Effortless Management:** Easily delete or modify existing bookmarks.
* **Command Completion:** Enjoy command-line completion for faster and more accurate command entry.

***

## Usage

### Commands

**BookWorm** provides a straightforward set of commands to manage your bookmarks. The primary executable is `bookworm`.

| Command | Description | Example |
| :--- | :--- | :--- |
| **`init`** | Initializes the BookWorm configuration and database. | `bookworm init` |
| **`make`** | Creates a new bookmark. | `bookworm make "My Blog" "https://myblog.com" --tag personal` |
| **`list`** | Lists all your saved bookmarks. You can filter by tags. | `bookworm list --tag dev` |
| **`open`** | Opens a saved bookmark in your default browser. | `bookworm open "My Blog"` |
| **`tag`** | Adds or updates tags for a bookmark. | `bookworm tag "My Blog" --add new_tag` |
| **`delete`** | Deletes a bookmark. | `bookworm delete "My Blog"` |
| **`completion`** | Generates shell completion scripts. | `bookworm completion bash` |

### Getting Started

1.  **Installation:** (Provide installation instructions here, e.g., `pip install bookworm` or `go get github.com/user/bookworm`).
2.  **Initialize:** Run `bookworm init` to set up your environment.
3.  **Create:** Add your first bookmark with `bookworm make "Google" "https://google.com"`.
4.  **Open:** Access it later with `bookworm open "Google"`.

***

## Why Use BookWorm?



Traditional browser-based bookmark managers can be slow and disruptive to your workflow, especially if you spend most of your time in the terminal. **BookWorm** streamlines this process by bringing bookmark management to your fingertips, allowing you to stay focused and productive. It's designed to be lightweight, fast, and completely keyboard-driven.
