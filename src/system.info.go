package main

import (
	"log"
	"net/http"
	"runtime"

	info "imuslab.com/arozos/mod/info/hardwareinfo"
	prout "imuslab.com/arozos/mod/prouter"
)

//InitShowSysInformation xxx
func SystemInfoInit() {
	log.Println("Operation System: " + runtime.GOOS)
	log.Println("System Architecture: " + runtime.GOARCH)

	if *allow_hardware_management {
		//Updates 5 Dec 2020, Added permission router
		router := prout.NewModuleRouter(prout.RouterOption{
			AdminOnly:   false,
			UserHandler: userHandler,
			DeniedHandler: func(w http.ResponseWriter, r *http.Request) {
				sendErrorResponse(w, "Permission Denied")
			},
		})

		//Create Info Server Object
		infoServer := info.NewInfoServer(info.ArOZInfo{
			BuildVersion: build_version + "." + internal_version,
			DeviceVendor: deviceVendor,
			DeviceModel:  deviceModel,
			VendorIcon:   "../../" + iconVendor,
			SN:           deviceUUID,
			HostOS:       runtime.GOOS,
			CPUArch:      runtime.GOARCH,
		})
		/*
			if runtime.GOOS == "windows" {
				//this features only working on windows, so display on win at now
				http.HandleFunc("/system/info/getCPUinfo", getCPUinfo)
				http.HandleFunc("/system/info/ifconfig", ifconfig)
				http.HandleFunc("/system/info/getDriveStat", getDriveStat)
				http.HandleFunc("/system/info/usbPorts", getUSB)
				http.HandleFunc("/system/info/getRAMinfo", getRAMinfo)

			} else if runtime.GOOS == "linux" {
				//this features only working on windows, so display on win at now
				http.HandleFunc("/system/info/getCPUinfo", getCPUinfoLinux)
				http.HandleFunc("/system/info/ifconfig", ifconfigLinux)
				http.HandleFunc("/system/info/getDriveStat", getDriveStatLinux)
				http.HandleFunc("/system/info/usbPorts", getUSBLinux)
				http.HandleFunc("/system/info/getRAMinfo", getRAMinfoLinux)
			}
		*/

		router.HandleFunc("/system/info/getCPUinfo", info.GetCPUInfo)
		router.HandleFunc("/system/info/ifconfig", info.Ifconfig)
		router.HandleFunc("/system/info/getDriveStat", info.GetDriveStat)
		router.HandleFunc("/system/info/usbPorts", info.GetUSB)
		router.HandleFunc("/system/info/getRAMinfo", info.GetRamInfo)

		//ArOZ Info do not need permission router
		http.HandleFunc("/system/info/getArOZInfo", infoServer.GetArOZInfo)

		//Register as a system setting
		registerSetting(settingModule{
			Name:     "Host Info",
			Desc:     "System Information",
			IconPath: "SystemAO/info/img/small_icon.png",
			Group:    "Info",
			StartDir: "SystemAO/info/index.html",
		})
	}

	//Register as a system setting
	if fileExists("web/SystemAO/vendor/platform/index.html") {
		registerSetting(settingModule{
			Name:     "Platform Info",
			Desc:     "Platform Information",
			IconPath: "SystemAO/info/img/small_icon.png",
			Group:    "Info",
			StartDir: "SystemAO/vendor/platform/index.html",
		})
	}

}
