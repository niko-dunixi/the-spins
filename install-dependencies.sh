#!/usr/bin/env bash
set -e

# Download and install Go
if [[ "$OSTYPE" == "darwin"* ]]; then
  if ! type brew > /dev/null; then
    /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
  fi
  brew update
  brew upgrade
  brew install go youtube-dl ffmpeg
else
  # Assuming linux if not OSX
  # Install Go
  if ! type go > /dev/null; then
    wget https://dl.google.com/go/go1.12.9.linux-amd64.tar.gz
    tar -xvf go1.12.9.linux-amd64.tar.gz
    sudo mkdir -p /usr/local/go
    sudo mv go /usr/local/go/go1.12.9
    echo "Adding GO variables to your .bashrc file"
    # Add to your .bashrc file

cat <<END_OF_BASHRC >> "${HOME}/.bashrc"
export GOPATH="${HOME}/go"
export GO111MODULE="on"
# We only want to use our personally managed version of go if
# there isn't one installed via homebrew or other outside process.
# More than one on PATH causes weird compiliation errors
if [ -f "/usr/local/bin/go" ]; then
  export PATH="${PATH}:${GOPATH}/bin"
else
  export GOROOT="/usr/local/go/go1.12.9"
  export PATH="${PATH}:${GOROOT}/bin:${GOPATH}/bin"
fi
END_OF_BASHRC
  fi
  # Install youtube-dl
  if ! type youtube-dl > /dev/null; then
    sudo pip install --upgrade youtube_dl
  fi
  # Install ffmpeg
  if ! type ffmpeg > /dev/null; then
    sudo apt install ffmpeg
  fi
fi

echo "Done!"
