Count the number of lines in your source files.

## Usage
```
./source_code_lines -directory {directory}
```

### Options
- directory: Specifies the directory to be analyzed.\
The program reads all the files in the specified directory and subdirectories recursively\
Default behavior is to analyze the current directory

### Ignored Directories
By default, the following directories are ignored
- .git
- vendor
- node_modules
- .idea
- .vscode
- logs
- venv
- tmp
- target
