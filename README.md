# The Spins
Now your computer can get them too!

## What is it?
I've been teaching myself how linear algebra works. Just as a proof of concept, I wanted to
move a point about an axis on the screen. Frustraighted with all the incantations I had to
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

## Running
In your terminal run `the-spins`, to exit simply hit `CTRL+C`

## Lame Installation
`go get github.com/paul-nelson-baker/the-spins`

## True Installation
Clone the repository locally and run the following:

```
make generate
make install
```



