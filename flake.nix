{
  inputs = {
    pkg.url = github:defn/pkg/0.0.166;
    terraform.url = github:defn/pkg/terraform-1.4.0-0?dir=terraform;
    godev.url = github:defn/pkg/godev-0.0.3?dir=godev;
    nodedev.url = github:defn/pkg/nodedev-0.0.1?dir=nodedev;
  };

  outputs = inputs:
    let
      goMain = caller:
        let
          defaultCaller = {
            extendBuild = ctx: { };
            extendShell = ctx: { };
          } // caller;

          goShell = ctx: ctx.wrap.nullBuilder ({ } // (defaultCaller.extendShell ctx));

          go = ctx: ctx.wrap.bashBuilder ({
            src = caller.src;

            buildInputs = [
              inputs.godev.defaultPackage.${ctx.system}
            ];

            installPhase = ''
              mkdir -p $out/bin
              ls -ltrhd ${ctx.goCmd}/bin/*
              cp ${ctx.goCmd}/bin/${ctx.config.slug} $out/bin/
              mkdir -p $out/share/bash-completion/completions
              $out/bin/${ctx.config.slug} completion bash > $out/share/bash-completion/completions/_${ctx.config.slug}
            '';
          } // (defaultCaller.extendBuild ctx));
        in
        inputs.pkg.main rec {
          src = caller.src;

          defaultPackage = ctx:
            let
              goEnv = ctx.pkgs.mkGoEnv {
                pwd = src;
              };

              goCmd = ctx.pkgs.buildGoApplication rec {
                inherit src;
                pwd = src;
                pname = ctx.config.slug;
                version = ctx.config.version;
              };
            in
            go (ctx // { inherit src; inherit goCmd; });

          devShell = ctx: ctx.wrap.devShell {
            devInputs = [
              ctx.pkgs.gomod2nix
              inputs.godev.defaultPackage.${ctx.system}
              inputs.nodedev.defaultPackage.${ctx.system}
              inputs.terraform.defaultPackage.${ctx.system}
              (goShell ctx)
            ];
          };
        };

      cdktfMain = caller:
        let
          cdktf = ctx: ctx.wrap.bashBuilder
            ({
              src = caller.src;

              buildInputs = [
                inputs.nodedev.defaultPackage.${ctx.system}
                inputs.terraform.defaultPackage.${ctx.system}
              ];

              installPhase = ''
                echo 1
                mkdir -p $out
                ${caller.infra.defaultPackage.${ctx.system}}/bin/${caller.infra_cli}
                cp -a cdktf.out/. $out/.
              '';
            }) // (caller.extend ctx);
        in
        inputs.pkg.main rec {
          src = caller.src;

          defaultPackage = ctx: cdktf (ctx // { inherit src; });

          devShell = ctx: ctx.wrap.devShell {
            devInputs = [
              inputs.nodedev.defaultPackage.${ctx.system}
              inputs.terraform.defaultPackage.${ctx.system}
              caller.infra.defaultPackage.${ctx.system}
            ];
          };
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
