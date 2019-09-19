/*
 * Redfish OAPI specification
 *
 * Partial Redfish OAPI specification for a limited client
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
type BootSource string

// List of BootSource
const (
	BOOTSOURCE_NONE BootSource = "None"
	BOOTSOURCE_PXE BootSource = "Pxe"
	BOOTSOURCE_FLOPPY BootSource = "Floppy"
	BOOTSOURCE_CD BootSource = "Cd"
	BOOTSOURCE_USB BootSource = "Usb"
	BOOTSOURCE_HDD BootSource = "Hdd"
	BOOTSOURCE_BIOS_SETUP BootSource = "BiosSetup"
	BOOTSOURCE_UTILITIES BootSource = "Utilities"
	BOOTSOURCE_DIAGS BootSource = "Diags"
	BOOTSOURCE_UEFI_SHELL BootSource = "UefiShell"
	BOOTSOURCE_UEFI_TARGET BootSource = "UefiTarget"
	BOOTSOURCE_SD_CARD BootSource = "SDCard"
	BOOTSOURCE_UEFI_HTTP BootSource = "UefiHttp"
	BOOTSOURCE_REMOTE_DRIVE BootSource = "RemoteDrive"
	BOOTSOURCE_UEFI_BOOT_NEXT BootSource = "UefiBootNext"
)