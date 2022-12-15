@echo off
set /p input= "Day Number: "
@REM echo Input is: %input%

mkdir day%input%
cd day%input%
go mod init aoc/day%input%
cd ..
go work use day%input%

