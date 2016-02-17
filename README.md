# `mol` The Molecule Renderer

![Examples](http://i.imgur.com/o5bzQ42.png)

## Installation

    go get github.com/fogleman/mol

## Usage

The program parses MDL Molfiles or SDF files and generates PNGs.

    mol input.sdf
    mol examples/*

There are no other command line options as of yet, though one can imagine several possibilities. Output resolution, output format (SVG), camera position, animations, etc. Quality pull requests are welcome. Also, I know very little about chemistry, so if something doesn't look right, let me know!

## Input Files

You can search for molecules here:

https://www.ncbi.nlm.nih.gov/pccompound/?term=benzene

Select a result and click `Download > 3D Conformer > SDF`.

There are also several examples in the `examples` folder.

## How it Works

`mol` uses another library I created called `ln`. `ln` is a 3D vector renderer that works somewhat like a ray tracer. `mol` simply constructs spheres and cylinders in the right places and uses `ln` to render.

https://github.com/fogleman/ln

## Camera Placement

The program automatically positions the camera to maximize visibility of all atoms in the molecule.

![Example](http://i.imgur.com/hrptdsp.png)

The two molecules shown above are the same (testosterone). One is shown from the most optimal point of view while the other is shown from the least optimal point of view. See [camera.go](https://github.com/fogleman/mol/blob/master/mol/camera.go) to see how this works. It's pretty straightforward.
