{
  inputs = {
    dev.url = github:defn/pkg/dev-0.0.23?dir=dev;
    terraform.url = github:defn/pkg/terraform-1.3.8-0?dir=terraform;
  };

  outputs = { self, ... }@inputs:
    let
      cdktfMain = { src }:
        let
          s = src;
          cdktf = { system, src, wrap }: wrap.bashBuilder {
            buildInputs = wrap.flakeInputs;

            inherit src;

            installPhase = ''
              mkdir -p $out
              ${self.packages.${system}.infra}/bin/infra
              cp -a cdktf.out/. $out/.
            '';
          };
        in
        inputs.dev.main rec {
          inherit inputs;

          src = builtins.path { path = s; name = (builtins.fromJSON (builtins.readFile "${s}/flake.json")).slug; };

          handler = { pkgs, wrap, system, builders, commands, config }: rec {
            defaultPackage = cdktf { inherit system; inherit src; inherit wrap; };
          };
        };

      goMain = { src }:
        let
          s = src;
        in
        inputs.dev.main rec {
          inherit inputs;

          src = builtins.path { path = s; name = (builtins.fromJSON (builtins.readFile "${s}/flake.json")).slug; };

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
    in
    {
      inherit cdktfMain;
      inherit goMain;
    } // inputs.dev.main rec {
      inherit inputs;

      src = builtins.path { path = ./.; name = (builtins.fromJSON (builtins.readFile "${./.}/flake.json")).slug; };

      handler = { pkgs, wrap, system, builders, commands, config }: {
        defaultPackage = wrap.nullBuilder { };
      };
    };
}
