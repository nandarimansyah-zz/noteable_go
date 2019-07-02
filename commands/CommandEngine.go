package commands

import (
	"github.com/nandarimansyah/noteable_go/commands/actions"
	"github.com/nandarimansyah/noteable_go/containers"
	"github.com/nandarimansyah/noteable_go/routes"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// CommandEngine is the structure of cli
type CommandEngine struct {
	rootCmd *cobra.Command
}

// NewCommandEngine the command line boot loader
func NewCommandEngine() *CommandEngine {
	var rootCmd = &cobra.Command{
		Use:   "noteable",
		Short: "noteable command line",
		Long:  "noteable-app command line",
	}
	defer func() {
		r := recover()
		if r != nil {
			log.Error(r)
		}
	}()

	rootCmd.PersistentFlags().StringP("config", "c", "configurations", "the config path location")

	//rootCmd.Execute()
	return &CommandEngine{
		rootCmd: rootCmd,
	}
}

// Run the all command line
func (c *CommandEngine) Run() {

	cn := containers.NewEngine()
	var commands = []*cobra.Command{
		// this for run server by comands
		{
			Use:   "serve",
			Short: "Noteable Listening HTTP server",
			Long:  "Noteable-app Listening HTTP server",
			//Args:  cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				actions.ServerListen(cn.Make(new(routes.Route)).GetRouter())
			},
		},
	}
	for _, command := range commands {
		c.rootCmd.AddCommand(command)
	}
	c.rootCmd.Execute()
}
