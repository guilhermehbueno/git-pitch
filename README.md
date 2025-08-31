# Pitch CLI - AI-Powered Git Workflow Tool

Transform your Git workflow with AI-generated commit messages and pull request descriptions. Pitch CLI integrates with Ollama to automatically create meaningful, contextual Git messages from your code changes.

## ✨ Features

- 🤖 **AI-Powered Commit Messages** - Automatically generate descriptive commit messages from staged changes
- 📝 **Pull Request Generation** - Create comprehensive PR titles and descriptions with structured markdown
- 🧠 **Model Management** - Easy installation and switching between different AI models
- 🔧 **Git Integration** - Seamless integration with Git hooks for automatic message generation
- 💬 **AI Consultation** - Interactive AI assistant for code questions and explanations
- 📚 **Documentation Generation** - Automatic README.md creation from project analysis
- ⚡ **Rich Terminal UI** - Beautiful, interactive interface with real-time feedback

## 🚀 Quick Start

### Installation

#### From GitHub Releases (Recommended)

Download and install the latest prebuilt binary for your system:

```bash
curl -sSfL https://raw.githubusercontent.com/guilhermehbueno/git-pitch/main/scripts/install.sh | bash
```

After installation, verify it works:

```bash
pitch --version
```

#### From Source (Requires Go >= 1.22)

```bash
go install github.com/guilhermehbueno/git-pitch@latest
```

### Initial Setup

1. **Install Ollama and configure models:**
   ```bash
   pitch install
   ```
   This will install Ollama, download AI models, and configure your system.

2. **Apply Git hooks to your repository:**
   ```bash
   cd your-git-repo
   pitch apply
   ```

3. **Start generating AI commit messages:**
   ```bash
   git add .
   pitch commit
   ```

## 📋 Main Commands

### Core Workflow Commands

| Command | Description | Example |
|---------|-------------|---------|
| `pitch commit` | Generate AI commit message from staged changes | `pitch commit -m "additional context"` |
| `pitch pr <branch>` | Create PR title and description | `pitch pr main --draft` |
| `pitch ask [query]` | Interactive AI consultation | `pitch ask "explain this function"` |

### Model Management

| Command | Description |
|---------|-------------|
| `pitch model list` | Show available AI models |
| `pitch model select` | Interactive model selection |
| `pitch model create <name>` | Create custom model with guided setup |

### System Management

| Command | Description |
|---------|-------------|
| `pitch install` | Complete setup (Ollama, models, configuration) |
| `pitch info` | System status and diagnostics |
| `pitch start/stop` | Control Ollama server |
| `pitch apply` | Install Git hooks for automatic commit messages |

### Utilities

| Command | Description |
|---------|-------------|
| `pitch readme` | Generate project documentation |
| `pitch update` | Update to latest version |

## ⚙️ Configuration

Pitch CLI looks for configuration files in this order:

1. `.git/git-pitch.yaml` (project-specific)
2. `~/git-pitch.yaml` (user-specific)
3. `/etc/git-pitch/git-pitch.yaml` (system-wide)

### Example Configuration

```yaml
# ~/.git-pitch.yaml
ollama:
  model: "llama3.2:latest"
  server_url: "http://localhost:11434"
  timeout: "30s"

git:
  allow_commit_override: true
  max_diff_lines: 500

ai:
  system_prompt: "You are an expert Git user..."
  temperature: 0.7

ui:
  theme: "dark"
  animations: true
```

### Environment Variables

Override any setting with environment variables prefixed with `GIT_PITCH_`:

```bash
export GIT_PITCH_DEBUG=true
export GIT_PITCH_OLLAMA_MODEL="codellama:latest"
```

## 🎯 Usage Examples

### Basic Commit Workflow

```bash
# Make your changes
echo "console.log('Hello, world!');" > hello.js

# Stage changes  
git add hello.js

# Generate AI commit message
pitch commit
# AI analyzes the diff and suggests: "Add hello world example in JavaScript"

# Commit with the suggested message
git commit -m "Add hello world example in JavaScript"
```

### Pull Request Generation

```bash
# Create a feature branch with changes
git checkout -b feature/user-authentication
# ... make changes ...
git push origin feature/user-authentication

# Generate PR with AI
pitch pr main
# Creates structured PR with title, summary, changes list, and testing instructions
```

### Interactive AI Consultation

```bash
# Ask about your code
pitch ask "How can I optimize this database query?"

# Multi-model comparison
pitch ask "Explain the difference between async/await and promises"
# Shows responses from multiple AI models side-by-side
```

## 🔧 Git Hook Integration

When you run `pitch apply`, it installs a Git hook that automatically generates commit messages:

```bash
# After applying hooks, normal git commit will use AI
git add .
git commit  # AI message appears in your editor automatically
```

Configure hook behavior in your config file:

```yaml
git:
  allow_commit_override: true  # AI message becomes default
  # OR
  allow_commit_override: false # AI message added as comment
```

## 🆘 Troubleshooting

### Common Issues

**"Ollama server not running"**
```bash
pitch start  # Start the Ollama server
pitch status # Check server status
```

**"Model not found"**
```bash
pitch model list    # See available models
pitch model select  # Choose a different model
```

**"No staged changes"**
```bash
git add .           # Stage your changes first
pitch commit        # Then generate message
```

### System Diagnostics

```bash
pitch info --health-check  # Comprehensive system check
pitch info --verbose       # Detailed debug information
```

## 🔄 Updates

Pitch CLI includes automatic update checking:

```bash
pitch update          # Update to latest version
pitch update --check-only  # Check for updates without installing
```

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

**Need help?** Run `pitch --help` or `pitch <command> --help` for detailed command information.