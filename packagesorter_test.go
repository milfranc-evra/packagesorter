package packagesorter

import (
	"testing"
)

func TestMinimumWidthDimensionNotExceeded(t *testing.T) {

	var width, height, length CmUnit = 149, 20, 30

	isExceeded := isMinimumDimensionExceeded(width, height, length)

	if isExceeded {
		t.Fatalf("Expecting %t, But received %t", false, isExceeded)
	}
}

func TestBareMinimumWidthDimensionExceeded(t *testing.T) {

	var width, height, length CmUnit = 150, 20, 30

	isExceeded := isMinimumDimensionExceeded(width, height, length)

	if !isExceeded {
		t.Fatalf("Expecting %t, But received %t", true, isExceeded)
	}
}

func TestOverBareMinimumWidthDimensionExceeded(t *testing.T) {

	var width, height, length CmUnit = 151, 20, 30

	isExceeded := isMinimumDimensionExceeded(width, height, length)

	if !isExceeded {
		t.Fatalf("Expecting %t, But received %t", true, isExceeded)
	}
}

func TestBareMinimumHeightDimensionNotExceeded(t *testing.T) {

	var width, height, length CmUnit = 10, 149, 10

	isExceeded := isMinimumDimensionExceeded(width, height, length)

	if isExceeded {
		t.Fatalf("Expecting %t, But received %t", true, isExceeded)
	}
}

func TestBareMinimumHeightDimensionExceeded(t *testing.T) {

	var width, height, length CmUnit = 10, 150, 10

	isExceeded := isMinimumDimensionExceeded(width, height, length)

	if !isExceeded {
		t.Fatalf("Expecting %t, But received %t", true, isExceeded)
	}
}

func TestOverBareMinimumHeightDimensionExceeded(t *testing.T) {

	var width, height, length CmUnit = 10, 151, 10

	isExceeded := isMinimumDimensionExceeded(width, height, length)

	if !isExceeded {
		t.Fatalf("Expecting %t, But received %t", true, isExceeded)
	}
}

func TestMinimumLengthDimensionExceeded(t *testing.T) {

	var width, height, length CmUnit = 10, 20, 150

	isExceeded := isMinimumDimensionExceeded(width, height, length)

	if !isExceeded {
		t.Fatalf("Expecting %t, But received %t", true, isExceeded)
	}
}

func TestVolumeBareMinimumExceeded(t *testing.T) {

	isExceeded := isVolumeExceeded(1000, 50, 20)

	if !isExceeded {
		t.Fatalf("Expecting %t, But received %t", true, isExceeded)
	}
}

func TestVolumeOverBareMinimumExceeded(t *testing.T) {

	isExceeded := isVolumeExceeded(1000, 50, 50)

	if !isExceeded {
		t.Fatalf("Expecting %t, But received %t", true, isExceeded)
	}
}

func TestVolumeNotExceeded(t *testing.T) {

	isExceeded := isVolumeExceeded(1000, 10, 20)

	if isExceeded {
		t.Fatalf("Expecting %t, But received %t", false, isExceeded)
	}
}

func TestPackageHeavyForUnderMinimum(t *testing.T) {

	isHeavy := isPackageHeavy(19)

	if isHeavy {
		t.Fatalf("Expecting %t, But received %t", false, isHeavy)
	}

}

func TestPackageHeavyForBareMinimum(t *testing.T) {

	isHeavy := isPackageHeavy(20)

	if !isHeavy {
		t.Fatalf("Expecting %t, But received %t", true, isHeavy)
	}

}

func TestPackageHeavyForOverMinimum(t *testing.T) {

	isHeavy := isPackageHeavy(21)

	if !isHeavy {
		t.Fatalf("Expecting %t, But received %t", true, isHeavy)
	}

}

func TestPackageBulkyViaDimension(t *testing.T) {

	var width, height, length CmUnit = 10, 151, 10

	isBulky := isPackageBulky(width, height, length)

	if !isBulky {
		t.Fatalf("Expecting %t, But received %t", true, isBulky)
	}

}

func TestPackageBulkyViaVolume(t *testing.T) {

	var width, height, length CmUnit = 1000, 50, 50

	isBulky := isPackageBulky(width, height, length)

	if !isBulky {
		t.Fatalf("Expecting %t, But received %t", true, isBulky)
	}
}

func TestSortRejected(t *testing.T) {

	var width, height, length CmUnit = 1000, 50, 50
	var mass KgUnit = 20

	packageStack := Sort(width, height, length, mass)

	if packageStack != PACKAGE_REJECTED {
		t.Fatalf("Expecting %s, But received %s", PACKAGE_REJECTED, packageStack)
	}
}

func TestSortSpecialViaVolume(t *testing.T) {

	var width, height, length CmUnit = 1000, 50, 50
	var mass KgUnit = 10

	packageStack := Sort(width, height, length, mass)

	if packageStack != PACKAGE_SPECIAL {
		t.Fatalf("Expecting %s, But received %s", PACKAGE_SPECIAL, packageStack)
	}

}

func TestSortSpecialViaMass(t *testing.T) {

	var width, height, length CmUnit = 1000, 50, 10
	var mass KgUnit = 10

	packageStack := Sort(width, height, length, mass)

	if packageStack != PACKAGE_SPECIAL {
		t.Fatalf("Expecting %s, But received %s", PACKAGE_SPECIAL, packageStack)
	}

}

func TestSortStandard(t *testing.T) {

	var width, height, length CmUnit = 18, 3, 10
	var mass KgUnit = 15

	packageStack := Sort(width, height, length, mass)

	if packageStack != PACKAGE_STANDARD {
		t.Fatalf("Expecting %s, But received %s", PACKAGE_STANDARD, packageStack)
	}

}
