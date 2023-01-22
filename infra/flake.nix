{
  inputs = {
    dev.url = github:defn/pkg/dev-0.0.22?dir=dev;
    defn-lib.url = github:defn/lib/0.0.11;
  };

  outputs = inputs: { main = inputs.dev.main; } // inputs.dev.main rec {
    inherit inputs;

    src = builtins.path { path = ./.; name = config.slug; };

    config = rec {
      slug = builtins.readFile ./SLUG;
      version = builtins.readFile ./VERSION;
    };

    handler = { pkgs, wrap, system, builders }: rec {
      devShell = wrap.devShell {
        devInputs = (
          [
            defaultPackage
            inputs.defn-lib.packages.${system}.infra
          ] ++
          pkgs.lib.attrsets.mapAttrsToList (name: value: value) commands
        );
      };

      defaultPackage = wrap.nullBuilder {
        propagatedBuildInputs = with pkgs; [
          bashInteractive
        ];
      };

      commands = pkgs.lib.attrsets.mapAttrs
        (name: value: (pkgs.writeShellScriptBin "this-${name}" value))
        scripts;

      scripts = {
        hello = ''
          echo hello
        '';
      };
    };
  };
}
