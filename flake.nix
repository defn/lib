{
  inputs = {
    dev.url = github:defn/pkg/dev-0.0.23?dir=dev;
  };

  outputs = inputs: inputs.dev.main rec {
    inherit inputs;

    src = builtins.path { path = ./.; name = builtins.readFile ./SLUG; };

    config = rec {
      apps = [ "hello" "bye" "api" "infra" ];
    };

    handler = { pkgs, wrap, system, builders, commands, config }:
      let
        goEnv = pkgs.mkGoEnv {
          pwd = src;
        };

        goCmd = pkgs.lib.genAttrs config.apps
          (name: pkgs.buildGoApplication {
            inherit src;
            pwd = src;
            version = config.version;
            pname = name;
            subPackages = [ "cmd/${name}" ];
          });

        deploy = {
          deploy = wrap.bashBuilder {
            inherit src;

            installPhase = ''
              mkdir -p $out/bin
              cp nix-entrypoint $out/bin/
            '';
          };
        };
      in
      rec {
        defaultPackage = wrap.nullBuilder {
          propagatedBuildInputs = [
            goEnv
            pkgs.gomod2nix
            deploy.deploy
          ];
        };

        packages = deploy // pkgs.lib.genAttrs config.apps
          (name: wrap.bashBuilder {
            inherit src;

            installPhase = ''
              mkdir -p $out/bin
              cp ${goCmd.${name}}/bin/${name} $out/bin/${name}
            '';
          });

        apps = pkgs.lib.genAttrs config.apps
          (name: {
            type = "app";
            program = "${packages.${name}}/bin/${name}";
          });
      };
  };
}
