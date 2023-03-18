{
  inputs = {
    pkg.url = github:defn/pkg/0.0.170;
    tired-proxy.url = github:defn/lib/tired-proxy-0.0.7?dir=cmd/tired-proxy;
    caddy.url = github:defn/pkg/caddy-2.6.4-3?dir=caddy;
    webhook.url = github:defn/pkg/webhook-2.8.0?dir=webhook;
  };

  outputs = inputs: inputs.pkg.main rec {
    src = ./.;

    devShell = ctx: ctx.wrap.devShell {
      devInputs = [
        (defaultPackage ctx)
      ];
    };

    defaultPackage = ctx: ctx.wrap.bashBuilder {
      inherit src;

      installPhase = ''
        mkdir -p $out/bin
        cp nix-* $out/bin/
        cp bin/* $out/bin/
      '';

      propagatedBuildInputs = [
        inputs.tired-proxy.defaultPackage.${ctx.system}
        inputs.caddy.defaultPackage.${ctx.system}
        inputs.webhook.defaultPackage.${ctx.system}
        ctx.pkgs.bashInteractive
        ctx.pkgs.curl
        ctx.pkgs.git
        ctx.pkgs.jq
      ];
    };

    packages = ctx: {
      this-gen-key = ctx.pkgs.writeShellScriptBin "this-gen-key" ''
        set -exfu

        nix-store --generate-binary-cache-key binarycache.example.com cache-priv-key.pem cache-pub-key.pem
        chmod 600 cache-priv-key.pem
      '';
    };
  };
}
