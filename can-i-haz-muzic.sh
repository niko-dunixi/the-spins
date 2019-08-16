#!/usr/bin/env bash
set -e

song_url="https://www.youtube.com/watch?v=z8cgNLGnnK4"
song_file="nsp-you-spin-me-cover.mp3"
loop_file="spin-loop.mp3"

if [ ! -f "./${song_file}" ]; then
  youtube-dl "${song_url}" -x --audio-format mp3 -o "${song_file}"
fi
if [ ! -f "./${loop_file}" ]; then
  ffmpeg -i "${song_file}" -ss 00:01:16.00 -to 00:01:23.00 -c copy "./${loop_file}"
fi
