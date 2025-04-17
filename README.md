# SKELGO
skelgo si a cli tool that prints a tree of ayoru current directory structure

# Example
```bash
$ skelgo
.
├── cmd
│   └── main.go
├── internal
│   ├── printer.go
│   └── walker.go
├── go.mod
└── README.md
```

# Install
```bash
go install github.com/yourusername/skelgo@latest
```
This will put skelgo into your $GOPATH/bin


# Build form source
```bash
git clone https://github.com/yourusername/skelgo.git
cd skelgo
go build -o skelgo
./skelgo
```