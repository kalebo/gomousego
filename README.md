# GomouseGo
---

## Motivation
This is a rewrite of a little utility that my coworker used to prank our old boss. That version was in C# but was lost in the sea of shifting bits. He later rewrote it in node.js however, this required the node runtime and would place a item in the taskbar alerting the 'user' to the prank.

This version has been redesigned to be easier to extend with other prank strategies and to compile as a dependency free binary. It runs without placing a entry in the taskbar.

## Installation
Just hide the binary somewhere and have it called on start up with the desired mode.

For example in the crontab have the entry:

```
0   *   *   *   *   /usr/local/bin/gomousego --mode momentum
```

## Build
If you have Go and all the libraries installed just run `go build`.

Make sure to `go get` the dependencies

### Linux library setup
#### Nix
```
nix-shell -p xorg.libX11 xorg.libXtst xorg.libXext xorg.libXi libpng libxkbcommon go
```

#### Ubuntu 
```
sudo apt-get install libx11-dev
sudo apt-get install libgtkglextmm-x11-dev
sudo apt-get install libghc6-x11-dev
sudo apt-get install libgl1-mesa-swx11-dev
sudo apt-get install xorg-dev

sudo apt-get install libxtst-dev libpng++-dev
``` 

### Window library setup
I recommend using the below to use as your Mingw installation. Simply clone it to the root of your C drive:

`git clone https://github.com/go-vgo/Mingw.git` 

Then setup your enviroment :
```
set GOPATH=c:\users\kalebo\repos\gocode
set PATH=C:\mingw\bin;%PATH%
```

## Code

This was written in Go with the awesome [robotgo](https://github.com/go-vgo/robotgo) library.

Overall the code is reasonably ideomatic and uses the strategy pattern in it's design. Because of this it should be pretty easy to add new modes: just implement the `IStrategy` interface on your new struct and add the instatiation call in the switch statement in `main()`.

## Contributers
Any pull requests are welcome.