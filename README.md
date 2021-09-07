# Software Identification Tags

The `swid` package provides a golang API for manipulating Software Identification (SWID) Tags as described by [ISO/IEC 19770-2:2015](https://www.iso.org/standard/65666.html), [NISTIR-8060](https://doi.org/10.6028/NIST.IR.8060), as well as by their "concise" counterpart [CoSWID](https://datatracker.ietf.org/doc/draft-ietf-sacm-coswid/).

## Resources

* [Package Documentation](https://pkg.go.dev/github.com/veraison/swid)

## Developer tips

Before doing a PR (and routinely during the dev/test cycle) run
```
make presubmit
```
and check its output to make sure your code coverage figures are in line with the target and that there are no newly introduced lint problems.
