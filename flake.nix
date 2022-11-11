{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    dev.url = "github:defn/pkg?dir=dev&ref=v0.0.11";
  };

  outputs =
    { self
    , nixpkgs
    , flake-utils
    , dev
    }:
    flake-utils.lib.eachDefaultSystem (system:
    let
      pkgs = import nixpkgs { inherit system; };
    in
    {
      devShell =
        pkgs.mkShell rec {
          buildInputs = with pkgs; [
            dev.defaultPackage.${system}
            go
            gotools
            go-tools
            golangci-lint
            gopls
            go-outline
            gopkgs
            nodejs-18_x
            rsync
          ];
        };

      defaultPackage =
        with import nixpkgs { inherit system; };
        stdenv.mkDerivation rec {
          name = "${slug}-${version}";

          slug = "defn-cloud";
          version = "0.0.1";

          dontUnpack = true;

          installPhase = "mkdir -p $out";

          propagatedBuildInputs = [ ];

          meta = with lib; {
            homepage = "https://defn.sh/${slug}";
            description = "nix golang / tilt integration";
            platforms = platforms.linux;
          };
        };
    });
}
