{
  inputs.lib.url = github:defn/lib/0.0.56;
  outputs = inputs: inputs.lib.goMain rec {
    src = ./.;

    extendShell = ctx: {
      propagatedBuildInputs = [ ctx.pkgs.irssi ];
    };
  };
}
