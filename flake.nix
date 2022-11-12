{
  inputs = {
    dev.url = github:defn/pkg?dir=dev&ref=v0.0.22;
    nixpkgs.url = github:NixOS/nixpkgs/nixpkgs-unstable;
  };

  outputs = inputs:
    inputs.dev.wrapper.flake-utils.lib.eachDefaultSystem (system:
      let
        latest = import inputs.nixpkgs { inherit system; };
        pkgs = import inputs.dev.wrapper.nixpkgs { inherit system; };
        wrap = inputs.dev.wrapper.wrap { other = inputs; inherit system; };
        slug = "defn-cloud";
        version = "0.0.1";
        buildInputs = [
          latest.rsync
          latest.go
          latest.gotools
          latest.go-tools
          latest.golangci-lint
          latest.gopls
          latest.go-outline
          latest.gopkgs
          latest.nodejs-18_x
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
