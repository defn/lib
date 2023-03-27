{
  inputs.lib.url = github:defn/lib/0.0.74;
  outputs = inputs: inputs.lib.goMain rec {
    src = ./.;
  };
}
