package orangepi

import "board-module/genericlinux"

const (
	opzero2 = "OrangePi Zero2"
	op3lts  = "OrangePi 3 LTS"
)

var BoardInfoMappings = map[string]genericlinux.BoardInformation{
	opzero2: {
		// OP zero 2 user manual: https://drive.google.com/drive/folders/1ToDjWZQptABxfiRwaeYW1WzQILM5iwpb
		// Gpio pins can be found on page 147.
		PinDefinitions: []genericlinux.PinDefinition{
			{Name: "3", DeviceName: "gpiochip0", LineNumber: 229, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "5", DeviceName: "gpiochip0", LineNumber: 228, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "7", DeviceName: "gpiochip0", LineNumber: 73, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "11", DeviceName: "gpiochip0", LineNumber: 70, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "13", DeviceName: "gpiochip0", LineNumber: 69, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "15", DeviceName: "gpiochip0", LineNumber: 72, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "19", DeviceName: "gpiochip0", LineNumber: 231, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "21", DeviceName: "gpiochip0", LineNumber: 232, PwmChipSysfsDir: "", PwmID: -1},
			// When we can switch between gpio and pwm, this would have line number 226.
			{Name: "8", DeviceName: "gpiochip0", LineNumber: -1, PwmChipSysfsDir: "300a000.pwm", PwmID: 2},
			// When we can switch between gpio and pwm, this would have line number 227.
			{Name: "10", DeviceName: "gpiochip0", LineNumber: -1, PwmChipSysfsDir: "300a000.pwm", PwmID: 1},
			{Name: "12", DeviceName: "gpiochip0", LineNumber: 75, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "16", DeviceName: "gpiochip0", LineNumber: 79, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "18", DeviceName: "gpiochip0", LineNumber: 78, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "22", DeviceName: "gpiochip0", LineNumber: 71, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "24", DeviceName: "gpiochip0", LineNumber: 233, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "26", DeviceName: "gpiochip0", LineNumber: 7, PwmChipSysfsDir: "", PwmID: -1},
		},
		Compats: []string{"xunlong,orangepi-zero2", "allwinner,sun50i-h616"},
	},
	op3lts: {
		// OP 3 LTS user manual: https://drive.google.com/file/d/1jka7avWnzNeTIQFkk78LoJdygWaGH2iu/view
		// Gpio pins can be found on page 145.
		PinDefinitions: []genericlinux.PinDefinition{
			{Name: "3", DeviceName: "gpiochip1", LineNumber: 122, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "5", DeviceName: "gpiochip1", LineNumber: 121, PwmChipSysfsDir: "", PwmID: -1},
			// When we can switch between gpio and pwm, the line number would be 118.
			{Name: "7", DeviceName: "gpiochip1", LineNumber: -1, PwmChipSysfsDir: "300a000.pwm", PwmID: 0},
			{Name: "8", DeviceName: "gpiochip0", LineNumber: 2, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "10", DeviceName: "gpiochip0", LineNumber: 3, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "11", DeviceName: "gpiochip1", LineNumber: 120, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "12", DeviceName: "gpiochip1", LineNumber: 114, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "13", DeviceName: "gpiochip1", LineNumber: 119, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "15", DeviceName: "gpiochip0", LineNumber: 10, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "16", DeviceName: "gpiochip1", LineNumber: 111, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "18", DeviceName: "gpiochip1", LineNumber: 112, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "19", DeviceName: "gpiochip1", LineNumber: 229, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "21", DeviceName: "gpiochip1", LineNumber: 230, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "22", DeviceName: "gpiochip1", LineNumber: 117, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "23", DeviceName: "gpiochip1", LineNumber: 228, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "24", DeviceName: "gpiochip1", LineNumber: 227, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "26", DeviceName: "gpiochip0", LineNumber: 8, PwmChipSysfsDir: "", PwmID: -1},
		},
		Compats: []string{"xunlong,orangepi-3-lts", "allwinner,sun50i-h6"},
	},
	"pi5": {
		PinDefinitions: []genericlinux.PinDefinition{
			{Name: "3", DeviceName: "gpiochip4", LineNumber: 2, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "5", DeviceName: "gpiochip4", LineNumber: 3, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "7", DeviceName: "gpiochip4", LineNumber: 4, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "8", DeviceName: "gpiochip4", LineNumber: 14, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "10", DeviceName: "gpiochip4", LineNumber: 15, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "11", DeviceName: "gpiochip4", LineNumber: 17, PwmChipSysfsDir: "", PwmID: -1},
			// When we can do GPIO and hardware PWM on the same pin, this will have line number 18.
			{Name: "12", DeviceName: "gpiochip4", LineNumber: -1, PwmChipSysfsDir: "1f00098000.pwm", PwmID: 2},
			{Name: "13", DeviceName: "gpiochip4", LineNumber: 27, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "15", DeviceName: "gpiochip4", LineNumber: 22, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "16", DeviceName: "gpiochip4", LineNumber: 23, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "18", DeviceName: "gpiochip4", LineNumber: 24, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "19", DeviceName: "gpiochip4", LineNumber: 10, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "21", DeviceName: "gpiochip4", LineNumber: 9, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "22", DeviceName: "gpiochip4", LineNumber: 25, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "23", DeviceName: "gpiochip4", LineNumber: 11, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "24", DeviceName: "gpiochip4", LineNumber: 8, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "26", DeviceName: "gpiochip4", LineNumber: 7, PwmChipSysfsDir: "", PwmID: -1},
			// Per https://www.raspberrypi.com/documentation/computers/images/GPIO-duplicate.png
			// Physical pins 27 and 28 (shown in white in that diagram) should not be used for
			// normal GPIO stuff.
			{Name: "29", DeviceName: "gpiochip4", LineNumber: 5, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "31", DeviceName: "gpiochip4", LineNumber: 6, PwmChipSysfsDir: "", PwmID: -1},
			// We'd expect pins 32 and 33 to have hardware PWM support, too, but we haven't gotten
			// that to work yet.
			{Name: "32", DeviceName: "gpiochip4", LineNumber: 12, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "33", DeviceName: "gpiochip4", LineNumber: 13, PwmChipSysfsDir: "", PwmID: -1},
			// When we can do GPIO and hardware PWM on the same pin, this will have line number 19.
			{Name: "35", DeviceName: "gpiochip4", LineNumber: -1, PwmChipSysfsDir: "1f00098000.pwm", PwmID: 3},
			{Name: "36", DeviceName: "gpiochip4", LineNumber: 16, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "37", DeviceName: "gpiochip4", LineNumber: 26, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "38", DeviceName: "gpiochip4", LineNumber: 20, PwmChipSysfsDir: "", PwmID: -1},
			{Name: "40", DeviceName: "gpiochip4", LineNumber: 21, PwmChipSysfsDir: "", PwmID: -1},
		},
		Compats: []string{"raspberrypi,5-model-b", "brcm,bcm2712"},
	},
}
