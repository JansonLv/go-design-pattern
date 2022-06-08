package main

func main() {

	// 一个接入设备（雷电接口的）
	client := &client{}
	mac := &mac{}

	// 我通过雷电接口插入mac电脑
	client.insertLightningConnectorIntoComputer(mac)

	// 但是设备是一个windows，没有雷电接口
	windowsMachine := &windows{}

	// 通过windowsAdapter适配到雷电接口上，简单理解就是windowsAdapter实现了雷电接口，调用了usb接口：雷电转usb数据线的功能
	windowsMachineAdapter := &windowsAdapter{
		windowMachine: windowsMachine,
	}

	client.insertLightningConnectorIntoComputer(windowsMachineAdapter)
}
