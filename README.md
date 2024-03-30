# sensitive


<p align="center">
<img src="./assets/read_image.jpeg" title="" alt="logo!" data-align="center">
</p>



[![Go Reference](https://pkg.go.dev/badge/github.com/ehrktia/sensitive@v0.1.0.svg)](https://pkg.go.dev/github.com/ehrktia/sensitive@v0.1.0)


This is a lib which can be used to hide sensitive values in struct.Mostly in microservices/applications config is managed via external (json/yaml) files. The values are parsed and mapped to struct. Some applications which speaks to external APIs will require key to be provided. These keys can be managed in config value fields.When logging the config values in public these secrets can be exposed and cause data leaks.

Focus of the lib is to manage data leaks,the lib has a `Mask` method which is

used to perform this operation.

There are 2 input fields required for the operation

1. data structure

2. list of fields which are required to be masked or hidden

### usage

refer to `example/main.go`  for how to use

*Note:-* please set env variable `SET_VALUE` with required value to be
used for masking sensitive value  

### Todo

- [ ] expand the functionality to non exported columns

### contributing

please feel free to raise a PR with your changes
