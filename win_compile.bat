echo "windows compile"
@echo off
SET curpath=%cd%
cd ..\..
set GOPATH=%cd%

for /F %%i in ('git describe --tags') do ( set gittag=%%i)
echo GITTAG=%gittag%

if [%BUILDTIME%] == [] (
    set BUILDTIME=%date:~0,4%-%date:~5,2%-%date:~8,2%
    rem set BUILDTIME=%date:~3,4%.%date:~8,2%.%date:~11,2%
)

for /F %%i in ('git rev-parse HEAD') do ( set githash=%%i)
echo GITHASH=%githash%

if [%CURRENTDATE%] == [] (
    set CURRENTDATE=%date:~0,4%.%date:~5,2%.%date:~8,2%
    rem set CURRENTDATE=%date:~3,4%.%date:~8,2%.%date:~11,2%
)

set VERSION=%GITTAG%-%CURRENTDATE%
echo %VERSION%

set "LDFLAG=-X my-help/src/common.Version=%VERSION% -X my-help/src/common.BuildTime=%BUILDTIME% -X my-help/src/common.GitHash=%GITHASH% -X my-help/src/common.Tag=%GITTAG%"

cd %curpath%
set bindir=%curpath%\bin
if not exist %bindir% (
	mkdir %bindir%
)

go build -ldflags "%LDFLAG%" -o %bindir%\my-help.exe %curpath%\src\main.go
