# Project 1

## Computation Theory

---

# Videos


# Getting Started

## Installation

For running the project the easy way, you just need to have [Nix](https://nixos.org/download/#nix-install-linux) package manager installed on your system. Something you can do by simply running:

**Linux & Windows**

```bash
$ sudo sh <(curl -L https://nixos.org/nix/install) --daemon
```

**MacOS**

```bash
 $ sh <(curl -L https://nixos.org/nix/install)
```

## Execution

Now that you have Nix in your computer, you can run the rest the excercises of this lab.

The following commands will create a shell environment with all the dependencies you need to run the project, in a similiar fashion as Docker do.

### AST (Abstract Syntax Tree)
By running this, a bunch of images will be created on the `./graphs` representing the AST of each regex

```bash
nix run .#project --experimental-features 'nix-command flakes'

```

## Troubleshooting

Depending on the shell you are running nix, you have to tweak the the command shown above this might be some variants:

```bash
nix run .\#project --experimental-features 'nix-command flakes'
nix run '.#project' --experimental-features 'nix-command flakes'
```
