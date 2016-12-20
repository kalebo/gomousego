#!/usr/bin/env bash
nix-shell -p xorg.libX11 xorg.libXtst xorg.libXext xorg.libXi libpng libxkbcommon go
