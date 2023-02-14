{
  inputs.lib.url = github:defn/lib/0.0.25;

  outputs = inputs: inputs.lib.main rec {
    src = builtins.path { path = ./.; name = builtins.readFile ./SLUG; };
  };
}
