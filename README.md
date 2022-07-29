# gobuildinfo
Show `debug.BuildInfo` for Go binaries.

## Installation
Install `gobuildinfo`, or update it:

```text
go install github.com/clfs/gobuildinfo@latest
```

Uninstall it:

```text
rm $(which gobuildinfo)
```

## Usage
Run `gobuildinfo` alone to display a usage message:

```text
$ gobuildinfo
usage: gobuildinfo FILE
```

## Example
Run `gobuildinfo` on a Go binary:

```text
$ gobuildinfo /Users/calvin/go/bin/staticcheck
{
  "GoVersion": "go1.18.4",
  "Path": "honnef.co/go/tools/cmd/staticcheck",
  "Main": {
    "Path": "honnef.co/go/tools",
    "Version": "v0.3.2",
    "Sum": "h1:ytYb4rOqyp1TSa2EPvNVwtPQJctSELKaMyLfqNP4+34=",
    "Replace": null
  },
  "Deps": [
    {
      "Path": "github.com/BurntSushi/toml",
      "Version": "v0.4.1",
      "Sum": "h1:GaI7EiDXDRfa8VshkTj7Fym7ha+y8/XxIgD2okUIjLw=",
      "Replace": null
    },
    {
      "Path": "golang.org/x/exp/typeparams",
      "Version": "v0.0.0-20220218215828-6cf2b201936e",
      "Sum": "h1:qyrTQ++p1afMkO4DPEeLGq/3oTsdlvdH4vqZUBWzUKM=",
      "Replace": null
    },
    {
      "Path": "golang.org/x/mod",
      "Version": "v0.6.0-dev.0.20220419223038-86c51ed26bb4",
      "Sum": "h1:6zppjxzCulZykYSLyVDYbneBfbaBIQPYMevg0bEwv2s=",
      "Replace": null
    },
    {
      "Path": "golang.org/x/sys",
      "Version": "v0.0.0-20211019181941-9d821ace8654",
      "Sum": "h1:id054HUawV2/6IGm2IV8KZQjqtwAOo2CYlOToYqa0d0=",
      "Replace": null
    },
    {
      "Path": "golang.org/x/tools",
      "Version": "v0.1.11-0.20220513221640-090b14e8501f",
      "Sum": "h1:OKYpQQVE3DKSc3r3zHVzq46vq5YH7x8xpR3/k9ixmUg=",
      "Replace": null
    }
  ],
  "Settings": [
    {
      "Key": "-compiler",
      "Value": "gc"
    },
    {
      "Key": "CGO_ENABLED",
      "Value": "1"
    },
    {
      "Key": "CGO_CFLAGS",
      "Value": ""
    },
    {
      "Key": "CGO_CPPFLAGS",
      "Value": ""
    },
    {
      "Key": "CGO_CXXFLAGS",
      "Value": ""
    },
    {
      "Key": "CGO_LDFLAGS",
      "Value": ""
    },
    {
      "Key": "GOARCH",
      "Value": "arm64"
    },
    {
      "Key": "GOOS",
      "Value": "darwin"
    }
  ]
}
```