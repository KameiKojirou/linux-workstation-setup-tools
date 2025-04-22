package core

import (
  "fmt"
  "os"
  "os/exec"
  "path/filepath"
  "strings"

  "github.com/charmbracelet/log"
)

// InstallAndroidStudio will:
//  1. Add Flathub if needed
//  2. Install Android Studio via Flatpak
//  3. Use the Studio's bundled sdkmanager to pull in
//     platform-tools, build-tools;34.0.0, platforms;android-34,
//     cmdline-tools;latest, cmake;3.10.2.4988404, ndk;23.2.8568313
//  4. Append ANDROID_SDK_ROOT, ANDROID_HOME, JAVA_HOME and PATH
//     exports to ~/.profile (if not already present)
func InstallAndroidStudio() {
  // 1) ensure Flathub is added
  if err := exec.Command(
    "flatpak", "remote-add", "--if-not-exists", "flathub",
    "https://dl.flathub.org/repo/flathub.flatpakrepo",
  ).Run(); err != nil {
    log.Fatal("Adding Flathub failed:", err)
  }

  // 2) install Android Studio
  if err := exec.Command(
    "flatpak", "install", "-y", "flathub", "com.google.AndroidStudio",
  ).Run(); err != nil {
    log.Fatal("Installing Android Studio failed:", err)
  }

  // 3) install SDK/NDK/CMake via sdkmanager inside the Flatpak
  home := os.Getenv("HOME")
  sdkRoot := filepath.Join(
    home, ".var/app/com.google.AndroidStudio/data/Android/Sdk",
  )
  sdkArgs := []string{
    "run", "--command=sdkmanager", "com.google.AndroidStudio",
    "--sdk_root=" + sdkRoot,
    "platform-tools",
    "build-tools;34.0.0",
    "platforms;android-34",
    "cmdline-tools;latest",
    "cmake;3.10.2.4988404",
    "ndk;23.2.8568313",
  }
  cmd := exec.Command("flatpak", sdkArgs...)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  if err := cmd.Run(); err != nil {
    log.Fatal("Installing Android SDK components failed:", err)
  }

  // 4) append exports to ~/.profile
  profile := filepath.Join(home, ".profile")
  lines := []string{
    fmt.Sprintf("export ANDROID_SDK_ROOT=%q", sdkRoot),
    fmt.Sprintf("export ANDROID_HOME=%q", sdkRoot),
    "export PATH=\"$PATH:$ANDROID_SDK_ROOT/platform-tools\"",
    "export PATH=\"$PATH:$ANDROID_SDK_ROOT/cmdline-tools/latest/bin\"",
    fmt.Sprintf("export JAVA_HOME=%q",
      filepath.Join(home,
        ".var/app/com.google.AndroidStudio/data/android-studio/jre")),
    "export PATH=\"$PATH:$JAVA_HOME/bin\"",
  }

  // read existing profile (if any)
  data, err := os.ReadFile(profile)
  if err != nil && !os.IsNotExist(err) {
    log.Fatal("Reading ~/.profile failed:", err)
  }
  existing := string(data)

  // open for append
  f, err := os.OpenFile(profile,
    os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
  if err != nil {
    log.Fatal("Opening ~/.profile failed:", err)
  }
  defer f.Close()

  for _, l := range lines {
    if !strings.Contains(existing, l) {
      if _, err := f.WriteString("\n" + l); err != nil {
        log.Fatal("Writing to ~/.profile failed:", err)
      }
    }
  }

  log.Info("Android Studio + SDK installed and env set up.")
}

// UninstallAndroidStudio uninstalls the Flatpak and removes its SDK data.
func UninstallAndroidStudio() {
  if err := exec.Command(
    "flatpak", "uninstall", "-y", "com.google.AndroidStudio",
  ).Run(); err != nil {
    log.Fatal("Uninstalling Android Studio failed:", err)
  }

  dir := filepath.Join(os.Getenv("HOME"),
    ".var/app/com.google.AndroidStudio")
  if err := os.RemoveAll(dir); err != nil {
    log.Warn("Removing SDK directory failed:", err)
  }

  log.Info("Android Studio uninstalled.")
}

// UpgradeAndroidStudio runs a Flatpak update on Android Studio.
func UpgradeAndroidStudio() {
  if err := exec.Command(
    "flatpak", "update", "-y", "com.google.AndroidStudio",
  ).Run(); err != nil {
    log.Fatal("Upgrading Android Studio failed:", err)
  }
  log.Info("Android Studio upgraded.")
}
