{
  inputs = {
    dev.url = github:defn/pkg?dir=dev&ref=v0.0.50;
    temporalite.url = github:defn/pkg?dir=temporalite&ref=v0.0.47;
    tilt.url = github:defn/pkg?dir=tilt&ref=v0.0.47;
    earthly.url = github:defn/pkg?dir=earthly&ref=v0.0.47;
    latest.url = github:NixOS/nixpkgs/nixpkgs-unstable;
  };

  outputs = inputs:
    inputs.dev.eachDefaultSystem (system:
      let
        site = import ./config.nix;
        pkgs = import inputs.dev.wrapper.nixpkgs { inherit system; };
        wrap = inputs.dev.wrapper.wrap { other = inputs; inherit system; inherit site; };
        latest = import inputs.latest { inherit system; };
      in
      with site;
      rec {
        devShell = wrap.devShell;
        defaultPackage = wrap.bashBuilder {
          src = ./.;

          installPhase = ''
            set -exu
            mkdir -p $out/bin
            for a in $src/y/*.go; do
              dst="$(basename "''${a%.go}")"
              (
                echo "#!/usr/bin/env yaegi"
                echo
                cat $a
              ) > $out/bin/$dst
              chmod 755 $out/bin/$dst
            done
          '';

          propagatedBuildInputs = with latest; [
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
        };

        packages = {
          go = wrap.nullBuilder {
            propagatedBuildInputs = [ latest.bash ];
          };
        };
      }
    );
}
