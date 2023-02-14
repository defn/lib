{
  inputs.lib.url = github:defn/lib/0.0.29;
  outputs = inputs: inputs.lib.main rec { src = ./.; };
}
