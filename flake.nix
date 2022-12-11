{
  inputs = {
    dev.url = github:defn/pkg/dev-0.0.11-rc22?dir=dev;
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
          modules = ./gomod2nix.toml;

          goEnv = pkgs.mkGoEnv {
            pwd = ./.;
          };

          goHello = pkgs.buildGoApplication {
            pname = "hello";
            version = "0.1";
            inherit modules;
            inherit src;
            pwd = ./cmd/hello;
          };
          goBye = pkgs.buildGoApplication {
            pname = "bye";
            version = "0.1";
            inherit modules;
            inherit src;
            pwd = ./cmd/bye;
          };
          goApi = pkgs.buildGoApplication {
            pname = "api";
            version = "0.1";
            inherit modules;
            inherit src;
            pwd = ./cmd/api;
          };
          #goHello = builders.go { cmd = "cmd/hello"; };
          #goBye = builders.go { cmd = "cmd/bye"; };
        in
        rec {
          defaultPackage = wrap.bashBuilder {
            src = ./.;

            installPhase = ''
              mkdir -p $out/bin
              cp ${goHello}/bin/hello $out/bin/hello
              cp ${goBye}/bin/bye $out/bin/bye
              cp ${goApi}/bin/api $out/bin/api
            '';

            propagatedBuildInputs = [
              builders.yaegi
              goEnv
            ];
          };
        };
    };
}
