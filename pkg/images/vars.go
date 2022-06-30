package images

var (
	// Binaries used
	// Debootstrap = "debootstrap"
	Mmdebstrap    = "mmdebstrap"
	QemuImg       = "qemu-img"
	VirtCustomize = "virt-customize"
	GuestFish     = "guestfish"

	Binaries = []string{
		Mmdebstrap,
		QemuImg,
		VirtCustomize,
		GuestFish,
	}

	DefaultConfFile  = "conf.json"
	DefaultImageExt  = "img"
	DefaultImageSize = "8G"
)