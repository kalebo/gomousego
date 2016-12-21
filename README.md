# GomouseGo

## Motivation
This is a rewrite of a little utility that my coworker used to prank our old boss. That version was in C# but was lost in the sea of shifting bits. He later rewrote it in node.js however, this required the node runtime and would place a item in the taskbar alerting the 'user' to the prank.

This version has been redesigned to be easier to extend with other prank strategies and to compile as a dependency free binary. It runs without placing a entry in the taskbar.

## Usage
Just hide the binary somewhere and have it called on start up with the desired mode.

For example in the crontab have the entry:

```
0   *   *   *   *   /usr/local/bin/gomousego --mode momentum
```

### Modes / Strategies
You can choose the mode desired by using the `--mode` flag. See below for a description of each strategy.

* `momentum` -- the original demonmouse mode. Cursor will retain it's momentum and keep moving.
* `scroll` -- spins the mouse wheel at random times.


## Build
If you have Go and all the libraries installed just run `go build`.

Make sure to `go get` the dependencies

### Linux library setup
Follow the instructions below relating to your setup.

#### Nix
If you have nix you can just run the following for a complete build environment: 

```
nix-shell -p xorg.libX11 xorg.libXtst xorg.libXext xorg.libXi libpng libxkbcommon go
```

#### Ubuntu / Debian
Install the following packages:

```
sudo apt-get install libx11-dev
sudo apt-get install libxtst-dev libpng++-dev
sudo apt-get install xcb libxcb-xkb-dev x11-xkb-utils libx11-xcb-dev libxkbcommon-x11-dev
sudo apt-get install libxkbcommon-dev
``` 

### Window library setup
I recommend using the below to use as your Mingw installation. Simply clone it to the root of your C drive:

`git clone https://github.com/go-vgo/Mingw.git` 

Then setup your environment :
```
set GOPATH=c:\users\kalebo\repos\gocode
set PATH=C:\mingw\bin;%PATH%
```

## Code

This was written in Go with the awesome [robotgo](https://github.com/go-vgo/robotgo) library.

Overall the code is reasonably idiomatic and uses the strategy pattern in it's design. Because of this it should be pretty easy to add new modes: just implement the `IStrategy` interface on your new struct and add the instantiation call in the switch statement in `main()`.

## Contributors
Any pull requests are welcome.
