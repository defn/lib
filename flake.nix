{
  inputs = {
    pkg.url = github:defn/pkg/0.0.158;
    terraform.url = github:defn/pkg/terraform-1.4.0-beta2-1?dir=terraform;
  };

  outputs = { self, ... }@inputs:
    let
      cdktfMain = { src, infra }:
        let
          s = src;
          cdktf = { system, src, wrap, pkgs }: wrap.bashBuilder {
            inherit src;

            buildInputs = wrap.flakeInputs ++ [
              pkgs.nodejs-18_x
            ];

            installPhase = ''
              mkdir -p $out
              ${infra.defaultPackage.${system}}/bin/infra
              cp -a cdktf.out/. $out/.
            '';
          };
        in
        inputs.dev.main rec {
          inherit inputs;

          src = builtins.path { path = s; name = (builtins.fromJSON (builtins.readFile "${s}/flake.json")).slug; };

          handler = { pkgs, wrap, system, builders, commands, config }: rec {
            defaultPackage = cdktf { inherit system; inherit src; inherit wrap; inherit pkgs; };
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
    } // inputs.pkg.main rec {
      src = ./.;
      defaultPackage = ctx: ctx.wrap.nullBuilder { };
    };
}
