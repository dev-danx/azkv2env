```markdown
# Azure Key Vault to Environment Variables

This is a Go application that loads secrets from Azure Key Vault and sets them as environment variables.

## Prerequisites

- Go 1.16 or later
- Azure account
- Azure Key Vault

## Usage

Set the `KEY_VAULT_NAME` environment variable to the name of your Azure Key Vault.

```sh
export KEY_VAULT_NAME=your-key-vault-name
```

Then, run the application with the names of the secrets you want to load:

```sh
go run main.go load secret1 secret2 secret3
```

The application will load the specified secrets from Azure Key Vault and set them as environment variables.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the terms of the MIT license.
```

Please replace "your-key-vault-name" and "secret1 secret2 secret3" with your actual Azure Key Vault name and the names of the secrets you want to load.