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
      defaultPackage = wrap.nullBuilder {
        propagatedBuildInputs = with pkgs; [
          bashInteractive
          inputs.defn-lib.packages.${system}.infra
        ];
      };
    };
  };
}
