{
  inputs = {
    dev.url = github:defn/pkg/dev-0.0.14?dir=dev;
  };

  outputs = inputs: inputs.dev.main rec {
    inherit inputs;

    src = builtins.path { path = ./.; name = config.slug; };

    config = rec {
      slug = builtins.readFile ./SLUG;
      version = builtins.readFile ./VERSION;
    };

    handler = { pkgs, wrap, system, builders }:
      let
        inherit src;
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
