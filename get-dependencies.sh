#!/usr/bin/env bash
set -ex

## Linux
# sudo apt-get install libx11-dev
# sudo apt-get install libx11-xcb-dev
# sudo apt-get install libxkbcommon-x11-dev
# # sudo apt-get install xcb libxcb-xkb-dev x11-xkb-utils libx11-xcb-dev libxkbcommon-x11-dev
# sudo apt-get install nx-x11proto-xext-dev

# robotgo
sudo apt-get install -y gcc libc6-dev
sudo apt-get install -y libx11-dev xorg-dev libxtst-dev libpng++-dev
sudo apt-get install -y xcb libxcb-xkb-dev x11-xkb-utils libx11-xcb-dev libxkbcommon-x11-dev
sudo apt-get install -y libxkbcommon-dev
sudo apt-get install -y xsel xclip
# beep
sudo apt-get install -y libasound2-dev
