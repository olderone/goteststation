package packet

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	log "github.com/sirupsen/logrus"
	"time"
)

var (
	device      string = "eth0"
	snapshotLen int32  = 1024
	Promiscuous bool   = false
	err         error
	timeout     time.Duration = 30 * time.Second
	handle      *pcap.Handle
)

func OpenDevice()  {
	// Open device
	handle, err = pcap.OpenLive(device, snapshotLen, Promiscuous, timeout)
	if err != nil {log.Fatal(err) }
	defer handle.Close()

	// Use the handle as a packet source to process all packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
		fmt.Println(packet)
	}
}

func FindAllDevices() ([]pcap.Interface, error){
	// Find all devices
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	// Print device information
	fmt.Println("Devices found:")
	for _, device := range devices {
		fmt.Println("\nName: ", device.Name)
		fmt.Println("Description: ", device.Description)
		fmt.Println("Devices addresses: ", device.Description)
		for _, address := range device.Addresses {
			fmt.Println("- IP address:", address.IP)
			fmt.Println("- Subnet mask: ", address.Netmask)
		}
	}
	return devices, err
}