{
  inputs = {
    dev.url = github:defn/pkg/dev-0.0.23?dir=dev;
  };

  outputs = { self, ... }@inputs: inputs.dev.main rec {
    inherit inputs;

    src = builtins.path { path = ./.; name = builtins.readFile ./SLUG; };

    handler = { pkgs, wrap, system, builders, commands, config }:
      let
        goEnv = pkgs.mkGoEnv { pwd = src; };
        goCmd = pkgs.buildGoApplication {
          inherit src;
          pwd = src;
          pname = config.slug;
          version = config.version;
        };
      in
      rec {
        devShell = wrap.devShell {
          devInputs = [
            goEnv
            pkgs.gomod2nix
          ];
        };

        defaultPackage = wrap.bashBuilder {
          inherit src;

          installPhase = ''
            mkdir -p $out/bin
            ls -ltrhd ${goCmd}/bin/*
            cp ${goCmd}/bin/${config.slug} $out/bin/
          '';
        };
      };
  };
}
