{
  inputs.lib.url = github:defn/lib/0.0.34;
  inputs.infra.url = github:defn/lib/infra-0.0.2?dir=cmd/infra;
  outputs = inputs: inputs.lib.cdktfMain rec {
    src = ./.;
    infra = inputs.infra;
  };
}
