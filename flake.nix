{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    dev.url = "github:defn/pkg?dir=dev&ref=v0.0.14";
  };

  outputs =
    { self
    , nixpkgs
    , flake-utils
    , dev
    }:
    flake-utils.lib.eachDefaultSystem (system:
    let
      pkgs = import nixpkgs { inherit system; };
    in
    rec {
      devShell =
        pkgs.mkShell rec {
          buildInputs = with pkgs; [
            dev.defaultPackage.${system}
            defaultPackage
            go
            gotools
            go-tools
            golangci-lint
            gopls
            go-outline
            gopkgs
            nodejs-18_x
            rsync
          ];
        };

      defaultPackage =
        with import nixpkgs { inherit system; };
        stdenv.mkDerivation rec {
          name = "${slug}-${version}";

          slug = "defn-cloud";
          version = "0.0.1";

          src = ./.;

          dontUnpack = true;

          installPhase = ''
            mkdir -p $out/bin
            for a in $src/y/*.go; do
              dst="$(basename "''${a%.go}")"
              cp $a $out/bin/$dst
              sed 's#^// yaegi#\#!/usr/bin/env yaegi#' -i $out/bin/$dst
            done
          '';

          propagatedBuildInputs = [ ];

          meta = with lib; {
            homepage = "https://defn.sh/${slug}";
            description = "nix golang / tilt integration";
            platforms = platforms.linux;
          };
        };
    });
}
