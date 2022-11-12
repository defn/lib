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
        buildInputs = with latest; [
          rsync
          go
          gotools
          go-tools
          golangci-lint
          gopls
          go-outline
          gopkgs
          nodejs-18_x
        ];
        site = import ./config.nix;
      in
      with site;
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

            meta = with pkgs.lib; with site; {
              inherit homepage;
              inherit description;
              platforms = platforms.linux;
            };
          };
      }
    );
}
