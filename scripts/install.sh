#!/usr/bin/env bash
set -euo pipefail

REPO="guilhermehbueno/git-pitch"
OS="$(uname -s)"
ARCH="$(uname -m)"

# normalize arch names to match your GoReleaser archives
case "$ARCH" in
  x86_64) ARCH="x86_64" ;;        # mac/linux amd64 mapped to x86_64 per your template
  amd64)  ARCH="x86_64" ;;
  aarch64) ARCH="arm64" ;;
esac

ASSET_URL=$(
  curl -s "https://api.github.com/repos/${REPO}/releases/latest" |
    grep -Eo '"browser_download_url": *"[^"]+"' |
    cut -d'"' -f4 |
    grep "${OS}" | grep "${ARCH}" || true
)

if [[ -z "${ASSET_URL}" ]]; then
  echo "Could not find a release asset for ${OS}_${ARCH}." >&2
  exit 1
fi

TMPDIR="$(mktemp -d)"
ARCHIVE="${TMPDIR}/artifact"

curl -L "${ASSET_URL}" -o "${ARCHIVE}"

# extract
case "${ASSET_URL}" in
  *.tar.gz) tar -xzf "${ARCHIVE}" -C "${TMPDIR}" ;;
  *.zip)    unzip -q "${ARCHIVE}" -d "${TMPDIR}" ;;
esac

# move binary (assuming it’s named 'git-pitch' at archive root)
BIN="${TMPDIR}/git-pitch"
chmod +x "${BIN}"

# destination
DEST="/usr/local/bin/git-pitch"
if [[ ! -w "$(dirname "${DEST}")" ]]; then
  sudo mv "${BIN}" "${DEST}"
else
  mv "${BIN}" "${DEST}"
fi

echo "Installed to ${DEST}"
"${DEST}" --version || true