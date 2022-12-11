{
  inputs = {
    dev.url = github:defn/pkg/dev-0.0.11-rc14?dir=dev;
    temporalite.url = github:defn/pkg/temporalite-0.3.0-1?dir=temporalite;
    yaegi.url = github:defn/pkg/yaegi-0.14.3-1?dir=yaegi;
  };

  outputs = inputs:
    let
      src = ./.;
    in
    inputs.dev.main {
      inherit src;
      inherit inputs;

      config = rec {
        slug = "lib";
        version_src = ./VERSION;
        version = builtins.readFile version_src;
      };

      handler = { pkgs, wrap, system, builders }:
        let goHello = builders.go { src = ./.; cmd = "cmd/hello"; buildInputs = [ pkgs.go ]; }; in rec {
          apps.default = {
            type = "app";
            program = "${defaultPackage}/bin/hello";
          };

          defaultPackage = wrap.bashBuilder {
            src = ./.;

            installPhase = ''
              mkdir -p $out/bin
              cp ${goHello} $out/bin/hello
            '';

            propagatedBuildInputs = [
              builders.yaegi
            ];
          };

          packages = rec {
            go = wrap.nullBuilder {
              propagatedBuildInputs = with pkgs; [
                nodejs-18_x
                terraform
              ];
            };
          };
        };
    };
}
