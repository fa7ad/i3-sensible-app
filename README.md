# i3-sensible-app

Opens a predefined app based on the currently focused i3 workspace.

## Usage
Change your hotkey from i3-sensible-terminal to i3-sensible-app

```diff
-bindsym Mod4+Return exec i3-sensible-terminal
+bindsym Mod4+Return exec i3-sensible-app
```

Create the default apps config in `~/.config/i3/defaults.json`

Example:

```json
{
  "1: ": ["albert", "toggle"],
  "2: ": ["google-chrome-stable"],
  "3: ": ["i3-sensible-terminal"],
  "4: ": ["code"],
  "5: ": ["nautilus"],
  "6: ?": ["albert", "toggle"],
  "7: ": ["evince"],
  "8: ?": ["albert", "toggle"],
  "9: ": ["xdg-open", "http://localhost:6680/iris"],
  "10: ": ["mpv", "--idle"]
}
```

**Now just shift to your desired workspace and press <kbd>⊞ Win</kbd> + <kbd>Enter ⏎</kbd>**


## Installation

### Pre-requisites

* i3 (and i3-msg)
* go (for compilation)

### INSTALL (with go) (*preferred*)
Open a terminal and run

```bash
go get github.com/fa7ad/i3-sensible-app
```

### INSTALL (from git)
Clone this repo and change into the directory

```bash
git clone https://github.com/fa7ad/i3-sensible-app.git
cd i3-sensible-app/
```

Build the binary

```bash
go build
```

Move the directory to a folder in PATH

```bash
sudo install -Dm755 i3-sensible-app /usr/bin
```

### INSTALL (manual)
Download the binary from [releases](https://github.com/fa7ad/i3-sensible-app/releases)

## LICENSE
MIT License