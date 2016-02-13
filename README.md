# `mol` The Molecule Renderer

![Examples](http://i.imgur.com/o5bzQ42.png)

## Installation

You need `cairo` and `pkg-config` installed first.

    go get github.com/fogleman/mol

## Usage

The program parses MDL Molfiles or SDF files and generates PNGs.

    mol input.sdf
    mol examples/*

The program automatically positions the camera to maximize visibility of all atoms in the molecule.

There are no other command line options as of yet, though one can imagine several possibilities. Output resolution, camera position, animations, etc. Quality pull requests are welcome.
