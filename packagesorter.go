// Packagesorter is a library that takes certain dimensions and a mass to determine in which stack to put a package
package packagesorter

type KgUnit int
type CmUnit int
type Cm3Unit int

const (
	MINIMUM_BULKY_PACKAGE_VOLUME = 1_000_000
	MINIMUM_BULKY_DIMENSION      = 150
	MINIMUM_HEAVY_MASS           = 20
)

const (
	PACKAGE_STANDARD = "STANDARD"
	PACKAGE_SPECIAL  = "SPECIAL"
	PACKAGE_REJECTED = "REJECTED"
)

// Sort takes a width, height, length for dimensions in cm
// and return a string: the name of the stack where the package should go.
func Sort(width, height, length CmUnit, mass KgUnit) string {

	isBulky := isPackageBulky(width, height, length)
	isHeavy := isPackageHeavy(mass)

	switch {
	case isBulky && isHeavy:
		return PACKAGE_REJECTED
	case isBulky || isHeavy:
		return PACKAGE_SPECIAL
	default:
		return PACKAGE_STANDARD
	}
}

func isMinimumDimensionExceeded(width, height, length CmUnit) bool {
	return width >= MINIMUM_BULKY_DIMENSION || height >= MINIMUM_BULKY_DIMENSION || length >= MINIMUM_BULKY_DIMENSION
}

func isVolumeExceeded(width, height, length CmUnit) bool {
	return Cm3Unit(width*height*length) >= MINIMUM_BULKY_PACKAGE_VOLUME
}

func isPackageHeavy(mass KgUnit) bool {
	return mass >= MINIMUM_HEAVY_MASS
}

func isPackageBulky(width, height, length CmUnit) bool {
	return isMinimumDimensionExceeded(width, height, length) || isVolumeExceeded(width, height, length)
}
