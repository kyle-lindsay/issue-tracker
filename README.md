# issue-tracker

A lightweight issue tracking cli application for local development work.

## Installation:

1. Clone or download this repo
2. Make sure go is installed with `go --version`
3. Run to command `make build` to compile the project

Begin your command with `./issues` in the directory you built the project in. (`issues.exe` for Windows)

## Run Anywhere:

To run `issues` from anywhere in your terminal, add its location to your system PATH.

### macOS & Linux
Move the binary to a standard system directory that is already in your PATH:
```bash
sudo mv issues /usr/local/bin/
```
### Windows
1. Move `issues.exe` to a permanent folder (e.g., `C:\Program Files\issues`).
2. Search for **"environment variables"** in the Windows start menu and select **Edit the system environment variables**.
3. Click **Environment Variables...**, select **Path** under User or System variables, and click **Edit...**.
4. Click **New**, paste the folder path (`C:\Program Files\issues`), and save all windows.
5. Restart your terminal.

### Storage

All issues are saved locally in a human-readable JSON file named `.issues.json` within your current working directory. This makes it easy to track individual sets of issues per repository or project workspace.

## Usage

### Global Options
* `-h, --help`    : Display contextual help and usage instructions for any command.

### Create an Issue

```bash
issues create "Fix bug" --priority high --description "fix the xyz bug"
```

**Syntax:** 
`issues create <title> [options]`

**Options:**
* `-p, --priority <low|medium|high>`: Set the urgency tier. (Defaults to `medium`).
* `-d, --description <string>`: Provide a detailed summary of the issue.

### Close an Issue

```bash
issues close 42
```

**Syntax:** 
`issues close <id>`

### Reopen an Issue

```bash
issues reopen 42
```

**Syntax:** 
`issues reopen <id>`

### List Issues

```bash
issues list --status open --priority high --oldest
```

**Syntax:** 
`issues list [options]`

**Options:**
* `-p, --priority <low|medium|high|any>`: Filters output to selected priority. (Defaults to `any`).
* `-s, --status <open|closed|deleted|any>`: Filters output to selected status. (Defaults to `open`).
* `-l, --limit <number>`: Limits to a specified number of issues listed (Defaults to `5`).
* `-o, --oldest`: Sorts the output by oldest first (Default sort is newest first).

### Get Tracked Issues Status

```bash
issues status
```

**Syntax:** 
`issues status`

**Output:**
Displays information on tracked issues, counting the number of each status, separated by priority.

### Delete an Issue

```bash
issues delete 42 --hard
```

**Syntax:** 
`issues delete <id> [options]`

**Options:**
* `--hard`: Carries out hard deletion. Cannot be undone (Default is soft delete, can be undone).

### Recover an Issue

```bash
issues recover 42
```

**Syntax:** 
`issues recover <id> [options]`

**Options:**
* `-s, --status <open|closed>`: Sets the status of the recovered issue. (Defaults to `open`).

> **NOTE:** Issues deleted with `issues delete <id> --hard` cannot be recovered.

## Licence
*Distributed under the MIT License. See the `LICENSE` file in the repository root for more details.*
