# The Spins
Now your computer can get them too!

*_Allergy Warning: May contain easter-eggs_*

## Installation

```bash
$ go install github.com/paul-nelson-baker/the-spins@latest
```

## Usage

In your terminal run `the-spins`, to exit simply hit `CTRL+C` or `ESC`

## Easter Egg
* [youtube-dl](https://ytdl-org.github.io/youtube-dl/download.html) (optional)
* [ffmpeg](https://github.com/adaptlearning/adapt_authoring/wiki/Installing-FFmpeg) (optional)

```bash
$ make audio/spin-loop.mp3 install
```

## What is it?
I've been teaching myself how linear algebra works. Just as a proof of concept, I wanted to
move a point about an axis on the screen. Frustrated with all the incantations I had to
cast to get a simple proof of concept to work in OpenGL, I decided to utilize... the mouse!

The mouse can be considered a vector in space because any point X and Y can be interpreted
as a vector.

So what we do, take an arbitrary point `P_arb` and decide it's X and Y value are 75 and 75
respectively. We now take this point and rotate it by `Theta`, which we increment everytime
our main loop iterates. So now we have a point rotating about the axis, but this won't work
for our computer screen because the origin is in the top-left corner and we cannot move to
a negative position when `P_arb` has left `Quadrant I`. We now translate `P_arb` by another
vector `P_midScreen` which is a vector with X and Y values that are half the width and height
of the screen.

*This occurs inifinitely until the user hits `CTRL+C` or `ESC` to exit*
