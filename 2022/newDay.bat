@echo off
@REM set /p input= "Day Number: "
@REM @REM echo Input is: %input%

mkdir day%1
cd day%1
go mod init aoc/day%1
cd ..
go work use day%1
cd day%1
