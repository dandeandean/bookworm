{
  pkgs ? import <nixpkgs> { },
  lib ? import <nixpkgs/lib>,
}:
let
  bwVersion = "0.0.0";
  fs = lib.fileset;
  srcDir = ../.;
  sourceFiles = fs.unions [
    # https://nix.dev/tutorials/working-with-local-files.html#id8
    (fs.fileFilter (file: file.hasExt "go" || file.hasExt "mod" || file.hasExt "sum") srcDir)
  ];
in
fs.trace sourceFiles pkgs.stdenv.mkDerivation {
  buildInputs = with pkgs; [
    go
  ];
  pname = "bookworm";
  version = bwVersion;
  src = fs.toSource {
    root = srcDir;
    fileset = sourceFiles;
  };
  installPhase = ''
    mkdir -p $out/bin
    mkdir -p $out/.go-mod-cache
    # Without this, we get a strange error about a homeless shelter
    export HOME=$out
    go mod download
    go build .
    cp bookworm $out/bin
    echo "BUILDING"
  '';
}
