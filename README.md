# Run
![logo.png](logo.png)
> A universal script runner that executes code in 30+ programming languages with automatic runtime detection, installation assistance, and built-in benchmarking.

[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
<div class="logo-card">
    <div class="logo-display dark-bg">
        <svg width="250" height="150" viewBox="0 0 250 150">
            <rect x="10" y="10" width="230" height="130" rx="8" fill="#1a1a2e" stroke="#667eea" stroke-width="3"/>
            <!-- Terminal header -->
            <rect x="10" y="10" width="230" height="30" rx="8" fill="#667eea"/>
            <circle cx="25" cy="25" r="5" fill="#ff5f56"/>
            <circle cx="40" cy="25" r="5" fill="#ffbd2e"/>
            <circle cx="55" cy="25" r="5" fill="#27c93f"/>
            <!-- Terminal content -->
            <text x="25" y="70" font-family="monospace" font-size="16" fill="#00ff00">$ run script.py</text>
            <rect x="160" y="57" width="10" height="15" fill="#00ff00" opacity="0.8">
                <animate attributeName="opacity" values="0.8;0;0.8" dur="1s" repeatCount="indefinite"/>
            </rect>
        </svg>
    </div>
    <p class="logo-info">Perfect for developer tools and CLI apps</p>
</div>


## Why Run?

Stop memorizing different commands for different languages. Whether it's `python3 script.py`, `node app.js`, `go run main.go`, `gcc program.c -o program && ./program`, or `javac Program.java && java Program`, just use:

```bash
run script.ext
```

Run automatically:
- ✅ Detects the language from file extension
- ✅ Checks if the required runtime is installed
- ✅ Offers to install missing runtimes (with your permission)
- ✅ Compiles code when necessary
- ✅ Executes your program
- ✅ Cleans up temporary files
- ✅ Measures execution time
- ✅ Runs comprehensive benchmarks

## Features

- **30+ Languages Supported** - From Python to Assembly, C++ to Julia
- **Smart Runtime Detection** - Automatically checks if required tools are installed
- **Assisted Installation** - Prompts to install missing runtimes (requires user consent)
- **Compiled Language Support** - Handles compilation, execution, and cleanup automatically
- **Built-in Benchmarking** - Run performance tests with statistical analysis
- **Execution Timing** - Measure how long your code takes to run
- **Dry Run Mode** - Preview what will happen without executing
- **Cross-Platform** - Works on Linux, macOS, and Windows
- **Zero Configuration** - Works out of the box

## Quick Start

### Installation

**Option 1: Using Go (Recommended)**

```bash
go install github.com/Khaliiloo/run@latest
```

**Option 2: Download Pre-built Binary**

```bash
# Linux/macOS
curl -L https://github.com/Khaliiloo/run/releases/latest/download/install.sh | bash

# Or download manually from releases page
# https://github.com/Khaliiloo/run/releases
```

**Option 3: Build from Source**

```bash
git clone https://github.com/Khaliiloo/run.git
cd run
go build -o run
sudo mv run /usr/local/bin/
```

[//]: # ()
[//]: # (**Option 4: Using Package Managers** &#40;Coming Soon&#41;)

[//]: # ()
[//]: # (```bash)

[//]: # (# Homebrew &#40;macOS/Linux&#41;)

[//]: # (brew install run)

[//]: # ()
[//]: # (# APT &#40;Ubuntu/Debian&#41;)

[//]: # (sudo apt install run)

[//]: # (```)

### Basic Usage

```bash
# Run any script
run script.py
run app.js
run main.go
run program.cpp
run test.rb

# That's it! Run handles the rest.
```

## Usage & Examples

### Running Scripts

```bash
# Python
run hello.py

# JavaScript
run app.js

# C++ (automatically compiles and runs)
run program.cpp

# Java (compiles and runs)
run HelloWorld.java

# Rust (compiles and runs)
run main.rs
```

### Timing Execution

Measure how long your code takes to run:

```bash
run --time script.py
```

Output:
```
Running script.py...
Hello, World!

⏱  Execution time: 234ms
```

### Benchmarking

Run comprehensive performance benchmarks:

```bash
# Default: 10 iterations
run --bench algorithm.py

# Custom number of runs
run --bench 50 sorting.cpp

# Benchmark compiled languages
run --bench 100 fibonacci.rs
```

Output:
```
🔥 Running benchmark with 10 iterations...
==================================================
Compiling program.cpp...
✓ Compilation successful

Run 10/10... ✓ 45ms

==================================================
📊 Benchmark Results:
--------------------------------------------------
Runs:         10
Total time:   450ms
Average:      45ms
Median:       44ms
Min:          43ms
Max:          48ms
Std Dev:      2ms
==================================================
```

### Dry Run Mode

Preview what will happen without actually executing:

```bash
run --dry-run script.rb
```

Output:
```
🔍 Dry Run Mode - No execution will occur
=========================================
File: script.rb
Language: .rb
Runtime: ruby
✓ File exists
✓ Runtime 'ruby' is installed

Execution step:
  Command: ruby script.rb

✓ Dry run complete
```

### List Supported Languages

See all 30+ supported languages:

```bash
run --list
```

Output:
```
Supported Languages:
--------------------
Extension  Runtime         Type         Command
----------------------------------------------------------------------
.asm       nasm            Compiled     nasm -f elf64
.awk       awk             Interpreted  awk -f
.c         gcc             Compiled     gcc
.cpp       g++             Compiled     g++
.cs        dotnet          Compiled     dotnet build
.dart      dart            Interpreted  dart
.ex        elixir          Interpreted  elixir
...

Total: 30 languages supported
```

### Version Information

```bash
run --version
# Output: run version 1.0.0
```

### Help

```bash
run --help
```

Output:
```
run - Universal script runner

Usage:
  run [options] <source_file>

Options:
  --version, -v        Show version information
  --list, -l           List all supported languages
  --dry-run, -d        Show what would be executed without running
  --time, -t           Measure and display execution time
  --bench [n], -b [n]  run benchmark (default: 10 iterations)
  --help, -h           Show this help message

Examples:
  run script.py                 # Run Python script
  run --time app.js             # Run with execution time
  run --bench 20 program.cpp    # Benchmark with 20 runs
  run --dry-run test.go         # Preview without executing
  run --list                    # Show all supported languages
```

## 🗂️ Supported Languages

| Language | Extension | Type | Runtime | Auto-Install |
|----------|-----------|------|---------|--------------|
| Assembly | `.asm` | Compiled | NASM | ✅ |
| AWK | `.awk` | Interpreted | AWK | ✅ |
| C | `.c` | Compiled | GCC | ✅ |
| C++ | `.cpp` | Compiled | G++ | ✅ |
| C# | `.cs` | Compiled | .NET | ✅ |
| Dart | `.dart` | Interpreted | Dart | ✅ |
| Elixir | `.ex` | Interpreted | Elixir | ✅ |
| F# | `.fs` | Compiled | F# | ✅ |
| Go | `.go` | Interpreted | Go | ✅ |
| Groovy | `.groovy` | Interpreted | Groovy | ✅ |
| Haskell | `.hs` | Compiled | GHC | ✅ |
| Java | `.java` | Compiled | JDK | ✅ |
| JavaScript | `.js` | Interpreted | Node.js | ✅ |
| Julia | `.jl` | Interpreted | Julia | ✅ |
| Kotlin | `.kt` | Interpreted | Kotlin | ✅ |
| Lua | `.lua` | Interpreted | Lua | ✅ |
| Nim | `.nim` | Compiled | Nim | ✅ |
| OCaml | `.ml` | Compiled | OCaml | ✅ |
| Pascal | `.pas` | Compiled | FPC | ✅ |
| Perl | `.pl` | Interpreted | Perl | ✅ |
| PHP | `.php` | Interpreted | PHP | ✅ |
| Python | `.py` | Interpreted | Python 3 | ✅ |
| R | `.r` | Interpreted | Rscript | ✅ |
| Raku | `.raku` | Interpreted | Raku | ✅ |
| Ruby | `.rb` | Interpreted | Ruby | ✅ |
| Rust | `.rs` | Compiled | Rustc | ⚠️ Manual |
| Scheme | `.scm` | Interpreted | MIT Scheme | ✅ |
| Shell | `.sh` | Interpreted | Bash | ✅ |
| Swift | `.swift` | Interpreted | Swift | ✅ |
| Tcl | `.tcl` | Interpreted | Tclsh | ✅ |
| TypeScript | `.ts` | Interpreted | ts-node | ⚠️ Manual |
| VB.NET | `.vb` | Compiled | VBC | ✅ |
| Zig | `.zig` | Compiled | Zig | ✅ |

**Total: 30+ languages and counting!**

## 🔧 How It Works

1. **Language Detection**: Identifies the programming language from the file extension
2. **Runtime Verification**: Checks if the required compiler/interpreter is installed
3. **Installation Assistance**: If missing, prompts user to install (with permission)
4. **Compilation** (if needed): Automatically compiles code for compiled languages
5. **Execution**: Runs your program and displays output
6. **Cleanup**: Removes temporary compiled binaries (configurable)
7. **Reporting**: Shows timing and benchmark statistics if requested

### Compiled vs Interpreted Languages

**Interpreted Languages** (Python, JavaScript, Ruby, etc.):
```bash
run script.py
# Directly executes: python3 script.py
```

**Compiled Languages** (C++, Rust, Java, etc.):
```bash
run program.cpp
# 1. Compiles: g++ program.cpp -o program
# 2. Executes: ./program
# 3. Cleans up: removes ./program
```

## 🖥️ Platform-Specific Notes

### Linux (Ubuntu/Debian)
- Most runtimes can be auto-installed using `apt`
- Requires `sudo` privileges for installations
- Works out of the box for most languages

### macOS
- Requires [Homebrew](https://brew.sh/) for automatic installations
- Some tools (like Xcode Command Line Tools) may be pre-installed
- Install Homebrew first: `/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`

### Windows
- Many runtimes require manual installation
- Run provides download links for each runtime
- Consider using [Windows Subsystem for Linux (WSL)](https://docs.microsoft.com/en-us/windows/wsl/) for better compatibility
- Some features work best with Git Bash or PowerShell

## ⚠️ Security Considerations

**Important Security Notes:**

1. **Sudo Access**: Run may execute `sudo` commands when installing runtimes on Linux. Always review what's being installed.

2. **Code Execution**: This tool executes arbitrary code. Never run untrusted scripts without reviewing them first.

3. **Auto-Installation**: The auto-install feature runs system commands. While convenient, be aware of what's being installed on your system.


**Best Practices:**
- Always use `--dry-run` first with unfamiliar scripts
- Review code before executing
- Use in sandboxed/containerized environments for untrusted code
- Keep your system and runtimes updated

## 🤝 Contributing

Contributions are welcome! Here's how you can help:

### Adding a New Language

1. Fork the repository
2. Edit `main.go` and add to `languageConfigs`:

```go
".xyz": {
    CheckCmd: []string{"xyz", "--version"},
    InstallCmd: func() []string {
        switch runtime.GOOS {
        case "linux":
            return []string{"sudo", "apt", "install", "-y", "xyz"}
        case "darwin":
            return []string{"brew", "install", "xyz"}
        case "windows":
            return []string{"echo", "Install from https://xyz.org"}
        default:
            return []string{"echo", "Unsupported OS"}
        }
    },
    RunCmd: []string{"xyz"},
    IsCompiled: false, // or true if it needs compilation
},
```

3. Test your changes
4. Submit a pull request

### Other Ways to Contribute

- **Bug Reports**: Open an issue with reproduction steps
- **Feature Requests**: Suggest new features or improvements
- **Documentation**: Improve README, add examples, write tutorials
- **Testing**: Test on different platforms and report issues
- **Code Quality**: Refactoring, optimization, better error handling

### Development Setup

```bash
git clone https://github.com/Khaliiloo/run.git
cd run
go build -o run
./run --version
```

## 🗺️ Roadmap

### Version 1.1
- [ ] Configuration file support (`~/.runrc`)
- [ ] Custom language definitions
- [ ] Environment variable support
- [ ] Output to file option

### Version 1.2
- [ ] Package manager integration (apt, brew, chocolatey)
- [ ] Docker container support
- [ ] Remote execution capabilities
- [ ] Parallel execution mode

### Version 2.0
- [ ] Interactive REPL mode
- [ ] Plugin system for custom languages
- [ ] Web UI for benchmarks
- [ ] CI/CD integration tools
- [ ] Cloud execution support

### Community Requests
- [ ] Jupyter notebook support
- [ ] Multi-file project support
- [ ] Dependency management
- [ ] Code linting integration
- [ ] Test runner integration

## 📊 Performance

Run is designed to be lightweight and fast:

- **Startup time**: < 10ms
- **Memory footprint**: ~5MB
- **Language detection**: < 10ms
- **Minimal overhead**: Benchmark mode measures pure execution time

## 🐛 Troubleshooting

### Runtime Not Found After Installation

If a runtime still isn't found after installation:

```bash
# Reload your shell configuration
source ~/.bashrc  # or ~/.zshrc

# Check if runtime is in PATH
which python3

# Try manually specifying the path (future feature)
```

### Compilation Errors

For compiled languages, ensure you have the full toolchain:

```bash
# C/C++
sudo apt install build-essential  # Linux
xcode-select --install            # macOS

# Rust
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh

# Java
sudo apt install default-jdk      # Linux
brew install openjdk              # macOS
```

### Permission Denied

If you encounter permission errors:

```bash
# Linux/macOS: Make script executable
chmod +x script.sh

# For compiled binaries
chmod +x ./program
```

## 📚 Additional Resources

[//]: # (- **Documentation**: [Full docs]&#40;https://github.com/Khaliiloo/run/wiki&#41;)

[//]: # (- **Examples**: [Example scripts]&#40;https://github.com/Khaliiloo/run/tree/main/examples&#41;)
- **Community**: [Discussions](https://github.com/Khaliiloo/run/discussions)
- **Issues**: [Bug tracker](https://github.com/Khaliiloo/run/issues)

## 📄 License

MIT License

Copyright (c) 2025 Run

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

## 🙏 Acknowledgments

- Inspired by the need to quickly test code snippets across multiple languages

[//]: # (- Thanks to all contributors and users who provided feedback)
- Built with [Go](https://go.dev/) - a fantastic language for CLI tools and backend services

## 📧 Contact

- **Author**: Khalil Abdulgawad   (https://www.linkedin.com/in/kabdulgawad)
- **Email**: khalil.abdulgawad@gmail.com


## 🌟 Show Your Support

If you find this project useful, please consider:

- ⭐ **Starring** the repository
- 🐛 **Reporting** bugs
- 💡 **Suggesting** features
- 🔀 **Contributing** code
- 📢 **Sharing** with others

---

<div align="center">

**[⬆ back to top](#run)**

Built with Go • Engineered for developers • One command to run them all
</div>