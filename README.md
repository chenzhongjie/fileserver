# fileserver

to receive uploaded file with web page or api

## Usage

**upload files with web page**

open ` http://ip:8080` in any browser.

**upload files with curl**

`curl -F 'file=@abc' http://ip:8080?token=ze4LApGm5T`

**download addr**

`wget --header="token:ze4LApGm5T" http://ip:8080/abc`

**other**

`./fileserver --help`

1. Please replace `ip` with real addr.
2. You can find token in `files/.token` file.
3. The token will be updated every time it starts.
4. Logs will be printed in console.

