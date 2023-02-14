{
  inputs = {
    dev.url = github:defn/pkg/dev-0.0.23?dir=dev;
    defn-lib.url = github:defn/lib/0.0.15;
    terraform.url = github:defn/pkg/terraform-1.3.4?dir=terraform;
  };

  outputs = inputs: inputs.dev.main rec {
    inherit inputs;

    src = builtins.path { path = ./.; name = builtins.readFile ./SLUG; };

    handler = { pkgs, wrap, system, builders, commands, config }: rec {
      defaultPackage = wrap.bashBuilder {
        buildInputs = with pkgs; [ nodejs-18_x ] ++ wrap.flakeInputs;

        inherit src;

        installPhase = ''
          mkdir -p $out
          ${inputs.defn-lib.packages.${system}.infra}/bin/infra
          cp -a cdktf.out/. $out/.
        '';
      };
    };
  };
}
