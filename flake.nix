{
  inputs = {
    dev.url = github:defn/pkg/dev-0.0.22?dir=dev;
  };

  outputs = inputs: inputs.dev.main rec {
    inherit inputs;

    src = builtins.path { path = ./.; name = config.slug; };

    config = rec {
      slug = builtins.readFile ./SLUG;
      version = builtins.readFile ./VERSION;
      apps = [ "hello" "bye" "api" "infra" ];
    };

    handler = { pkgs, wrap, system, builders }:
      let
        goEnv = pkgs.mkGoEnv {
          pwd = src;
        };

        go = pkgs.lib.genAttrs config.apps
          (name: pkgs.buildGoApplication {
            inherit src;
            pwd = src;
            version = config.version;
            pname = name;
            subPackages = [ "cmd/${name}" ];
          });
      in
      rec {
        defaultPackage = wrap.nullBuilder {
          propagatedBuildInputs = [
            builders.yaegi
            goEnv
            pkgs.gomod2nix
          ];
        };

        packages = pkgs.lib.genAttrs config.apps
          (name: wrap.bashBuilder {
            inherit src;

            installPhase = ''
              mkdir -p $out/bin
              cp ${go.${name}}/bin/${name} $out/bin/${config.slug}
            '';
          });
      };
  };
}
