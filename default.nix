{
  pkgs ? import <nixpkgs> { },
  lib ? import <nixpkgs/lib>,
}:
let
  bwVersion = "0.0.1";
  # https://nix.dev/tutorials/working-with-local-files.html#id8
  fs = lib.fileset;
  sourceFiles = fs.unions [
    (fs.fileFilter (file: file.hasExt "go" || file.hasExt "mod" || file.hasExt "sum") ./.)
  ];
in
fs.trace sourceFiles pkgs.stdenv.mkDerivation {
  buildInputs = with pkgs; [
    go
  ];
  pname = "bookworm";
  version = bwVersion;
  src = fs.toSource {
    root = ./.;
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
