package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"
)

/*
	Developer: Khalil Abdulgawad

	This Go program is a command-line utility that checks for the presence of
		various programming language runtimes, installs them if they are missing,
		and then compiles or runs a given source code file based on its extension.
	Designed to support a wide range of programming languages, it uses OS-specific
		commands for installation and execution.
*/

const version = "1.0.0"

// LanguageConfig holds configuration for each supported language
type LanguageConfig struct {
	CheckCmd    []string
	InstallCmd  func() []string // Function to return OS-specific install commands
	RunCmd      []string
	CompileCmd  []string // For compiled languages
	IsCompiled  bool
	ClassNameFn func(string) string // For Java, to get class name from file name
}

var languageConfigs = map[string]LanguageConfig{
	".py": {
		CheckCmd: []string{"python3", "--version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "python3"}
			case "darwin":
				return []string{"brew", "install", "python"}
			case "windows":
				return []string{"echo", "Please install Python from https://www.python.org/downloads/"}
			default:
				return []string{"echo", "Unsupported OS for automatic Python installation."}
			}
		},
		RunCmd: []string{"python3"},
	},
	".go": {
		CheckCmd: []string{"go", "version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "golang-go"}
			case "darwin":
				return []string{"brew", "install", "go"}
			case "windows":
				return []string{"echo", "Please install Go from https://go.dev/dl"}
			default:
				return []string{"echo", "Unsupported OS for automatic Go installation."}
			}
		},
		RunCmd: []string{"go", "run"},
	},
	".js": {
		CheckCmd: []string{"node", "--version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "nodejs"}
			case "darwin":
				return []string{"brew", "install", "node"}
			case "windows":
				return []string{"echo", "Please install Node.js from https://nodejs.org/en/download/"}
			default:
				return []string{"echo", "Unsupported OS for automatic Node.js installation."}
			}
		},
		RunCmd: []string{"node"},
	},
	".rb": {
		CheckCmd: []string{"ruby", "--version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "ruby"}
			case "darwin":
				return []string{"brew", "install", "ruby"}
			case "windows":
				return []string{"echo", "Please install Ruby from https://rubyinstaller.org/"}
			default:
				return []string{"echo", "Unsupported OS for automatic Ruby installation."}
			}
		},
		RunCmd: []string{"ruby"},
	},
	".java": {
		CheckCmd: []string{"java", "--version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "default-jdk"}
			case "darwin":
				return []string{"brew", "install", "openjdk"}
			case "windows":
				return []string{"echo", "Please install Java JDK from https://www.oracle.com/java/technologies/downloads/"}
			default:
				return []string{"echo", "Unsupported OS for automatic Java installation."}
			}
		},
		CompileCmd: []string{"javac"},
		RunCmd:     []string{"java"},
		IsCompiled: true,
		ClassNameFn: func(filename string) string {
			return strings.TrimSuffix(filename, filepath.Ext(filename))
		},
	},
	".cpp": {
		CheckCmd: []string{"g++", "--version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "build-essential"}
			case "darwin":
				return []string{"xcode-select", "--install"}
			case "windows":
				return []string{"echo", "Please install MinGW-w64 or Visual Studio with C++ tools."}
			default:
				return []string{"echo", "Unsupported OS for automatic C++ installation."}
			}
		},
		CompileCmd: []string{"g++"},
		IsCompiled: true,
	},
	".c": {
		CheckCmd: []string{"gcc", "--version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "build-essential"}
			case "darwin":
				return []string{"xcode-select", "--install"}
			case "windows":
				return []string{"echo", "Please install MinGW-w64 or Visual Studio with C tools."}
			default:
				return []string{"echo", "Unsupported OS for automatic C installation."}
			}
		},
		CompileCmd: []string{"gcc"},
		IsCompiled: true,
	},
	".rs": {
		CheckCmd: []string{"rustc", "--version"},
		InstallCmd: func() []string {
			return []string{"echo", "Please install Rust from https://rustup.rs/ by running: curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh"}
		},
		CompileCmd: []string{"rustc"},
		RunCmd:     []string{filepath.Base(strings.TrimSuffix(os.Args[1], filepath.Ext(os.Args[1])))},
		IsCompiled: true,
	},
	".cs": {
		CheckCmd: []string{"dotnet", "--version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "dotnet-sdk-8.0"}
			case "darwin":
				return []string{"brew", "install", "dotnet"}
			case "windows":
				return []string{"echo", "Please install .NET SDK from https://dotnet.microsoft.com/download"}
			default:
				return []string{"echo", "Unsupported OS for automatic C# installation."}
			}
		},
		CompileCmd: []string{"dotnet", "build"},
		RunCmd:     []string{"dotnet", "run"},
		IsCompiled: true,
	},
	".sh": {
		CheckCmd: []string{"bash", "--version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "bash"}
			case "darwin":
				return []string{"brew", "install", "bash"}
			case "windows":
				return []string{"echo", "Please install Git Bash from https://gitforwindows.org/"}
			default:
				return []string{"echo", "Unsupported OS for automatic Bash installation."}
			}
		},
		RunCmd: []string{"bash"},
	},
	".pl": {
		CheckCmd: []string{"perl", "--version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "perl"}
			case "darwin":
				return []string{"brew", "install", "perl"}
			case "windows":
				return []string{"echo", "Please install Strawberry Perl from http://strawberryperl.com/"}
			default:
				return []string{"echo", "Unsupported OS for automatic Perl installation."}
			}
		},
		RunCmd: []string{"perl"},
	},
	".php": {
		CheckCmd: []string{"php", "--version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "php"}
			case "darwin":
				return []string{"brew", "install", "php"}
			case "windows":
				return []string{"echo", "Please install PHP from https://windows.php.net/download/"}
			default:
				return []string{"echo", "Unsupported OS for automatic PHP installation."}
			}
		},
		RunCmd: []string{"php"},
	},
	".ts": {
		CheckCmd: []string{"ts-node", "--version"},
		InstallCmd: func() []string {
			return []string{"echo", "Please install Node.js and then run: npm install -g ts-node typescript"}
		},
		RunCmd: []string{"ts-node"},
	},
	".lua": {
		CheckCmd: []string{"lua", "--version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "lua5.3"}
			case "darwin":
				return []string{"brew", "install", "lua"}
			case "windows":
				return []string{"echo", "Please install Lua from https://www.lua.org/download.html"}
			default:
				return []string{"echo", "Unsupported OS for automatic Lua installation."}
			}
		},
		RunCmd: []string{"lua"},
	},
	".r": {
		CheckCmd: []string{"Rscript", "--version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "r-base"}
			case "darwin":
				return []string{"brew", "install", "r"}
			case "windows":
				return []string{"echo", "Please install R from https://cran.r-project.org/bin/windows/base/"}
			default:
				return []string{"echo", "Unsupported OS for automatic R installation."}
			}
		},
		RunCmd: []string{"Rscript"},
	},
	".hs": {
		CheckCmd: []string{"ghc", "--version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "ghc"}
			case "darwin":
				return []string{"brew", "install", "ghc"}
			case "windows":
				return []string{"echo", "Please install GHC from https://www.haskell.org/ghc/download_ghc_9_10_3.html"}
			default:
				return []string{"echo", "Unsupported OS for automatic Haskell installation."}
			}
		},
		CompileCmd: []string{"ghc"},
		IsCompiled: true,
	},
	".swift": {
		CheckCmd: []string{"swift", "--version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"echo", "Please install Swift from https://swift.org/download/#releases"}
			case "darwin":
				return []string{"brew", "install", "swift"}
			case "windows":
				return []string{"echo", "Please install Swift from https://swift.org/download/#releases"}
			default:
				return []string{"echo", "Unsupported OS for automatic Swift installation."}
			}
		},
		RunCmd: []string{"swift"},
	},
	".groovy": {
		CheckCmd: []string{"groovy", "--version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "groovy"}
			case "darwin":
				return []string{"brew", "install", "groovy"}
			case "windows":
				return []string{"echo", "Please install Groovy from https://groovy-lang.org/download.html"}
			default:
				return []string{"echo", "Unsupported OS for automatic Groovy installation."}
			}
		},
		RunCmd: []string{"groovy"},
	},
	".kt": {
		CheckCmd: []string{"kotlinc", "-version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "kotlin"}
			case "darwin":
				return []string{"brew", "install", "kotlin"}
			case "windows":
				return []string{"echo", "Please install Kotlin from https://kotlinlang.org/docs/command-line.html"}
			default:
				return []string{"echo", "Unsupported OS for automatic Kotlin installation."}
			}
		},
		RunCmd: []string{"kotlinc", "-script"},
	},
	".ex": {
		CheckCmd: []string{"elixir", "--version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "elixir"}
			case "darwin":
				return []string{"brew", "install", "elixir"}
			case "windows":
				return []string{"echo", "Please install Elixir from https://elixir-lang.org/install.html"}
			default:
				return []string{"echo", "Unsupported OS for automatic Elixir installation."}
			}
		},
		RunCmd: []string{"elixir"},
	},
	".ml": {
		CheckCmd: []string{"ocamlc", "-version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "ocaml"}
			case "darwin":
				return []string{"brew", "install", "ocaml"}
			case "windows":
				return []string{"echo", "Please install OCaml from https://ocaml.org/docs/install.html"}
			default:
				return []string{"echo", "Unsupported OS for automatic OCaml installation."}
			}
		},
		CompileCmd: []string{"ocamlc"},
		IsCompiled: true,
	},
	".nim": {
		CheckCmd: []string{"nim", "--version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "nim"}
			case "darwin":
				return []string{"brew", "install", "nim"}
			case "windows":
				return []string{"echo", "Please install Nim from https://nim-lang.org/install.html"}
			default:
				return []string{"echo", "Unsupported OS for automatic Nim installation."}
			}
		},
		CompileCmd: []string{"nim", "c"},
		IsCompiled: true,
	},
	".dart": {
		CheckCmd: []string{"dart", "--version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "dart"}
			case "darwin":
				return []string{"brew", "install", "dart"}
			case "windows":
				return []string{"echo", "Please install Dart from https://dart.dev/get-dart"}
			default:
				return []string{"echo", "Unsupported OS for automatic Dart installation."}
			}
		},
		RunCmd: []string{"dart"},
	},
	".raku": {
		CheckCmd: []string{"raku", "--version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "raku"}
			case "darwin":
				return []string{"brew", "install", "raku"}
			case "windows":
				return []string{"echo", "Please install Raku from https://raku.org/downloads/"}
			default:
				return []string{"echo", "Unsupported OS for automatic Raku installation."}
			}
		},
		RunCmd: []string{"raku"},
	},
	".tcl": {
		CheckCmd: []string{"tclsh"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "tcl"}
			case "darwin":
				return []string{"brew", "install", "tcl-tk"}
			case "windows":
				return []string{"echo", "Please install Tcl from https://www.activestate.com/products/tcl/"}
			default:
				return []string{"echo", "Unsupported OS for automatic Tcl installation."}
			}
		},
		RunCmd: []string{"tclsh"},
	},
	".vb": {
		CheckCmd: []string{"vbc", "--version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "mono-complete"}
			case "darwin":
				return []string{"brew", "install", "mono"}
			case "windows":
				return []string{"echo", "Please install Visual Studio with VB.NET support."}
			default:
				return []string{"echo", "Unsupported OS for automatic VB.NET installation."}
			}
		},
		CompileCmd: []string{"vbc"},
		IsCompiled: true,
	},
	".fs": {
		CheckCmd: []string{"fsharpc", "--version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "fsharp"}
			case "darwin":
				return []string{"brew", "install", "fsharp"}
			case "windows":
				return []string{"echo", "Please install Visual Studio with F# support."}
			default:
				return []string{"echo", "Unsupported OS for automatic F# installation."}
			}
		},
		CompileCmd: []string{"fsharpc"},
		IsCompiled: true,
	},
	".pas": {
		CheckCmd: []string{"fpc", "--version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "fpc"}
			case "darwin":
				return []string{"brew", "install", "fpc"}
			case "windows":
				return []string{"echo", "Please install Free Pascal from https://www.freepascal.org/download.var"}
			default:
				return []string{"echo", "Unsupported OS for automatic Pascal installation."}
			}
		},
		CompileCmd: []string{"fpc"},
		IsCompiled: true,
	},
	".jl": {
		CheckCmd: []string{"julia", "--version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "julia"}
			case "darwin":
				return []string{"brew", "install", "julia"}
			case "windows":
				return []string{"echo", "Please install Julia from https://julialang.org/downloads/"}
			default:
				return []string{"echo", "Unsupported OS for automatic Julia installation."}
			}
		},
		RunCmd: []string{"julia"},
	},
	".scm": {
		CheckCmd: []string{"scheme", "--version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "mit-scheme"}
			case "darwin":
				return []string{"brew", "install", "mit-scheme"}
			case "windows":
				return []string{"echo", "Please install MIT/GNU Scheme from https://www.gnu.org/software/mit-scheme/"}
			default:
				return []string{"echo", "Unsupported OS for automatic Scheme installation."}
			}
		},
		RunCmd: []string{"scheme"},
	},
	".awk": {
		CheckCmd: []string{"awk", "--version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "gawk"}
			case "darwin":
				return []string{"brew", "install", "gawk"}
			case "windows":
				return []string{"echo", "Please install Gawk from http://gnuwin32.sourceforge.net/packages/gawk.htm"}
			default:
				return []string{"echo", "Unsupported OS for automatic Awk installation."}
			}
		},
		RunCmd: []string{"awk", "-f"},
	},
	".asm": {
		CheckCmd: []string{"nasm", "--version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "nasm"}
			case "darwin":
				return []string{"brew", "install", "nasm"}
			case "windows":
				return []string{"echo", "Please install NASM from https://www.nasm.us/pub/nasm/releasebuilds/"}
			default:
				return []string{"echo", "Unsupported OS for automatic NASM installation."}
			}
		},
		CompileCmd: []string{"nasm", "-f", "elf64"},
		IsCompiled: true,
	},
	".zig": {
		CheckCmd: []string{"zig", "version"},
		InstallCmd: func() []string {
			switch runtime.GOOS {
			case "linux":
				return []string{"sudo", "apt", "install", "-y", "zig"}
			case "darwin":
				return []string{"brew", "install", "zig"}
			case "windows":
				return []string{"echo", "Please install Zig from https://ziglang.org/download/"}
			default:
				return []string{"echo", "Unsupported OS for automatic Zig installation."}
			}
		},
		CompileCmd: []string{"zig", "build-exe"},
		IsCompiled: true,
	},
}

func main() {

	// Handle flags
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--version", "-v":
			fmt.Printf("run version %s\n", version)
			os.Exit(0)
		case "--list", "-l":
			listLanguages()
			os.Exit(0)
		case "--help", "-h":
			printHelp()
			os.Exit(0)
		}
	} else {
		printHelp()
		os.Exit(1)
	}

	// Parse flags and file
	var dryRun, timeExec, bench bool
	var sourceFile string
	benchRuns := 10 // Default number of benchmark runs

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		switch {
		case arg == "--dry-run" || arg == "-d":
			dryRun = true
		case arg == "--time" || arg == "-t":
			timeExec = true
		case arg == "--bench" || arg == "-b":
			bench = true
			// Check if next arg is a number for bench runs
			if i+1 < len(os.Args) && isNumeric(os.Args[i+1]) {
				fmt.Sscanf(os.Args[i+1], "%d", &benchRuns)
				i++
			}
		case !strings.HasPrefix(arg, "--"):
			sourceFile = arg
		}
	}

	if sourceFile == "" {
		fmt.Println("Usage: run [options] <source_file>")
		fmt.Println("\nOptions:")
		fmt.Println("  --version, -v        Show version")
		fmt.Println("  --list, -l           List all supported languages")
		fmt.Println("  --dry-run, -d            Show what would be executed without running")
		fmt.Println("  --time, -t               Measure and display execution time")
		fmt.Println("  --bench [n], -b [n]          Run benchmark (default: 10 runs)")
		fmt.Println("  --help, -h           Show this help message")
		os.Exit(1)
	}

	// Validate conflicting flags
	if bench && timeExec {
		fmt.Println("Warning: --bench already includes timing. Ignoring --time flag.")
		timeExec = false
	}
	if dryRun && (timeExec || bench) {
		fmt.Println("Warning: --dry-run cannot be used with --time or --bench. Ignoring timing flags.")
		timeExec = false
		bench = false
	}

	ext := filepath.Ext(sourceFile)

	config, ok := languageConfigs[ext]

	if !ok {
		fmt.Printf("Unsupported file type: %s\n", ext)
		fmt.Println("Run 'run --list' to see supported languages.")
		os.Exit(1)
	}

	installCmd := config.InstallCmd()

	if !checkRuntime(config.CheckCmd) {
		if dryRun {
			fmt.Printf("✗ Runtime '%s' not found (would prompt for installation)\n", config.CheckCmd[0])
			os.Exit(1)
		}
		fmt.Printf("%s not found. Do you want to install it? (y/n): ", config.CheckCmd[0])
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		if strings.ToLower(strings.TrimSpace(input)) == "y" {
			if installCmd[0] == "echo" {
				fmt.Println(installCmd[1])
				fmt.Println("Please install the runtime manually and re-run the command.")
				os.Exit(1)
			}
			if !installRuntime(installCmd) {
				fmt.Println("Installation failed. Exiting.")
				os.Exit(1)
			}
			// Re-check after installation
			if !checkRuntime(config.CheckCmd) {
				fmt.Println("Runtime still not found after installation. Exiting.")
				os.Exit(1)
			}
		} else {
			fmt.Println("Installation declined. Exiting.")
			os.Exit(1)
		}
	}

	if dryRun {
		performDryRun(sourceFile, config, ext)
		os.Exit(0)
	}

	if bench {
		performBenchmark(sourceFile, config, ext, benchRuns)
		os.Exit(0)
	}

	// Normal execution with optional timing
	var start time.Time
	if timeExec {
		start = time.Now()
	}

	executeFile(sourceFile, config, ext)

	if timeExec {
		elapsed := time.Since(start)
		fmt.Printf("\n⏱  Execution time: %v\n", elapsed)
	}

	fmt.Println()
}

func listLanguages() {
	fmt.Println("Supported Languages:")
	fmt.Println("--------------------")

	// Sort extensions for consistent output
	extensions := make([]string, 0, len(languageConfigs))
	for ext := range languageConfigs {
		extensions = append(extensions, ext)
	}
	sort.Strings(extensions)

	fmt.Printf("%-10s %-15s %-12s %s\n", "Extension", "Runtime", "Type", "Command")
	fmt.Println(strings.Repeat("-", 70))

	for _, ext := range extensions {
		config := languageConfigs[ext]
		runtime := config.CheckCmd[0]
		langType := "Interpreted"
		if config.IsCompiled {
			langType = "Compiled"
		}

		cmdStr := strings.Join(config.RunCmd, " ")
		if config.IsCompiled && len(config.CompileCmd) > 0 {
			cmdStr = strings.Join(config.CompileCmd, " ")
		}

		fmt.Printf("%-10s %-15s %-12s %s\n", ext, runtime, langType, cmdStr)
	}

	fmt.Printf("\nTotal: %d languages supported\n", len(languageConfigs))
}

func performDryRun(sourceFile string, config LanguageConfig, ext string) {
	fmt.Println(" Dry Run Mode - No execution will occur")
	fmt.Println("=========================================")
	fmt.Printf("File: %s\n", sourceFile)
	fmt.Printf("Language: %s\n", ext)
	fmt.Printf("Runtime: %s\n", config.CheckCmd[0])

	// Check if file exists
	if _, err := os.Stat(sourceFile); os.IsNotExist(err) {
		fmt.Printf("✗ File not found: %s\n", sourceFile)
		return
	} else {
		fmt.Printf("✓ File exists\n")
	}

	// Check runtime
	if checkRuntime(config.CheckCmd) {
		fmt.Printf("✓ Runtime '%s' is installed\n", config.CheckCmd[0])
	} else {
		fmt.Printf("✗ Runtime '%s' not found\n", config.CheckCmd[0])
		return
	}

	if config.IsCompiled {
		fmt.Println("\nCompilation step:")
		executableName := strings.TrimSuffix(sourceFile, filepath.Ext(sourceFile))
		compileArgs := []string{}

		if ext == ".rs" {
			compileArgs = append(config.CompileCmd[1:], sourceFile)
		} else if ext == ".cs" {
			fmt.Printf("  Would create .NET project and compile\n")
		} else {
			compileArgs = append(config.CompileCmd[1:], sourceFile, "-o", executableName)
		}

		if len(compileArgs) > 0 {
			fmt.Printf("  Command: %s %s\n", config.CompileCmd[0], strings.Join(compileArgs, " "))
		}

		fmt.Println("\nExecution step:")
		if ext == ".java" {
			fmt.Printf("  Command: %s %s\n", config.RunCmd[0], config.ClassNameFn(filepath.Base(sourceFile)))
		} else if ext == ".cs" {
			fmt.Printf("  Command: dotnet run\n")
		} else {
			fmt.Printf("  Command: ./%s\n", executableName)
		}

		fmt.Println("\nCleanup step:")
		fmt.Printf("  Would remove: %s\n", executableName)
	} else {
		fmt.Println("\nExecution step:")
		runArgs := append(config.RunCmd[1:], sourceFile)
		fmt.Printf("  Command: %s %s\n", config.RunCmd[0], strings.Join(runArgs, " "))
	}

	fmt.Println("\n✓ Dry run complete")
}

func performBenchmark(sourceFile string, config LanguageConfig, ext string, runs int) {
	fmt.Printf("🔥  Running benchmark with %d iterations...\n", runs)
	fmt.Println(strings.Repeat("=", 50))

	times := make([]time.Duration, runs)
	var totalTime time.Duration

	// Compile once if needed
	var executableName string
	var compiledForBench bool

	if config.IsCompiled {
		executableName = strings.TrimSuffix(sourceFile, filepath.Ext(sourceFile))
		fmt.Printf("Compiling %s...\n", sourceFile)

		var compileArgs []string
		if ext == ".rs" {
			compileArgs = append(config.CompileCmd[1:], sourceFile)
		} else if ext == ".cs" {
			// Handle .NET compilation
			projectDir := strings.TrimSuffix(sourceFile, filepath.Ext(sourceFile))
			if _, err := os.Stat(projectDir); os.IsNotExist(err) {
				cmd := exec.Command("dotnet", "new", "console", "-o", projectDir)
				cmd.Stdout = nil
				cmd.Stderr = os.Stderr
				cmd.Run()
				os.Rename(sourceFile, filepath.Join(projectDir, "Program.cs"))
			}
			os.Chdir(projectDir)
			compileArgs = config.CompileCmd[1:]
		} else {
			compileArgs = append(config.CompileCmd[1:], sourceFile, "-o", executableName)
		}

		cmd := exec.Command(config.CompileCmd[0], compileArgs...)
		cmd.Stdout = nil
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Compilation failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("✓ Compilation successful\n")
		compiledForBench = true
	}

	// Run benchmark iterations
	for i := 0; i < runs; i++ {
		fmt.Printf("Run %d/%d... ", i+1, runs)

		start := time.Now()

		var cmd *exec.Cmd
		if config.IsCompiled {
			if ext == ".java" {
				cmd = exec.Command(config.RunCmd[0], config.ClassNameFn(filepath.Base(sourceFile)))
			} else if ext == ".cs" {
				cmd = exec.Command(config.RunCmd[0], config.RunCmd[1:]...)
			} else if ext == ".rs" {
				cmd = exec.Command("./" + executableName)
			} else {
				cmd = exec.Command(executableName)
			}
		} else {
			runArgs := append(config.RunCmd[1:], sourceFile)
			cmd = exec.Command(config.RunCmd[0], runArgs...)
		}

		cmd.Stdout = nil // Suppress output during benchmark
		cmd.Stderr = nil
		err := cmd.Run()

		elapsed := time.Since(start)
		times[i] = elapsed
		totalTime += elapsed

		if err != nil {
			fmt.Printf("✗ Failed (%v)\n", err)
		} else {
			fmt.Printf("✓ %v\r", elapsed)
		}
	}

	// Clean up if compiled
	if compiledForBench {
		if ext == ".cpp" || ext == ".c" || ext == ".rs" || ext == ".nim" || ext == ".zig" || ext == ".hs" || ext == ".pas" || ext == ".fs" || ext == ".ml" {
			os.Remove(executableName)
			if runtime.GOOS == "windows" {
				os.Remove(executableName + ".exe")
			}
		}
	}

	// Calculate statistics
	sort.Slice(times, func(i, j int) bool { return times[i] < times[j] })

	min := times[0]
	max := times[len(times)-1]
	avg := totalTime / time.Duration(runs)
	median := times[len(times)/2]

	var sumSquaredDiffs float64
	for _, t := range times {
		diff := float64(t - avg)
		sumSquaredDiffs += diff * diff
	}

	// Standard deviation is the square root of variance
	stdDev := time.Duration(math.Sqrt(sumSquaredDiffs / float64(len(times))))

	// Print results
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("  Benchmark Results:")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf("Runs:         %d\n", runs)
	fmt.Printf("Total time:   %v\n", totalTime)
	fmt.Printf("Average:      %v\n", avg)
	fmt.Printf("Median:       %v\n", median)
	fmt.Printf("Min:          %v\n", min)
	fmt.Printf("Max:          %v\n", max)
	fmt.Printf("Std Dev:      %v\n", stdDev)
	fmt.Println(strings.Repeat("=", 50))
}

func executeFile(sourceFile string, config LanguageConfig, ext string) {
	if config.IsCompiled {
		executableName := strings.TrimSuffix(sourceFile, filepath.Ext(sourceFile))
		compileArgs := []string{}
		if ext == ".rs" {
			compileArgs = append(config.CompileCmd[1:], sourceFile)
		} else if ext == ".cs" {
			// For C#, we need to create a project first, then build
			projectDir := strings.TrimSuffix(sourceFile, filepath.Ext(sourceFile))
			if _, err := os.Stat(projectDir); os.IsNotExist(err) {
				fmt.Printf("Creating .NET project in %s...\n", projectDir)
				cmd := exec.Command("dotnet", "new", "console", "-o", projectDir)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err = cmd.Run()
				if err != nil {
					fmt.Printf("Failed to create .NET project: %v\n", err)
					os.Exit(1)
				}
				// Move the source file into the project directory
				fmt.Printf("Moving %s to %s...\n", sourceFile, filepath.Join(projectDir, "Program.cs"))
				os.Rename(sourceFile, filepath.Join(projectDir, "Program.cs"))
			}
			// Change directory to projectDir for dotnet build and run
			os.Chdir(projectDir)
			compileArgs = append(config.CompileCmd[1:])
		} else {
			compileArgs = append(config.CompileCmd[1:], sourceFile, "-o", executableName)
		}

		cmd := exec.Command(config.CompileCmd[0], compileArgs...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		fmt.Printf("Compiling %s...\n", sourceFile)
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Compilation failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Compilation successful.")

		runArgs := []string{executableName}
		if ext == ".java" {
			// For Java, the executable is the class name
			runArgs = []string{config.ClassNameFn(filepath.Base(sourceFile))}
		} else if ext == ".cs" {
			// For C#, dotnet run handles execution from the project directory
			runArgs = config.RunCmd[1:]
			cmd = exec.Command(config.RunCmd[0], runArgs...)
		} else if ext == ".rs" {
			// For Rust, the executable is in the current directory
			cmd = exec.Command("./" + executableName)
		} else {
			cmd = exec.Command(runArgs[0], runArgs[1:]...)
		}

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		fmt.Printf("Running %s...\n", executableName)
		err = cmd.Run()
		if err != nil {
			fmt.Printf("Execution failed: %v\n", err)
			os.Exit(1)
		}

		// Clean up compiled executable for C/C++/Rust/...
		if ext == ".cpp" || ext == ".c" || ext == ".rs" || ext == ".nim" || ext == ".zig" || ext == ".hs" || ext == ".pas" || ext == ".fs" || ext == ".ml" {
			if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
				os.Remove(executableName)
			} else if runtime.GOOS == "windows" {
				os.Remove(executableName + ".exe")
			}
		}

	} else {
		runArgs := append(config.RunCmd[1:], sourceFile)
		cmd := exec.Command(config.RunCmd[0], runArgs...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		fmt.Printf("Running %s...\n", sourceFile)
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Execution failed: %v\n", err)
			os.Exit(1)
		}
	}
}

func checkRuntime(cmdArgs []string) bool {
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	err := cmd.Run()
	return err == nil
}

func installRuntime(cmdArgs []string) bool {
	if len(cmdArgs) == 0 || (len(cmdArgs) == 2 && cmdArgs[0] == "echo" &&
		strings.Contains(cmdArgs[1], "Please install")) {
		fmt.Println(strings.Join(cmdArgs, " "))
		return false // Indicate that automatic installation is not supported or user needs to manually install
	}
	fmt.Printf("Attempting to install %s...\n", cmdArgs[0])
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err == nil
}

func isNumeric(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return len(s) > 0
}

func printHelp() {
	fmt.Println("run - Universal script runner")
	fmt.Println("\nUsage:")
	fmt.Println("  run [options] <source_file>")
	fmt.Println("\nOptions:")
	fmt.Println("  --version, -v        Show version information")
	fmt.Println("  --list, -l           List all supported languages")
	fmt.Println("  --dry-run, -d            Show what would be executed without running")
	fmt.Println("  --time, -t               Measure and display execution time")
	fmt.Println("  --bench [n], -b [n]          Run benchmark (default: 10 iterations)")
	fmt.Println("  --help, -h           Show this help message")
	fmt.Println("\nExamples:")
	fmt.Println("  run script.py                 # Run Python script")
	fmt.Println("  run --time app.js             # Run with execution time")
	fmt.Println("  run --bench 20 program.cpp    # Benchmark with 20 runs")
	fmt.Println("  run --dry-run test.go         # Preview without executing")
	fmt.Println("  run --list                    # Show all supported languages")
}
