#!/usr/bin/env bash
set -e

REPO="guilhermehbueno/git-pitch"
LATEST=$(curl -s https://api.github.com/repos/$REPO/releases/latest \
  | grep "browser_download_url.*$(uname -s)_$(uname -m)" \
  | cut -d '"' -f 4)

curl -L "$LATEST" -o /tmp/git-pitch.tar.gz
tar -xzf /tmp/git-pitch.tar.gz -C /tmp
sudo mv /tmp/git-pitch /usr/local/bin