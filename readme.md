# cpsub
Copy files with specific extension in subdirectories while keeping the actual directory structure. Something either `cp` can't do, or I'm too stupid to figure out.

## Install
```
go get github.com/kraxarn/cpsub
```
Binary is stored in:
```
`go env GOPATH`/bin/cpsub
```
## Usage
```
cpsub <source-directory> <target-directory> <file-extension>
```