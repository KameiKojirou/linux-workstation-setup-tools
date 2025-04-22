package core

import (
	"os"
	"os/exec"
	"github.com/charmbracelet/log"
)


func InstallJava() {
	// check if  java is already installed
	if _, err := exec.LookPath("java"); err == nil {
		log.Info("java is already installed.")
		return
	}
	// sudo apt install default-jdk
	cmd := exec.Command("sudo", "apt", "install", "openjdk-17-jdk")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to install java: ", err)
	}
	// add java to path
	cmd = exec.Command("sh", "-c", "echo 'export PATH=$PATH:/usr/lib/jvm/java-17-openjdk-amd64/bin' >> ~/.bashrc")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to add java to path: ", err)
	}
	// add JAVA_HOME to path
	cmd = exec.Command("sh", "-c", "echo 'export JAVA_HOME=/usr/lib/jvm/java-17-openjdk-amd64' >> ~/.bashrc")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to add JAVA_HOME to path: ", err)
	}
	// add export PATH="$JAVA_HOME/bin:$PATH" to path
	cmd = exec.Command("sh", "-c", "echo 'export PATH=\"$JAVA_HOME/bin:$PATH\"' >> ~/.bashrc")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to add export PATH=\"$JAVA_HOME/bin:$PATH\" to path: ", err)
	}
}

func UninstallJava() {
	cmd := exec.Command("sudo", "apt", "remove", "openjdk-17-jdk")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to uninstall java: ", err)
	}
	// remove java from path
	cmd = exec.Command("sh", "-c", "sed -i '/java/d' ~/.bashrc")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to remove java from path: ", err)
	}
	// remove JAVA_HOME from path
	cmd = exec.Command("sh", "-c", "sed -i '/JAVA_HOME/d' ~/.bashrc")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to remove JAVA_HOME from path: ", err)
	}
	// remove export PATH="$JAVA_HOME/bin:$PATH" from path
	cmd = exec.Command("sh", "-c", "sed -i '/export PATH=\"$JAVA_HOME/bin:$PATH\"/d' ~/.bashrc")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to remove export PATH=\"$JAVA_HOME/bin:$PATH\" from path: ", err)
	}
}


func UpgradeJava() {
	cmd := exec.Command("sudo", "apt", "upgrade", "openjdk-17-jdk")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to upgrade java: ", err)
	}
}





