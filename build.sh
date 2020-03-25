export GOOS=darwin
export GOARCH=amd64
go build -o web-cli-mac
export GOOS=linux
go build -o web-cli-linux
export GOOS=windows 
go build -o web-cli-win.exe
​