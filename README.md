
# git-pitch

A simple CLI tool written in Go.

## Installation

### From GitHub Releases (recommended)

Download and install the latest prebuilt binary for your system:

```bash
curl -sSfL https://raw.githubusercontent.com/guilhermehbueno/git-pitch/main/scripts/install.sh | bash
````

This will download the latest release for your OS/architecture and place the `git-pitch` binary in `/usr/local/bin`.

After installation, check it with:

```bash
git-pitch --version
```

### From Source (requires Go)

If you have Go installed (>= 1.22):

```bash
go install github.com/guilhermehbueno/git-pitch/cmd/git-pitch@latest
```

This compiles from source and puts the binary into `$GOPATH/bin`.

---

## Usage

```bash
git-pitch --help
```

---

## Extending the CLI with Cobra

This project uses [Cobra](https://github.com/spf13/cobra) to organize commands.
Cobra lets you build a hierarchy of commands and subcommands, keeping `main.go` minimal while putting the real logic in dedicated files.

### Adding a new command

1. Install the generator (optional, but handy):

   ```bash
   go install github.com/spf13/cobra-cli@latest
   ```

2. Add a new command:

   ```bash
   cobra-cli add <command-name>
   ```

   This creates a new file under `cmd/` with a scaffolded `*cobra.Command`.

3. Implement the command logic in that file:

   ```go
   var helloCmd = &cobra.Command{
     Use:   "hello",
     Short: "Say hello",
     Run: func(cmd *cobra.Command, args []string) {
       fmt.Println("Hello from git-pitch!")
     },
   }

   func init() {
     rootCmd.AddCommand(helloCmd)
   }
   ```

4. Rebuild:

   ```bash
   go build -o bin/git-pitch ./cmd/git-pitch
   ```

5. Run your new command:

   ```bash
   git-pitch hello
   ```

### Command hierarchy

* `git-pitch` → root command

    * `init` → initializes configuration
    * `pitch` → runs your pitch logic
    * (your new commands here)

Cobra also automatically provides `--help` and `--version` flags.

---

## Releases

Releases are automatically managed by [release-please](https://github.com/googleapis/release-please-action)
and built with [GoReleaser](https://goreleaser.com/). New versions are created when changes are merged to `main` with Conventional Commit messages (`feat:`, `fix:`, etc.).

```
