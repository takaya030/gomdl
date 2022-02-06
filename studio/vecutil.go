package studio

import (
	"github.com/chewxy/math32"
)

func AngleQuaternion(angles *Vec3, qt *Vec4) {

	var ang float32
	var sr, sp, sy, cr, cp, cy float32

	ang = angles[2] * 0.5
	sy = math32.Sin(ang)
	cy = math32.Cos(ang)
	ang = angles[1] * 0.5
	sp = math32.Sin(ang)
	cp = math32.Cos(ang)
	ang = angles[0] * 0.5
	sr = math32.Sin(ang)
	cr = math32.Cos(ang)

	qt[0] = sr * cp * cy - cr * sp * sy		// X
	qt[1] = cr * sp * cy + sr * cp * sy		// Y
	qt[2] = cr * cp * sy - sr * sp * cy		// Z
	qt[3] = cr * cp * cy + sr * sp * sy		// W
}
