# WebAssembly Image Tool

## Introduction

This repository contains a simple Go-based WebAssembly (WASM) module for resizing images in the browser. The WASM module processes image resizing operations entirely on the client side, leveraging Go’s powerful image manipulation capabilities.

## Prerequisites

Before building the WASM module, ensure you have the following installed:

Go (version 1.18+ required)

## Demo

You can explore a live demo of the [Web Image Resizer](https://github.com/yuelone/Web-Image-Resizer) at the following link:

- [Web Image Resizer Demo](https://web-image-resizer.vercel.app/)

## Building the WASM Module

To compile the Go code into a WebAssembly module, run the following command:

```bash
# Compile Go code to WASM
GOOS=js GOARCH=wasm go build -o main.wasm main.go

# Copy the necessary WebAssembly runtime script
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./
```
