package scraper

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"time"

	"github.com/chromedp/chromedp"
)

func init() {
	if !isChromeInstalled() {
		installChrome()
	} else {
		log.Println("Chrome is already installed.")
	}
}

func isChromeInstalled() bool {
	var checkCmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		checkCmd = exec.Command("where", "chrome")
	case "darwin":
		checkCmd = exec.Command("which", "google-chrome")
	default:
		checkCmd = exec.Command("which", "google-chrome")
	}
	if err := checkCmd.Run(); err != nil {
		return false
	}
	return true
}

func installChrome() {
	switch runtime.GOOS {
	case "windows":
		fmt.Println("Downloading Chrome for Windows...")
		err := exec.Command("powershell", "Start-Process", "powershell", "-ArgumentList", "'-Command", "Invoke-WebRequest", "https://dl.google.com/chrome/install/latest/chrome_installer.exe", "-OutFile", "chrome_installer.exe; Start-Process", "chrome_installer.exe", "-Wait'").Run()
		if err != nil {
			log.Fatal("Failed to install Chrome on Windows:", err)
		}
	case "darwin":
		fmt.Println("Installing Chrome on macOS...")
		err := exec.Command("brew", "install", "--cask", "google-chrome").Run()
		if err != nil {
			log.Fatal("Failed to install Chrome on macOS:", err)
		}
	case "linux":
		fmt.Println("Installing Chrome on Linux...")
		err := exec.Command("sudo", "apt-get", "update").Run()
		if err != nil {
			log.Fatal("Failed to update package list:", err)
		}
		err = exec.Command("sudo", "apt-get", "install", "-y", "google-chrome-stable").Run()
		if err != nil {
			log.Fatal("Failed to install Chrome on Linux:", err)
		}
	default:
		fmt.Println("Unsupported OS. Please install Chrome manually.")
	}

	fmt.Println("Chrome installation complete.")
}

// New creates a new scraping context and extracts data from the given URL
func New(url string) string {
	var data string
	var by string

	dataXpath := "/html/body/div[1]/div/div/div[2]/main/div/div/div/div[1]/div/section/div/div/div[1]/div/div/article/div/div/div[3]/div[1]"
	byXpath := "/html/body/div[1]/div/div/div[2]/main/div/div/div/div[1]/div/section/div/div/div[1]/div/div/article/div/div/div[2]/div[2]/div/div/div[1]/div/div/div[1]/div/a/div/div[1]/span/span"

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	log.Println("Scraping:", url)
	log.Println("Creating context...")
	log.Println("Performing scraping...")

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(time.Second*2),
		chromedp.Text(dataXpath, &data, chromedp.BySearch),
		chromedp.Text(byXpath, &by, chromedp.BySearch),
	)
	if err != nil {
		log.Fatalf("Scraping failed: %v", err)
	}

	log.Println("Scraping successful.")
	return data + " by " + by
}
