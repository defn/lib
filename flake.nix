{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    home.url = "path:/home/ubuntu/dev";
  };

  outputs =
    { self
    , nixpkgs
    , flake-utils
    , home
    }:
    flake-utils.lib.eachDefaultSystem (system:
    let
      pkgs = nixpkgs.legacyPackages.${system};
    in
    {
      devShell =
        pkgs.mkShell rec {
          buildInputs = with pkgs; [
            home.defaultPackage.${system}
            go
            gotools
            go-tools
            golangci-lint
            gopls
            go-outline
            gopkgs
          ];
        };

      defaultPackage =
        with import nixpkgs { inherit system; };
        stdenv.mkDerivation rec {
          name = "${slug}-${version}";

          slug = "defn-cloud";
          version = "0.0.1";

          src = ./bin;

          dontUnpack = true;

          installPhase = ''
            install -m 0755 -D $src/image.sh $out/bin/hello
            chmod 755 $out/bin/hello
          '';

          meta = with lib; {
            homepage = "https://defn.sh/${slug}";
            description = "nix golang / tilt integration";
            platforms = platforms.linux;
          };
        };
    });
}
