{
  inputs = {
    dev.url = github:defn/pkg/dev-0.0.23-rc8?dir=dev;
    defn-lib.url = github:defn/lib/0.0.13;
  };

  outputs = inputs: inputs.dev.main rec {
    inherit inputs;

    src = builtins.path { path = ./.; name = builtins.readFile ./SLUG; };

    handler = { pkgs, wrap, system, builders, commands, config }: rec {
      defaultPackage = wrap.nullBuilder {
        propagatedBuildInputs = with pkgs; [
          bashInteractive
          inputs.defn-lib.packages.${system}.infra
        ];
      };
    };
  };
}
