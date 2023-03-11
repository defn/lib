{
  inputs.lib.url = github:defn/lib/0.0.52;
  outputs = inputs: inputs.lib.goMain rec { src = ./.; };
}
