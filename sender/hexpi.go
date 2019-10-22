package sender

import (
	"errors"
	"os"
	"os/exec"
	"strconv"

	"github.com/sirupsen/logrus"
)

// Send - sedn to hexpi
func Send(code []int) error {
	path := os.Getenv("HEXPI_PATH")
	if len(path) == 0 {
		logrus.Fatalf("[Sender] Failed get HEXPI_PATH. defined?")
		return errors.New("failed get hexpi_path")
	}

	gpio := os.Getenv("HEXPI_GPIO")
	if len(gpio) == 0 {
		gpio = "2"
	}

	spacedCode := ""
	for v := range code {
		spacedCode += strconv.Itoa(v) + " "
	}

	_, err := exec.Command(path, gpio, spacedCode).Output()
	if err != nil {
		logrus.WithError(err).Errorf("[Sender] Failed send to hexpi.")
		return err
	}
	return nil
}