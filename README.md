# Patchwork quilt patterns generator

[![Build Status](https://travis-ci.org/borodiychuk/patchwork.svg?branch=master)](https://travis-ci.org/borodiychuk/patchwork) [![Go Report Card](https://goreportcard.com/badge/github.com/borodiychuk/patchwork)](https://goreportcard.com/report/github.com/borodiychuk/patchwork)

This is a tool that generates patchwork quilt pattern out of given square samples. 

<p align="center" ><img src="https://user-images.githubusercontent.com/1705072/27860688-e5ccdb04-617d-11e7-9778-bec33440710d.png" alt="Patchwork"></p>

## How is it built

The application ties together concepts of:
1. *Sample provider*. That is the data source for At this moment the sample can be imported from PNG file.
2. *Pattern composer*. It composes a pattern out of available samples. Right now it can combine random samples in a way that no similar samples stay aside horizontally or vertically.
3. *Result renderer*. It exports the result in particular format. Right now it can export to PNG.

## TODO

* Sample provider that generates sample of given color.
* Composer that generates tetris board filled with respective figures.
* Composer that generates patchwork representation out of provided image.
