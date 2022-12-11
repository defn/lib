{
  inputs = {
    dev.url = github:defn/pkg/dev-0.0.11-rc21?dir=dev;
    temporalite.url = github:defn/pkg/temporalite-0.3.0-1?dir=temporalite;
    yaegi.url = github:defn/pkg/yaegi-0.14.3-1?dir=yaegi;
  };

  outputs = inputs:
    let
      src = ./.;
    in
    inputs.dev.main {
      inherit src;
      inherit inputs;

      config = rec {
        slug = "lib";
        version_src = ./VERSION;
        version = builtins.readFile version_src;
      };

      handler = ele@{ pkgs, wrap, system, builders }:
        let
          goEnv = pkgs.mkGoEnv { pwd = ./.; };
          goApp = pkgs.buildGoApplication {
            pname = "bleh";
            version = "0.1";
            pwd = ./.;
            src = ./.;
            modules = ./gomod2nix.toml;
          };
          goHello = builders.go { cmd = "cmd/hello"; };
          goBye = builders.go { cmd = "cmd/bye"; };
        in
        rec {
          apps.default = {
            type = "app";
            program = "${defaultPackage}/bin/hello";
          };

          defaultPackage = wrap.bashBuilder {
            src = ./.;

            installPhase = ''
              mkdir -p $out/bin
              cp ${goHello} $out/bin/hello
              cp ${goBye} $out/bin/bye
            '';

            propagatedBuildInputs = [
              builders.yaegi
              pkgs.gomod2nix
              goEnv
            ];
          };
        };
    };
}
