{ pkgs ? (
    let
      inherit (builtins) fetchTree fromJSON readFile;
      inherit ((fromJSON (readFile ./flake.lock)).nodes) nixpkgs gomod2nix;
    in
    import (fetchTree nixpkgs.locked) {
      overlays = [
        (import "${fetchTree gomod2nix.locked}/overlay.nix")
      ];
    }
  )
}:

pkgs.buildGoApplication {
  pname = "screenshot2";
  version = "0.0.17";
  pwd = ./.;
  src = ./.;
  modules = ./gomod2nix.toml;
}
