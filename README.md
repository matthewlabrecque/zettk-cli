# Zettk-CLI

A KISS (Keep It Simple, Stupid) command-line interface for managing a markdown-based Zettelkasten.

## What is a Zettelkasten?

A **Zettelkasten** (German for "slip box") is a note-taking and knowledge management system that originated with German sociologist Niklas Luhmann. The core idea is to store individual notes ("Zettel"), each with a unique identifier, and link them together to form a web of knowledge. Rather than organizing notes hierarchically in folders, a Zettelkasten emphasizes:

- **Atomic notes**: Each note contains one idea.
- **Unique IDs**: Every note is identifiable by a persistent ID (often a timestamp).
- **Bi-directional linking**: Notes reference each other using wiki-style links (e.g., `[[another-note]]`).
- **Emergent structure**: Knowledge organization develops organically through connections rather than pre-defined categories.

## What is Zettk-CLI?

Zettk-CLI is a lightweight, opinionated command-line tool for creating and navigating a markdown-based Zettelkasten. It manages notes in a directory structure under `~/zettlekasten` and integrates directly with your preferred terminal text editor (set via the `EDITOR` environment variable). Zettk-CLI is intentionally minimal — it provides just enough structure to keep a Zettelkasten functional while staying out of your way.

## Directory Structure

After running `zettk-cli init`, your Zettelkasten will look like this:

```
~/zettelkasten/
├── 00-INBOX/              # New notes (default)
├── 01-ARCHIVE/            # Processed notes
│   └── daily-notes/       # Automatic daily notes
├── 02-INPUT/              # External input / reference notes
├── scratchpad/            # Quick notes directory
└── templates/             # Note templates
    └── note.md            # Default template
```

## Commands

### `init`

Initialize a new Zettelkasten in `~/zettelkasten`.

```sh
zettk-cli init
```

Creates the folder structure (`00-INBOX`, `01-ARCHIVE`, `02-INPUT`, `daily-notes`, `templates`) and a default `note.md` template. If a Zettelkasten already exists, it will prompt you before overwriting.

---

### `new [title]`

Create a new note.

```sh
zettk-cli new "my-idea"
```

Generates a new markdown note with a timestamp ID (e.g., `202601021430-my-idea.md`), writes the selected template into it, appends a wiki-link to the current day's daily note, and opens it in your `$EDITOR`.

**Flags:**

- `-t, --template <name>` — Use a custom template from `~/zettelkasten/templates/<name>.md` (default: `note`). If the template name is `input`, the note is saved to `02-INPUT/` instead of `00-INBOX/`.

---

### `open [search]`

Open an existing note in your `$EDITOR`.

```sh
zettk-cli open my-idea
```

Searches `00-INBOX`, `01-ARCHIVE`, and `02-INPUT` for filenames matching the query. If multiple files match, you will be prompted to select one.

---

### `find [search]`

Search for a note and display its metadata.

```sh
zettk-cli find my-idea
```

Returns the filename, Zettelkasten ID, creation time, and last modified time. If multiple files match, you will be prompted to select one.

**Flags:**

- `--inbox` — Search only `00-INBOX`
- `--archive` — Search only `01-ARCHIVE`
- `--input` — Search only `02-INPUT`

*(These flags are mutually exclusive.)*

---

### `sp`

Open the scratchpad for quick, disposable notes.

```sh
zettk-cli sp
```

Opens (or creates) `~/zettelkasten/scratchpad/scratchpad.md` in your `$EDITOR`.

---

### `daily`

Open today's daily note.

```sh
zettk-cli daily
```

Opens (or creates) the daily note at `~/zettelkasten/01-ARCHIVE/daily-notes/YYYY-MM-DD.md` in your `$EDITOR`. New notes created with `zettk-cli new` are automatically linked here.

---

## Building & Installing

### Prerequisites

- [Go](https://go.dev/) 1.26.3 or later
- A text editor set in your `EDITOR` environment variable (e.g., `nvim`, `vim`, `emacs`, `nano`)

### Build locally

```sh
git clone <repository-url>
cd zettk
go build
```

This produces a `zettk-cli` binary in the current directory.

### Install to `$GOPATH/bin`

```sh
go install
```

Ensure `$GOPATH/bin` (or `$HOME/go/bin`) is in your `PATH` to run `zettk-cli` from anywhere.

## License

MIT © 2026 Matthew Labrecque
