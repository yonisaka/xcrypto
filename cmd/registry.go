package cmd

import "github.com/urfave/cli/v2"

// Build is a method
func (cmd *Command) Build() []*cli.Command {
	cmd.registerCLI(cmd.newGenerateKeyPair())
	cmd.registerCLI(cmd.newRSAEncryptor())
	cmd.registerCLI(cmd.newAESEncryptor())
	cmd.registerCLI(cmd.newHashing())

	return cmd.CLI
}
