GoRender: A software rendering experiment in Go
===============================================

GoRender is a 3D software renderer for Go. It is an experiment to put together
a neat API and glue together what I have learnt about graphics programming over
the years.

I'd like it to use as much textbook notation as possible, which is possible in
Go due to UTF8 magic. I'd also like this to be fully self-sufficient in that it
will not use any external libraries - though should hopefully make it possible
to compose them in future.

To Do
-----

### Pre-requisites

 + [x] Decide on SDL2 binding for Go
 + [x] Create drawable backbuffer and test example

### Primitives

 + [ ] Implemenent file reader (.obj)
 + [ ] Build a few primitives (cuboid, sphere)

### Transformation

 + [ ] Build format for vector representations
 + [ ] Build vector and matrix libraries
 + [ ] Write transformation and rotation tests

### Camera

### Lighting

### Clipping

### Texturing

### Interpolation and rasterization

### Shading and rendering

### Scene graph

### Benchmarking and examples
