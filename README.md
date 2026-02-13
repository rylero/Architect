# Architect

A toolchain written in Go designed to build, debug, and test digital logic circuts. It works of the idea of increasing levels of abstraction, starting from basic logic gates, and then creating builder functions that build layers of abstraction on top. Using this pattern you get both the ability to sketch the overal architecture while still being able to debug the individual logic gates themselves.

## Features:
 - Basic Logic Gates
 - Simple Tests
 - Built in Debugging with Scoped Probes

## Todo:
 - [ ] Add support for direct outputs in builder functions
 - [ ] Continuous Simulation Runner
 - [ ] Breakpoints
 - [ ] Output Options (binary, base10, lcd, etc.
 - [ ] Standard Library (similar to logisim evolution standard library)
 - [ ] Better Testing Support (Unit Tests, Integration Tests, High Level System Tests)
