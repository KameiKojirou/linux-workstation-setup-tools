package core

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings" // Still needed for robust sed pattern

	"github.com/charmbracelet/log"
)

func InstallDotnetCore() {
	log.Info("Starting Dotnet Core installation...")

	// check if dotnet core is already installed
	if _, err := exec.LookPath("dotnet"); err == nil {
		log.Info("dotnet core is already installed. Skipping installation.")
		log.Info("If you wish to reinstall or update, please use the appropriate commands.")
		return
	}

	// Install the latest version of dotnet core SDK
	log.Info("Updating apt and installing dotnet-sdk-8.0...")
	cmd := exec.Command("sh", "-c", "sudo apt-get update && sudo apt-get install -y dotnet-sdk-8.0")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to install dotnet core SDK: ", err)
	}

	// Install the runtime of dotnet core
	log.Info("Installing aspnetcore-runtime-8.0...")
	cmd = exec.Command("sh", "-c", "sudo apt-get update && sudo apt-get install -y aspnetcore-runtime-8.0")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to install dotnet core runtime: ", err)
	}

	// Add dotnet core to .profile
	log.Info("Adding dotnet core environment variables to ~/.profile...")

	// Append # DOTNET CORE
	cmd = exec.Command("sh", "-c", "echo '# DOTNET CORE' >> ~/.profile")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to add DOTNET CORE comment to ~/.profile: ", err)
	}

	// Append export DOTNET_ROOT
	cmd = exec.Command("sh", "-c", "echo 'export DOTNET_ROOT=$HOME/.dotnet' >> ~/.profile")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to add DOTNET_ROOT to ~/.profile: ", err)
	}

	// Append export PATH
	cmd = exec.Command("sh", "-c", `echo 'export PATH=$PATH:$DOTNET_ROOT:$DOTNET_ROOT/tools' >> ~/.profile`)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to add PATH to ~/.profile: ", err)
	}

	log.Info("Dotnet Core installation complete.")
	log.Info("Please run 'source ~/.profile' or restart your terminal to apply changes.")
}

func UpdateDotnetCore() {
	log.Info("Updating dotnet core...")
	// This relies on the 'dotnet update' command.
	// For apt-installed versions, consider also running `sudo apt-get update && sudo apt-get upgrade`.
	cmd := exec.Command("dotnet", "update")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal("Failed to update dotnet core: ", err)
	}
	log.Info("Dotnet Core update complete.")
}

func UninstallDotnetCore() {
  log.Info("Uninstalling all .NET installations…")

  // 1) Purge all APT packages matching dotnet*, aspnet*, and runtime*
  log.Info("Purging apt packages: dotnet*, aspnet*, dotnet-runtime*, dotnet-hosting*")
  purgePkgs := []string{
    "dotnet-sdk-*",
    "aspnetcore-runtime-*",
    "dotnet-runtime-*",
    "dotnet-hosting-*",
  }
  args := append([]string{"apt-get", "purge", "-y"}, purgePkgs...)
  runSudo(args...)

  // 2) Autoremove and clean up
  log.Info("Running apt autoremove and autoclean")
  runSudo("apt-get", "autoremove", "-y")
  runSudo("apt-get", "autoclean", "-y")

  // 3) Remove Microsoft APT feed & GPG key
  log.Info("Removing Microsoft APT feed & GPG key")
  runSudo("rm", "-f", "/etc/apt/sources.list.d/microsoft-prod.list")
  runSudo("rm", "-f", "/etc/apt/trusted.gpg.d/microsoft.gpg")

  // 4) Delete system-wide dotnet folders & alternatives
  log.Info("Removing system-wide dotnet directories")
  runSudo("rm", "-rf", "/usr/share/dotnet")
  runSudo("rm", "-f", "/etc/alternatives/dotnet")

  // 5) Remove any Snap-installed dotnet
  log.Info("Checking for snap-installed dotnet")
  if _, err := exec.LookPath("snap"); err == nil {
    runSudo("snap", "remove", "dotnet-sdk", "--purge")
  }

  // 6) Remove user-local installations (~/.dotnet & ~/.nuget)
  home, err := os.UserHomeDir()
  if err != nil {
    log.Warn("Could not determine $HOME:", err)
  } else {
    for _, d := range []string{".dotnet", ".nuget"} {
      p := filepath.Join(home, d)
      if _, err := os.Stat(p); err == nil {
        log.Info("Removing ", p)
        if e := os.RemoveAll(p); e != nil {
          log.Warn("Failed to remove ", p, ":", e)
        }
      }
    }
  }

  // 7) Scrub environment-vars from shell rc files
  log.Info("Removing DOTNET_* lines from shell rc files")
  patterns := []string{
    `# DOTNET CORE`,
    `export DOTNET_ROOT=`,
    `export PATH=.*DOTNET_ROOT`,
  }
  shells := []string{".profile", ".bashrc", ".zshrc"}
  for _, f := range shells {
    if home != "" {
      path := filepath.Join(home, f)
      for _, pat := range patterns {
        cmd := exec.Command("sh", "-c",
          `grep -q -E "`+pat+`" "`+path+`" && sed -i -E '/`+pat+`/d' "`+path+`"`)
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        _ = cmd.Run() // we don’t fatally bail on missing files or patterns
      }
    }
  }

  log.Info("Rehashing shell command lookup (hash -r)…")
  exec.Command("sh", "-c", "hash -r").Run()

  log.Info("✔ .NET fully uninstalled.")
  log.Info("You may want to restart your shell or run 'source ~/.profile'.")
}

// runSudo runs “sudo <args…>” with stdout/stderr hooked up.
func runSudo(args ...string) {
  cmd := exec.Command("sudo", args...)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  if err := cmd.Run(); err != nil {
    log.Warn("‘sudo ", strings.Join(args, " "), "’ failed: ", err)
  }
}