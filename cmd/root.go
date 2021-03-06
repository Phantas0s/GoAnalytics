package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Phantas0s/devdash/internal"
	"github.com/Phantas0s/devdash/internal/platform"
	"github.com/spf13/cobra"
)

var (
	// Used for flags
	file    string
	logpath string
	debug   bool

	rootCmd = &cobra.Command{
		Use:   "devdash",
		Short: "DevDash is a highly configurable terminal dashboard for developers and creators",
		Long:  `DevDash is a highly flexible terminal dashboard for developers and creators, which allows you to gather and refresh the data you really need from Google Analytics, Google Search Console, Github, TravisCI, and more.`,
		Run: func(cmd *cobra.Command, args []string) {
			run(args)
		},
	}
)

func init() {
	rootCmd.Flags().StringVarP(&file, "config", "c", "", "A valid dashboard configuration")
	// TODO logger
	// rootCmd.Flags().StringVarP(&logpath, "logpath", "l", "", "Path for logging")
	rootCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Debug Mode - doesn't display graph")
	rootCmd.AddCommand(listCmd())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error: " + err.Error())
		os.Exit(1)
	}
}

func run(args []string) {
	termui, err := platform.NewTermUI(debug)
	if err != nil {
		fmt.Println(err)
	}

	tui := internal.NewTUI(termui)
	defer tui.Close()

	if file == "" && len(args) > 0 {
		file = args[0]
	}

	cfg := mapConfig(file)
	hotReload := make(chan time.Time)
	tui.AddKHotReload(cfg.KHotReload(), hotReload)
	tui.AddKQuit(cfg.KQuit())

	// first display
	build(file, tui)

	ticker := time.NewTicker(time.Duration(cfg.RefreshTime()) * time.Second)

	go func() {
		for {
			hotReload <- <-ticker.C
		}
	}()

	go func() {
		for hr := range hotReload {
			tui.HotReload()
			build(file, tui)
			if debug {
				fmt.Println("Last reload: " + hr.Format("2006-01-02 15:04:05"))
			}
		}
	}()

	align := time.NewTicker(time.Duration(3) * time.Second)
	go func() {
		for range align.C {
			tui.Align()
			tui.Render()
		}
	}()

	tui.Loop()
}

// build every services present in the configuration
func build(file string, tui *internal.Tui) {
	cfg := mapConfig(file)
	for _, p := range cfg.Projects {
		rows, sizes := p.OrderWidgets()
		project := internal.NewProject(p.Name, p.NameOptions, rows, sizes, p.Themes, tui)

		gaService := p.Services.GoogleAnalytics
		if !gaService.empty() {
			gaWidget, err := internal.NewGaWidget(gaService.Keyfile, gaService.ViewID)
			if err != nil {
				internal.DisplayError(tui, err)()
			}
			project.WithGa(gaWidget)
		}

		gscService := p.Services.GoogleSearchConsole
		if !gscService.empty() {
			gscWidget, err := internal.NewGscWidget(gscService.Keyfile, gscService.Address)
			if err != nil {
				internal.DisplayError(tui, err)()
			}
			project.WithGoogleSearchConsole(gscWidget)
		}

		monService := p.Services.Monitor
		if !monService.empty() {
			monWidget, err := internal.NewMonitorWidget(monService.Address)
			if err != nil {
				internal.DisplayError(tui, err)()
			}
			project.WithMonitor(monWidget)
		}

		githubService := p.Services.Github
		if !githubService.empty() {
			githubWidget, err := internal.NewGithubWidget(
				githubService.Token,
				githubService.Owner,
				githubService.Repository,
			)
			if err != nil {
				internal.DisplayError(tui, err)()
			}
			project.WithGithub(githubWidget)
		}

		travisService := p.Services.TravisCI
		if !travisService.empty() {
			travisWidget := internal.NewTravisCIWidget(travisService.Token)
			project.WithTravisCI(travisWidget)
		}

		feedlyService := p.Services.Feedly
		if !feedlyService.empty() {
			feedlyWidget := internal.NewFeedlyWidget(feedlyService.Address)
			project.WithFeedly(feedlyWidget)
		}

		gitService := p.Services.Git
		if !gitService.empty() {
			gitWidget := internal.NewGitWidget(gitService.Path)
			project.WithGit(gitWidget)
		}

		remoteHostService := p.Services.RemoteHost
		if !remoteHostService.empty() {
			remoteHostWidget, err := internal.NewHostWidget(
				remoteHostService.Username,
				remoteHostService.Address,
			)
			if err != nil {
				fmt.Println(err)
				internal.DisplayError(tui, err)()
			}
			project.WithRemoteHost(remoteHostWidget)
		}

		localhost, err := internal.NewHostWidget("localhost", "localhost")
		if err != nil {
			fmt.Println(err)
			internal.DisplayError(tui, err)()
		}
		project.WithLocalhost(localhost)

		// TODO choice between concurency and non concurency
		// renderFuncs := project.CreateNonConcWidgets()
		renderFuncs := project.CreateWidgets()
		if !debug {
			project.Render(renderFuncs)
		}
	}
}

// TODO - Wrap logger. If logger nil, drop the message
func InitLoggerFile(logpath string) *log.Logger {
	if logpath == "" {
		return nil
	}

	file, err := os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	return log.New(file, "", 0)
}
