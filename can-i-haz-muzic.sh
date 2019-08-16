#!/usr/bin/env bash
set -e

if ! type youtube-dl > /dev/null; then
  echo "You need: youtube-dl"
  exit 127
elif ! type ffmpeg > /dev/null; then
  echo "You need: ffmpeg"
  exit 127
fi

song_url="https://www.youtube.com/watch?v=z8cgNLGnnK4"
song_file="nsp-you-spin-me-cover.mp3"
loop_file="spin-loop.mp3"

if [ ! -f "./${song_file}" ]; then
  youtube-dl "${song_url}" --extract-audio --audio-format mp3 --exec "mv {} ${song_file}"
fi
if [ ! -f "./${loop_file}" ]; then
  ffmpeg -i "${song_file}" -ss 00:01:15.50 -to 00:01:23.00 -c copy "./${loop_file}"
fi
echo "I haz muzic: ${loop_file}"
