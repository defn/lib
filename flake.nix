{
  inputs = {
    nixpkgs.url = github:NixOS/nixpkgs/nixpkgs-unstable;
    flake-utils.url = github:numtide/flake-utils;
    wrapper.url = github:defn/pkg?dir=wrapper&ref=v0.0.16;

    dev.url = github:defn/pkg?dir=dev&ref=v0.0.16;
  };

  outputs = inputs:
    inputs.flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import inputs.nixpkgs { inherit system; };
        wrap = inputs.wrapper.wrap { other = inputs; inherit system; inherit pkgs; };
        slug = "defn-cloud";
        version = "0.0.1";
        buildInputs = [
          pkgs.rsync
          pkgs.go
          pkgs.gotools
          pkgs.go-tools
          pkgs.golangci-lint
          pkgs.gopls
          pkgs.go-outline
          pkgs.gopkgs
          pkgs.nodejs-18_x
        ];
      in
      rec {
        devShell = wrap.devShell;
        defaultPackage = pkgs.stdenv.mkDerivation
          rec {
            name = "${slug}-${version}";

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

            propagatedBuildInputs = buildInputs;

            meta = with pkgs.lib; {
              homepage = "https://defn.sh/${slug}";
              description = "nix golang / tilt integration";
              platforms = platforms.linux;
            };
          };
      }
    );
}
