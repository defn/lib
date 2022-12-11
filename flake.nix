{
  inputs = {
    dev.url = github:defn/pkg/dev-0.0.11-rc22?dir=dev;
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
          pwd = ./.;
          src = ./.;

          goEnv = pkgs.mkGoEnv {
            inherit pwd;
          };

          goHello = pkgs.buildGoApplication {
            inherit pwd;
            inherit src;
            pname = "hello";
            version = "0.1";
            subPackages = [ "cmd/hello" ];
          };
          goBye = pkgs.buildGoApplication {
            inherit pwd;
            inherit src;
            pname = "bye";
            version = "0.1";
            subPackages = [ "cmd/bye" ];
          };
          goApi = pkgs.buildGoApplication {
            inherit pwd;
            inherit src;
            pname = "api";
            version = "0.1";
            subPackages = [ "cmd/api" ];
          };
          #goHello = builders.go { cmd = "cmd/hello"; };
          #goBye = builders.go { cmd = "cmd/bye"; };
        in
        rec {
          defaultPackage = wrap.nullBuilder {
            propagatedBuildInputs = [
              builders.yaegi
              goEnv
            ];
          };

          packages.bins = wrap.bashBuilder {
            src = ./.;

            installPhase = ''
              mkdir -p $out/bin
              cp ${goHello}/bin/hello $out/bin/hello
              cp ${goBye}/bin/bye $out/bin/bye
              cp ${goApi}/bin/api $out/bin/api
            '';

            propagatedBuildInputs = [
              builders.yaegi
            ];
          };
        };
    };
}
