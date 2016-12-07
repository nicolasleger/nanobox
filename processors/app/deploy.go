package app

import (
	"fmt"

	"github.com/nanobox-io/nanobox-boxfile"

	generator "github.com/nanobox-io/nanobox/generators/hooks/code"
	"github.com/nanobox-io/nanobox/models"
	"github.com/nanobox-io/nanobox/processors/code"
	"github.com/nanobox-io/nanobox/processors/component"
	"github.com/nanobox-io/nanobox/processors/platform"
	"github.com/nanobox-io/nanobox/processors/provider"
	"github.com/nanobox-io/nanobox/util/display"
	"github.com/nanobox-io/nanobox/util/hookit"
)

// Deploy ...
func Deploy(envModel *models.Env, appModel *models.App) error {

	// init docker client
	if err := provider.Init(); err != nil {
		return fmt.Errorf("failed to init docker client: %s", err.Error())
	}

	// syncronize the services as per the new boxfile
	if err := component.Sync(envModel, appModel); err != nil {
		return fmt.Errorf("failed to sync components: %s", err.Error())
	}

	// if the app is a dev app then we should leave here
	if appModel.Name == "dev" {
		return nil
	}

	// setup the platform services
	if err := platform.Setup(appModel); err != nil {
		return fmt.Errorf("failed to setup platform services: %s", err.Error())
	}

	// create the warehouse config for child processes
	hoarder, _ := models.FindComponentBySlug(appModel.ID, "hoarder")

	warehouseConfig := code.WarehouseConfig{
		BuildID:        "1234",
		WarehouseURL:   hoarder.IPAddr(),
		WarehouseToken: "123",
	}

	// publish the code
	if err := code.Publish(envModel, warehouseConfig); err != nil {
		return fmt.Errorf("unable to publish code: %s", err.Error())
	}

	// start code
	if err := code.Sync(appModel, warehouseConfig); err != nil {
		return fmt.Errorf("failed to add code components: %s", err.Error())
	}

	if err := finalizeDeploy(appModel); err != nil {
		return fmt.Errorf("failed to finalize deploy: %s", err.Error())
	}

	// give the user some helpful information
	display.InfoSimDeploy(appModel.LocalIPs["env"])

	return platform.MistListen(appModel)
}

// update the router and run deploy hooks
func finalizeDeploy(appModel *models.App) error {
	display.OpenContext("Finalizing deploy")
	defer display.CloseContext()

	display.StartTask("Running before_live hooks")
	if err := runDeployHook(appModel, "before_live"); err != nil {
		display.ErrorTask()
		return fmt.Errorf("failed to run before deploy hooks: %s", err.Error())
	}
	display.StopTask()

	// update nanoagent portal
	display.StartTask("Updating router")
	if err := platform.UpdatePortal(appModel); err != nil {
		display.ErrorTask()
		return fmt.Errorf("failed to update router: %s", err.Error())
	}
	display.StopTask()

	display.StartTask("Running after_live hooks")
	if err := runDeployHook(appModel, "after_live"); err != nil {
		display.ErrorTask()
		return fmt.Errorf("failed to run after deloy hooks: %s", err.Error())
	}
	display.StopTask()

	return nil
}

// run the before/after hooks and populate the necessary data
func runDeployHook(appModel *models.App, hookType string) error {
	box := boxfile.New([]byte(appModel.DeployedBoxfile))

	// run the hooks for each service in the boxfile
	for _, componentName := range box.Nodes("code") {

		component, err := models.FindComponentBySlug(appModel.ID, componentName)
		if err != nil {
			// no component for that thing in the database..
			// prolly need to report this error but we might not want to fail
			continue
		}

		if _, err := hookit.DebugExec(component.ID, hookType, generator.DeployPayload(appModel, component), "info"); err != nil {
			return err
		}
	}

	return nil
}
