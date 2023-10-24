# sdPreCommit

`sdPreCommit` is a tool designed to be a hard-coded pre-commit equivalent for specific use cases. It automates the process of running `yarn lint` and checks for the presence of specific comment prefixes in your codebase before committing.

## Getting Started

### Prerequisites

- Go installed on your machine
- Your project should be using `yarn` for package management

### Installation

Clone the repository and navigate to its directory in your terminal.

### Run

To execute the program, run:

```bash
go run main.go
```

### Build

To build an executable, run:

```bash
go build main.go
```

### Usage
Run sdPreCommit in your repository before making a commit. It will check your staged files against your configuration settings for linting and comment prefix checking.

### Customization

- File Types: Modify the var interestedExtensions in main.go to include the file extensions you're interested in checking.
```bash
var interestedExtensions = []string{".js", ".jsx", ".ts", ".tsx", ".vue", ".scss", ".css"}
```

- Comment Prefix: Change the const commentString to set your desired comment prefix that you'd like to search for in the codebase.
```bash
const commentString = "TODO_CMNT"
```

- Linting Commands: At the runEslint function, you can specify which yarn commands you'd like to run. Alternatively, you can define your own function and call it from main().
```bash
func runEslint(files []string) {
  // Your linting logic here
}
```
