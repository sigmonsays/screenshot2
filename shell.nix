{ pkgs ? import <nixpkgs> {} }:

with pkgs;

mkShell {
  buildInputs = [
        go_1_20
        gomod2nix
  ];
}
