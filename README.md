# `mol` The Molecule Renderer

![Examples](http://i.imgur.com/o5bzQ42.png)

## Installation

You need `cairo` and `pkg-config` installed first.

    go get github.com/fogleman/mol

## Usage

The program parses MDL Molfiles or SDF files and generates PNGs. There are no other command line options as of yet.

    mol input.sdf
    mol examples/*

The program automatically positions the camera to maximize visibility of all atoms in the molecule.
