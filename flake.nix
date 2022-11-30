{
  inputs = {
    dev.url = github:defn/pkg?dir=dev&ref=dev-0.0.1;
    temporalite.url = github:defn/pkg?dir=temporalite&ref=v0.0.47;
    latest.url = github:NixOS/nixpkgs/nixpkgs-unstable;
  };

  outputs = inputs: inputs.dev.main {
    inherit inputs;

    config = rec {
      slug = "lib";
      version = "0.0.1";
      homepage = "https://github.com/defn/${slug}";
      description = "cloud ibrary";
    };

    handler = { pkgs, wrap, system }:
      let
        latest = import inputs.latest { inherit system; };
      in
      rec {
        devShell = wrap.devShell;

        apps.default = {
          type = "app";
          program = "${defaultPackage}/bin/hello";
        };

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
