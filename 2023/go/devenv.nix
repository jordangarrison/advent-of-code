{ pkgs, ... }:

{
  dotenv.enable = true;
  # https://devenv.sh/basics/

  # https://devenv.sh/packages/
  packages = with pkgs; [
    git
    go
  ];

  # https://devenv.sh/scripts/
  scripts.aoc-fetch.exec = ''
    __day="$1"
    echo "Fetching Day $__day aoc"
    mkdir -p data
    curl -s https://adventofcode.com/2023/day/$__day/input -H "Cookie: session=$AOC_SESSION_TOKEN" > data/day$__day.txt
  '';

  scripts.aoc-run.exec = ''
    go run main.go $@
  '';

  enterShell = ''
    git --version
    go version
  '';

  # https://devenv.sh/languages/
  languages.nix.enable = true;
  languages.go.enable = true;

  # https://devenv.sh/pre-commit-hooks/
  # pre-commit.hooks.shellcheck.enable = true;

  # https://devenv.sh/processes/
  # processes.ping.exec = "ping example.com";

  # See full reference at https://devenv.sh/reference/options/
}
