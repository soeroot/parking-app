# Parking App CLI

A simple command-line tool written in Go to manage parking slots using a
min-heap strategy. It assigns the nearest available slot to cars entering the
parking lot and frees the slot when the car leaves.

---

## Features

- Read command from file only
- Allocate nearest available parking slot
- Release parking slot when a car leaves
- Show current parking status (sorted by slot number)
- Simple, dependency-light (no Cobra)
- Ready to install via `go install`

---

## Installation

Make sure Go 1.21+ is installed.

```bash
$ go install github.com/soeroot/parking-app@latest
$ parking-app <filePath>
