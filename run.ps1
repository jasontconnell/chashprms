$env:GODEBUG=''
go build

For ($i = 0; $i -lt 2000; $i++) {
    .\chashprms "$i is the value"
}

go version
