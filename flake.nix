{
  inputs = {
    pkg.url = github:defn/pkg/0.0.158;
    terraform.url = github:defn/pkg/terraform-1.4.0-beta2-1?dir=terraform;
  };

  outputs = inputs:
    let
      goMain = caller:
        let
          go = ctx: ctx.wrap.bashBuilder {
            src = caller.src;

            installPhase = ''
              mkdir -p $out/bin
              ls -ltrhd ${ctx.goCmd}/bin/*
              cp ${ctx.goCmd}/bin/${caller.config.slug} $out/bin/
            '';
          };
        in
        inputs.pkg.main rec {
          src = caller.src;

          defaultPackage = ctx:
            let
              goEnv = caller.pkgs.mkGoEnv {
                pwd = caller.src;
              };

              goCmd = caller.pkgs.buildGoApplication rec {
                src = caller.src;
                pwd = src;
                pname = caller.config.slug;
                version = caller.config.version;
              };
            in
            go (ctx // { inherit src; inherit goCmd; });

          devShell = caller.wrap.devShell {
            devInputs = [
              caller.pkgs.gomod2nix
            ];
          };
        };

      cdktfMain = caller:
        let
          cdktf = ctx: ctx.wrap.bashBuilder {
            src = caller.src;

            buildInputs = [
              ctx.pkgs.nodejs-18_x
              inputs.terraform.defaultPackage.${ctx.system}
            ];

            installPhase = ''
              mkdir -p $out
              ${caller.infra.defaultPackage.${ctx.system}}/bin/infra
              cp -a cdktf.out/. $out/.
            '';
          };
        in
        inputs.pkg.main rec {
          src = caller.src;

          defaultPackage = ctx: cdktf (ctx // { inherit src; });
        };
    in
    {
      inherit goMain;
      inherit cdktfMain;
    } // inputs.pkg.main rec {
      src = ./.;
      defaultPackage = ctx: ctx.wrap.nullBuilder { };
    };
}
