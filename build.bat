set GOARCH=amd64
set GOOS=linux
go build -o targets/Hyperion
set GOOS=windows
go build -o targets/Ares-windows.exe
pause