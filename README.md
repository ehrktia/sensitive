# sensitive

![logo!](./assets/read_image.jpeg)





This is a lib which can be used to hide sensitive values in struct.Mostly in microservices/applications config is managed via external (json/yaml) files. The values are parsed and mapped to struct. Some applications which speaks to external APIs will

require key to be provided. These keys can be managed in config value fields.When logging the config values in public these secrets can be exposed and cause data leaks.

Focus of the lib is to manage data leaks,the lib has a `CreateStruct` method which is

used to perform this operation.

There are 2 input fields required for the operation

1. data structure

2. list of fields which are required to be masked or hidden

### usage

refer to `example/main.go`  for how to use

### Todo

- [ ] expand the functionality to non exported columns

### contributing

please feel free to raise a PR with your changes
