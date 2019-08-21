# The Spins
Now your computer can get them too!

*_Allergy Warning: May contain easter-eggs_*

## Installation and Running

### Dependencies
You can install them yourself, or run the `install-dependencies.sh` script if you're on Mac or Linux
* [GoLang](https://golang.org/dl/)
* [youtube-dl](https://ytdl-org.github.io/youtube-dl/download.html) (optional)
* [ffmpeg](https://github.com/adaptlearning/adapt_authoring/wiki/Installing-FFmpeg) (optional)

### The Project, Itself
Like any other golang project you _can_ `go get github.com/paul-nelson-baker/the-spins`,
but you won't really get the full effect. What you should do is clone the repository
locally and run the following:

```
make generate
make install
```
> Note: For `make generate` to work you will need `youtube-dl` and `ffmpeg` installed locally

In your terminal run `the-spins`, to exit simply hit `CTRL+C`

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

*This occurs inifinitely until the user hits `CTRL+C` to exit*
