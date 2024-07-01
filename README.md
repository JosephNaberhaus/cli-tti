# CLI-TTI
A utility to measure how long a CLI program takes to be interactive.
![Timing the Node.js based commitizen client](meta/time.gif)

## Installation
### From source (recommended)
With v1.21 or greater of [Golang](https://go.dev/) you can easily install from source:

```shell
go install github.com/JosephNaberhaus/cli-tti@latest
```

### From binary
Download and extract the appropriate binary for your platform from the [releases page](https://github.com/JosephNaberhaus/cli-tti/releases).

## Usage
```
cli-tti [--wait-time] [-n] [--print-expected] <command to test>
```

#### Flags
`--wait-time`: The number of milliseconds to wait until the program is assumed to be interactive.

`-n`: The number of cycles to time the command before computing the average time-to-interactive.

`--print-expected`: Print the output that the utility has determined to be the interactive state of the command.