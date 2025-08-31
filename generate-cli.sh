#!/bin/bash

# Pitch CLI - Complete Command Generation Script
# Run this script to generate the entire CLI structure using cobra-cli

set -e  # Exit on any error

echo "🚀 Initializing Pitch CLI project..."

echo "📁 Project initialized. Generating commands..."

#━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
# INSTALLATION & ENVIRONMENT MANAGEMENT COMMANDS
#━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

echo "⚙️  Adding installation & environment commands..."

# Complete environment setup and Ollama installation
cobra-cli add install

# Complete removal with cleanup
cobra-cli add uninstall

# Self-update system
cobra-cli add update

#━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
# CORE AI-POWERED GIT WORKFLOW COMMANDS
#━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

echo "🤖 Adding AI-powered git workflow commands..."

# AI-powered commit message generation
cobra-cli add commit

# Pull request title and description generation
cobra-cli add pr

# Interactive AI consultation system
cobra-cli add ask

# Automatic documentation generation
cobra-cli add readme

#━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
# AI MODEL MANAGEMENT COMMANDS
#━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

echo "🧠 Adding AI model management commands..."

# Parent command for AI model management
cobra-cli add model

# Model subcommands
cobra-cli add list -p modelCmd
cobra-cli add select -p modelCmd
cobra-cli add create -p modelCmd
cobra-cli add delete -p modelCmd

# Using modelupdate to avoid conflict with main update command
cobra-cli add modelupdate -p modelCmd

#━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
# OLLAMA SERVER MANAGEMENT COMMANDS
#━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

echo "🖥️  Adding Ollama server management commands..."

# Start Ollama server with progress feedback
cobra-cli add start

# Gracefully stop Ollama server
cobra-cli add stop

# Restart server with health checks
cobra-cli add restart

# Real-time server status monitoring
cobra-cli add status

#━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
# GIT INTEGRATION & SYSTEM INFO COMMANDS
#━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

echo "📊 Adding Git integration & system info commands..."

# Git hooks installation and configuration
cobra-cli add apply

# System diagnostics and status dashboard
cobra-cli add info

echo ""
echo "✅ All commands generated successfully!"
echo ""
echo "📋 Generated command structure:"
echo "   pitch install       - Environment setup"
echo "   pitch uninstall     - Complete removal"
echo "   pitch update        - Self-update"
echo "   pitch commit        - AI commit messages"
echo "   pitch pr            - PR generation"
echo "   pitch ask           - AI consultation"
echo "   pitch readme        - Documentation generation"
echo "   pitch model         - Model management (with subcommands)"
echo "   pitch start         - Start Ollama server"
echo "   pitch stop          - Stop Ollama server"
echo "   pitch restart       - Restart Ollama server"
echo "   pitch status        - Server status"
echo "   pitch apply         - Install Git hooks"
echo "   pitch info          - System diagnostics"
echo ""
echo "🔧 Next steps:"
echo "   2. Edit cmd/*.go files to add flags and functionality"
echo "   3. Add internal/ directory structure"
echo "   4. go mod tidy"
echo "   5. go build"
echo ""
echo "🎉 CLI skeleton ready for development!"