@echo off

set GOOS_ORIG=%GOOS%
set GOARCH_ORIG=%GOARCH%

set GOOS=js
set GOARCH=wasm

cd wasm
go build -ldflags="-s -w" -o ../dist/hello.wasm
cd ..


copy html\*.html dist > nul
copy html\*.js dist > nul


set GOOS=%GOOS_ORIG%
set GOARCH=%GOARCH_ORIG%

go build -ldflags="-s -w" -o bin/


