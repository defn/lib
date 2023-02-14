{
  inputs = {
    dev.url = github:defn/pkg/dev-0.0.23?dir=dev;
    lib.url = github:defn/lib/0.0.20;
  };

  outputs = inputs: inputs.dev.main rec {
    inherit inputs;

    src = builtins.path { path = ./.; name = builtins.readFile ./SLUG; };

    handler = { pkgs, wrap, system, builders, commands, config }: rec {
      defaultPackage = inputs.lib.cdktf { inherit src; inherit wrap; };
    };
  };
}
