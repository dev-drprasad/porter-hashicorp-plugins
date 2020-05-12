#!/usr/bin/env bash

EXECUTABLE=${1:-porter}

mkdir dist/$2

cp 'dist/porter-hashicorp-plugins_darwin_amd64/porter-hashicorp-plugins' "dist/$2/hashicorp-darwin-amd64"
cp "dist/porter-hashicorp-plugins_linux_amd64/porter-hashicorp-plugins" "dist/$2/hashicorp-linux-amd64"
# cp "dist/porter-hashicorp-plugins_windows_amd64/porter-hashicorp-plugins.exe" "dist/$2/porter-hashicorp-plugins-windows-amd64.exe"

$EXECUTABLE mixin feed generate -d dist/$2 -f dist/atom.xml -t atom-template.xml
exit 0
