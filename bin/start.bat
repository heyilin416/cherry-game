@echo off
set "current_dir=%cd%"

echo start nets-server
cd nats-server
start run_nats.bat
cd /d "%current_dir%"

timeout /T 1

if exist "server.exe" (
    del /Q "server.exe"
    echo server.exe has been removed.
) else (
    echo server.exe does not exist.
)

cd /d "%current_dir%"
echo build server.exe
go mod tidy
go build -o server.exe ../nodes

cd /d "%current_dir%"
timeout /T 1

echo start master
start "master" server.exe master --path=./config/game.json --node=master

timeout /T 2

echo start center
start "center" server.exe center --path=./config/game.json --node=center

timeout /T 1

echo start gate
start "gate-1" server.exe gate --path=./config/game.json --node=gate-1

echo start web
start "web-1" server.exe web --path=./config/game.json --node=web-1

echo start hall
start "hall-1" server.exe hall --path=./config/game.json --node=hall-1

timeout /T 2
