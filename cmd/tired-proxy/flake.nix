{
  inputs.lib.url = github:defn/lib/0.0.50;
  outputs = inputs: inputs.lib.goMain rec {
    src = ./.;

    extend = { };
  };
}
