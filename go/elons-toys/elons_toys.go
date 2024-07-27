package elon

import "fmt"

func (c Car) canDrive() bool {
	return c.battery >= c.batteryDrain
}

func (c *Car) Drive() {
	if c.canDrive() {
		c.battery -= c.batteryDrain
		c.distance += c.speed
	}
}

func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

func (c Car) CanFinish(trackDistance int) bool {
	initialDistance := c.distance
	for c.distance < trackDistance+initialDistance && c.canDrive() {
		c.Drive()
	}
	return c.distance >= initialDistance+trackDistance
}
