package core

import (
	"os"
	"os/exec"
	"github.com/charmbracelet/log"
)


func InstallDotnetCore() {
	// check if dotnet core is already installed
	if _, err := exec.LookPath("dotnet"); err == nil {
		log.Info("dotnet core is already installed.")
		return
	}
	// install the latest version of dotnet core
	/* sudo apt-get update && \
  sudo apt-get install -y dotnet-sdk-8.0 */

	cmd := exec.Command("sh", "-c", "sudo apt-get update && sudo apt-get install -y dotnet-sdk-8.0")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to install dotnet core: ", err)
	}
	// install the runtime of dotnet core
	/* sudo apt-get update && \
  sudo apt-get install -y aspnetcore-runtime-8.0 */

	cmd = exec.Command("sh", "-c", "sudo apt-get update && sudo apt-get install -y aspnetcore-runtime-8.0")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to install dotnet core runtime: ", err)
	}

	// add dotnet core to path
	/*
	# DOTNET CORE
	export "DOTNET_ROOT=$HOME/.dotnet" 
	export "PATH=$PATH:$DOTNET_ROOT:$DOTNET_ROOT/tools"
	*/

	cmd = exec.Command("sh", "-c", "echo 'export DOTNET_ROOT=$HOME/.dotnet' >> ~/.bashrc")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to add dotnet core to path: ", err)
	}

	cmd = exec.Command("sh", "-c", "echo 'export PATH=$PATH:$DOTNET_ROOT:$DOTNET_ROOT/tools' >> ~/.bashrc")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to add dotnet core to path: ", err)
	}
}

func UpdateDotnetCore() {
	// update dotnet core
	log.Info("Updating dotnet core")
	cmd := exec.Command("sh", "-c", "dotnet update")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func UninstallDotnetCore() {

	// remove dotnet core from install and path

	log.Info("Uninstalling dotnet core")
	cmd := exec.Command("sh", "-c", "dotnet uninstall")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// remove dotnet core from path
	/*
	export "DOTNET_ROOT=$HOME/.dotnet"
	*/

	cmd = exec.Command("sh", "-c", "sed -i '/DOTNET_ROOT/d' ~/.bashrc")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to remove dotnet core from path: ", err)
	}
	// remove export "PATH=$PATH:$DOTNET_ROOT:$DOTNET_ROOT/tools" from path
	cmd = exec.Command("sh", "-c", "sed -i '/export \"PATH=$PATH:$DOTNET_ROOT:$DOTNET_ROOT/tools\"/d' ~/.bashrc")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to remove export \"PATH=$PATH:$DOTNET_ROOT:$DOTNET_ROOT/tools\" from path: ", err)
	}

	// delete .dotnet folder
	cmd = exec.Command("sh", "-c", "rm -rf $HOME/.dotnet")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to delete .dotnet folder: ", err)
	}

}
