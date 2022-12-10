{
  inputs = {
    dev.url = github:defn/pkg/dev-0.0.10?dir=dev;
    temporalite.url = github:defn/pkg/temporalite-0.3.0-1?dir=temporalite;
    yaegi.url = github:defn/pkg/yaegi-0.14.3-1?dir=yaegi;
  };

  outputs = inputs: inputs.dev.main {
    inherit inputs;

    config = rec {
      slug = "lib";
      version_src = ./VERSION;
      version = builtins.readFile version_src;
    };

    handler = { pkgs, wrap, system }: rec {
      apps.default = {
        type = "app";
        program = "${defaultPackage}/bin/hello";
      };

      defaultPackage = wrap.bashBuilder {
        buildInputs = with pkgs; [
          perl
        ];

        src = ./.;

        installPhase = ''
          set -exu
          mkdir -p $out/bin
          for a in $src/y/*.go; do
            dst="$(basename "''${a%.go}")"
            (
              echo "#!${inputs.yaegi.defaultPackage.${system}}/bin/yaegi"
              echo
              cat $a
            ) > $out/bin/$dst
            chmod 755 $out/bin/$dst
          done
        '';
      };

      packages = {
        go = wrap.nullBuilder {
          propagatedBuildInputs = with pkgs; [
            nodejs-18_x
            terraform
          ];
        };
      };
    };
  };
}
