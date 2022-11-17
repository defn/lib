{
  inputs = {
    dev.url = github:defn/pkg?dir=dev&ref=v0.0.57;
    temporalite.url = github:defn/pkg?dir=temporalite&ref=v0.0.47;
    tilt.url = github:defn/pkg?dir=tilt&ref=v0.0.47;
    earthly.url = github:defn/pkg?dir=earthly&ref=v0.0.47;
    stern.url = github:defn/pkg?dir=stern&ref=v0.0.56;
    c.url = github:defn/pkg?dir=c&ref=v0.0.56;
    latest.url = github:NixOS/nixpkgs/nixpkgs-unstable;
  };

  outputs = inputs: inputs.dev.main {
    inherit inputs;

    config = rec {
      slug = "cloud";
      version = "0.0.1";
      homepage = "https://github.com/defn/${slug}";
      description = "cloud infra and services";
    };

    handler = { pkgs, wrap, system }:
      let
        latest = import inputs.latest { inherit system; };
      in
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
            terraform
            aws-vault
          ];
        };

        packages = {
          go = wrap.nullBuilder {
            propagatedBuildInputs = with latest; [
              nodejs-18_x
              terraform
            ];
          };
        };
      };
  };
}
