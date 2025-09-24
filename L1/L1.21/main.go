package main

import "fmt"

type VGA interface {
	SendVGA() string
}

type VGAMonitor struct {
	name string
}

func (m *VGAMonitor) SendVGA() string {
	return fmt.Sprintf("%s displays VGA signal", m.name)
}

type HDMIGraphicCard struct {
	name string
}

func (c *HDMIGraphicCard) SendHDMI() string {
	return fmt.Sprintf("%s outputs HDMI signal", c.name)
}

type HDMItoVGAAdapter struct {
	hdmi *HDMIGraphicCard
}

func (a *HDMItoVGAAdapter) SendVGA() string {
	hdmiSignal := a.hdmi.SendHDMI()
	return fmt.Sprintf("%s -> adapter converts to VGA", hdmiSignal)
}

func Display(vga VGA) {
	fmt.Printf("%s\n", vga.SendVGA())
}

func main() {
	oldMonitor := &VGAMonitor{"LG Flatron E2210S"}

	newHDMICard := &HDMIGraphicCard{"AMD Radeon RX 480"}

	adapter := &HDMItoVGAAdapter{newHDMICard}

	fmt.Printf("=== Direct VGA Monitor ===\n")
	Display(oldMonitor)

	fmt.Printf("=== HDMI Graphic Card via Adapter ===\n")
	Display(adapter)
}
