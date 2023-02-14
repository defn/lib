{
  inputs.lib.url = github:defn/lib/0.0.28;

  outputs = inputs: inputs.lib.main rec {
    src = ./.;
  };
}
