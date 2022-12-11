{
  inputs = {
    dev.url = github:defn/pkg/dev-0.0.11?dir=dev;
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
        version = builtins.readFile ./VERSION;
      };

      handler = { pkgs, wrap, system, builders }:
        let
          src = ./.;
          pwd = src;
          version = builtins.readFile ./VERSION;
          apps = [ "hello" "bye" "api" "infra" ];

          goEnv = pkgs.mkGoEnv {
            inherit pwd;
          };

          go = pkgs.lib.genAttrs apps
            (name: pkgs.buildGoApplication {
              inherit pwd;
              inherit src;
              inherit version;
              pname = name;
              subPackages = [ "cmd/${name}" ];
            });
        in
        rec {
          defaultPackage = wrap.nullBuilder {
            propagatedBuildInputs = [
              builders.yaegi
              goEnv
            ];
          };

          packages = pkgs.lib.genAttrs apps
            (name: wrap.bashBuilder {
              inherit src;

              installPhase = ''
                mkdir -p $out/bin
                cp ${go.${name}}/bin/${name} $out/bin/lib
              '';
            });
        };
    };
}
