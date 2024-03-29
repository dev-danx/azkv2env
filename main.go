package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "A brief description of your application",
}

// LoadCmd should have a flag to specify the Key Vault name
var loadCmd = &cobra.Command{
	Use:   "load [names]",
	Short: "Load secrets from Azure Key Vault and set them as environment variables",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Create a new Key Vault client

		keyVaultName := cmd.Flag("key-vault-name").Value.String()
		if keyVaultName == "" {
			log.Fatalf("key-vault-name flag is not set")
			return
		}
		keyVaultUrl := fmt.Sprintf("https://%s.vault.azure.net/", keyVaultName)

		// create credential
		cred, err := azidentity.NewDefaultAzureCredential(nil)
		if err != nil {
			log.Fatalf("failed to create a new credential: %v", err)
			return
		}
		// Establish a connection to the Key Vault client
		client, err := azsecrets.NewClient(keyVaultUrl, cred, nil)
		if err != nil {
			log.Fatalf("failed to establish a connection to the Key Vault client: %v", err)
			return
		}

		// Authenticate with Azure using the managed identity of the VM
		// Create a new authorizer
		// Set the authorizer on the client

		// Load secrets from Azure Key Vault and set them as environment variables
		for _, name := range args {
			fmt.Printf("Loading secret %s\n", name)
			secret, err := client.GetSecret(context.Background(), name, "", nil)
			if err != nil {
				fmt.Println(err)
				return
			}
			os.Setenv(name, *secret.Value)
		}
	},
}

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove all loaded secrets that were set as environment variables",
	Run: func(cmd *cobra.Command, args []string) {
		// Remove all loaded secrets that were set as environment variables
		for _, name := range args {
			os.Unsetenv(name)
		}
	},
}

func init() {
	loadCmd.Flags().String("key-vault-name", "", "The name of the Azure Key Vault")
	rootCmd.AddCommand(loadCmd)
	rootCmd.AddCommand(cleanCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
