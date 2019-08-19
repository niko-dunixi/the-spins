#!/usr/bin/env bash
set -e

if ! type youtube-dl > /dev/null; then
  echo "You need: youtube-dl"
  exit 127
elif ! type ffmpeg > /dev/null; then
  echo "You need: ffmpeg"
  exit 127
fi

bin_directory="./bin"
asset_directory="./data/assets"
song_url="https://www.youtube.com/watch?v=z8cgNLGnnK4"
song_file="${bin_directory}/nsp-you-spin-me-cover.mp3"
loop_file="${asset_directory}/spin-loop.mp3"

mkdir -p "${bin_directory}"
mkdir -p "${asset_directory}"

if [ ! -f "${song_file}" ]; then
  youtube-dl "${song_url}" --extract-audio --audio-format mp3 --exec "mv {} ${song_file}"
fi
if [ ! -f "${loop_file}" ]; then
  ffmpeg -i "${song_file}" -ss 00:01:13.30 -to 00:01:30.38 -c copy "${loop_file}"
fi
echo "I haz muzic: ${loop_file}"
