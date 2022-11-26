{
  inputs = {
    dev.url = github:defn/pkg?dir=dev&ref=v0.0.57;
    temporalite.url = github:defn/pkg?dir=temporalite&ref=v0.0.47;
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
