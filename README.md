# rnum

### Random Number Generator (CLI)

A simple command-line tool to generate random numbers within a specified range.

## Usage

```bash
# Generate a random number between 1 and N
rnum 10              # Returns random number between 1-10

# Generate a random number between two values
rnum 5 15            # Returns random number between 5-15
```

## Installation

### From Source

1. **Prerequisites**: Make sure you have Go installed (1.16 or later)

    ```bash
    go version
    ```

2. **Clone and build**:

    ```bash
    git clone https://github.com/terryluciano/rnum.git
    cd rnum
    go build -o rnum main.go
    ```

3. **Install system-wide**:

    ```bash
    sudo mv rnum /usr/local/bin/
    ```

    Or install for current user only:

    ```bash
    mkdir -p ~/.local/bin
    mv rnum ~/.local/bin/
    # Add to PATH if not already there
    echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
    source ~/.bashrc
    ```

### Using go install

```bash
go install github.com/terryluciano/rnum@latest
```

## Examples

```bash
# Roll a dice
rnum 1 6

# Generate a random percentage
rnum 0 100

# Pick a random port number
rnum 8000 9000

# Single argument defaults to range 1-N
rnum 100    # Same as: rnum 1 100
```

## License

MIT License. See [LICENSE](LICENSE) for details.
