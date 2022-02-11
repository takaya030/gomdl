package studio

import (
	"github.com/chewxy/math32"
)

func (qt *Vec4) QuaternionMatrix(mat *Mat34) {
	mat[0][0] = 1.0 - 2.0 * qt[1] * qt[1] - 2.0 * qt[2] * qt[2]
	mat[1][0] = 2.0 * qt[0] * qt[1] + 2.0 * qt[3] * qt[2]
	mat[2][0] = 2.0 * qt[0] * qt[2] - 2.0 * qt[3] * qt[1]

	mat[0][1] = 2.0 * qt[0] * qt[1] - 2.0 * qt[3] * qt[2]
	mat[1][1] = 1.0 - 2.0 * qt[0] * qt[0] - 2.0 * qt[2] * qt[2]
	mat[2][1] = 2.0 * qt[1] * qt[2] + 2.0 * qt[3] * qt[0]

	mat[0][2] = 2.0 * qt[0] * qt[2] + 2.0 * qt[3] * qt[1]
	mat[1][2] = 2.0 * qt[1] * qt[2] - 2.0 * qt[3] * qt[0]
	mat[2][2] = 1.0 - 2.0 * qt[0] * qt[0] - 2.0 * qt[1] * qt[1]
}

func (p *Vec4) QuaternionSlerp(q Vec4, t float32, qt *Vec4) {

	// decide if one of the quaternions is backwards
	var a, b float32 = 0, 0

	for i := 0; i < 4; i++ {
		a += (p[i] - q[i]) * (p[i] - q[i])
		b += (p[i] + q[i]) * (p[i] + q[i])
	}

	if a > b {
		for i := 0; i < 4; i++ {
			q[i] = -q[i]
		}
	}

	var omega, cosom, sinom, sclp, sclq float32

	cosom = p[0]*q[0] + p[1]*q[1] + p[2]*q[2] + p[3]*q[3]

	if (1.0 + cosom) > 0.000001 {
		if (1.0 - cosom) > 0.000001 {
			omega = math32.Acos(cosom)
			sinom = math32.Sin(omega)
			sclp = math32.Sin((1.0 - t) * omega) / sinom
			sclq = math32.Sin(t * omega) / sinom
		} else {
			sclp = 1.0 - t
			sclq = t
		}

		for i := 0; i < 4; i++ {
			qt[i] = sclp * p[i] + sclq * q[i]
		}
	} else {
		qt[0] = -q[1]
		qt[1] = q[0]
		qt[2] = -q[3]
		qt[3] = q[2]
		sclp = math32.Sin((1.0 - t) * (0.5 * math32.Pi))
		sclq = math32.Sin(t * (0.5 * math32.Pi))
		for i := 0; i < 3; i++ {
			qt[i] = sclp * p[i] + sclq * qt[i]
		}
	}
}

func (angles *Vec3) AngleQuaternion(qt *Vec4) {

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
