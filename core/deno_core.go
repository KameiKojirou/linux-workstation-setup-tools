package core

import (
	"os"
	"os/exec"
	"github.com/charmbracelet/log"
)


func InstallDeno() {
    // Check if Deno is already installed.
    if _, err := exec.LookPath("deno"); err == nil {
        log.Info("Deno is already installed.")
        return
    }

	log.Info("Installing deno")
    cmd := exec.Command("sh", "-c", "echo '# DENO' >> ~/.profile")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}


	cmd = exec.Command("sh", "-c", "curl -fsSL https://deno.land/x/install/install.sh | sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}


func UninstallDeno() {
    log.Info("Starting Deno uninstallation process")

    // Remove the Deno installation directory.
    cmd := exec.Command("sh", "-c", "rm -rf ~/.deno")
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