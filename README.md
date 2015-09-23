# onmove-watch-data-extractor
Geonaute Onmove 510/710 GPS Watch Cross Platform Data Extractor 

export GOPATH=`pwd`
 
```go
cd $GOPATH
go build extractor/parser
go install extractor
```

### Usage
```go
$GOPATH/bin/extractor input/DCKD2647
```

### Run test 
```go 
go test extractor   
``` 