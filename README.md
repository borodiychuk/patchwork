# Patchwork quilt patterns generator

[![Build Status](https://travis-ci.org/borodiychuk/patchwork.svg?branch=master)](https://travis-ci.org/borodiychuk/patchwork) [![Go Report Card](https://goreportcard.com/badge/github.com/borodiychuk/patchwork)](https://goreportcard.com/report/github.com/borodiychuk/patchwork)

This is a tool that generates patchwork quilt pattern out of given square samples.

<p align="center" ><img src="https://user-images.githubusercontent.com/1705072/27992353-02a562f6-6493-11e7-8e57-b6b976c97f5a.png" alt="Patchwork output example with shuffle composer"> <img src="https://user-images.githubusercontent.com/1705072/27992352-02a2d824-6493-11e7-9b61-21f36a0b6cba.png" alt="Patchwork output example with crosses composer"></p>

## How to run it

Get sure you have Go installed, and then run:
```
go get github.com/borodiychuk/patchwork
```
After that you can just call it like that:
```
patchwork -out patchwork.png -sample-file s1.png -sample-file s2.png -sample-color 255,0,127 -sample-color 127,127,255 # ... and so on
```

## How is it built

The application ties together concepts of:
1. *Sample provider*. That is the data source for At this moment the sample can be imported from PNG file.
2. *Pattern composer*. It composes a pattern out of available samples. Right now it can combine random samples in a way that no similar samples stay aside horizontally or vertically.
3. *Result renderer*. It exports the result in particular format. Right now it can export to PNG.

## TODO

* Sample provider that generates sample of given color.
* Composer that generates tetris board filled with respective figures.
* Composer that generates patchwork representation out of provided image.
* Exporter tht exports the result as ASCII text with symbolic schema and pieces count.
