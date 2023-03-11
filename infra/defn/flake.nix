{
  inputs.lib.url = github:defn/lib/0.0.57;
  inputs.infra.url = github:defn/lib/infra-0.0.11?dir=cmd/infra;
  outputs = inputs: inputs.lib.cdktfMain rec {
    src = ./.;
    infra = inputs.infra;
    infra_cli = "infra";
  };
}
