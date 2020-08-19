
go build

For ($i = 0; $i -lt 100; $i++) {
    Write-Host $i
    .\chashprms "$i is the value"
}