mkdir "%~dp0\build\_workspace_win"

go env -w GOPATH="%~dp0\build\_workspace_win"
go env -w GOBIN="%~dp0\build\bin"

go build -v -o "%~dp0\build\bin\stratum-ping.exe"

pause