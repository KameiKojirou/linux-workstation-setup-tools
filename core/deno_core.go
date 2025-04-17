package core

import (
	"os"
	"os/exec"
	"github.com/charmbracelet/log"
)


func InstallDeno() {
	log.Info("Installing deno")
	cmd := exec.Command("sh", "-c", "curl -fsSL https://deno.land/x/install/install.sh | sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}


func UninstallDeno() {
    log.Info("Starting Deno uninstallation process")

    // Attempt to uninstall any Deno-installed scripts (if applicable)
    // Note: `deno uninstall` typically expects a script name.
    cmd := exec.Command("sh", "-c", "deno uninstall")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        log.Warn("deno uninstall: ", err)
    }

    // Remove the Deno installation directory.
    cmd = exec.Command("sh", "-c", "rm -rf ~/.deno")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        log.Fatal("removing ~/.deno failed: ", err)
    }

    // Remove references to Deno in shell config files.
    cmd = exec.Command("sh", "-c", "sed -i '/deno/d' ~/.profile")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        log.Fatal("cleaning ~/.profile failed: ", err)
    }

    cmd = exec.Command("sh", "-c", "sed -i '/deno/d' ~/.bashrc")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        log.Fatal("cleaning ~/.bashrc failed: ", err)
    }

    log.Info("Deno uninstallation completed")
}



func UpgradeDeno() {
	log.Info("Updating deno")
	cmd := exec.Command("sh", "-c", "deno upgrade")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}