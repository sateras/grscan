# grscan

CLI tool for finding Git repositories by remote (name or URL) within a directory.

## Features

* Recursively scans directories
* Detects all `.git` repositories
* Runs `git remote -v` for each repo
* Searches by:

  * remote name (e.g. `person-companies`)
  * URL substring (e.g. `gitlab.khc.kz`)

---

## Installation

```bash id="inst01"
go mod init grscan
go build -o grscan
```

---

## Usage

### Basic usage

```bash id="use01"
./grscan -q person-companies
```

Scans the current directory (`.`)

---

### Specify a path

```bash id="use02"
./grscan -path ~/projects -q person-companies
```

---

## Flags

| Flag    | Description                        | Default            |
| ------- | ---------------------------------- | ------------------ |
| `-path` | root directory to scan             | `.`                |
| `-q`    | search string (remote name or URL) | `person-companies` |

---

## Examples

### Find by remote name

```bash id="ex01"
./grscan -q person-companies
```

---

### Find by URL

```bash id="ex02"
./grscan -q gitlab.khc.kz
```

---

## Example Output

```bash id="out01"
FOUND: /Users/user/projects/repo1
origin  git@gitlab.khc.kz:repo.git (fetch)
origin  git@gitlab.khc.kz:repo.git (push)
-----------
```

---

## How it works

1. Walks the directory tree (`filepath.WalkDir`)
2. Finds `.git` directories
3. For each repo:

   * runs `git remote -v`
   * checks for matches with `-q`
4. Prints matching repositories

---

## Requirements

* Go 1.18+
* `git` installed and available in PATH

---

## Ideas for improvement

* Parallel scanning (faster)
* JSON output
* Branch filtering
* Quiet mode (`--only-paths`)
* Read `.git/config` instead of calling git

---

## License

MIT
