package studio

import (
	"github.com/chewxy/math32"
)

func (v *Vec3) VectorLength() float32 {
	return math32.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
}

func (v *Vec3) VectorNormalize() float32 {
	length := v.VectorLength()

	if length > 0.0 {
		ilength := 1/length
		v[0] *= ilength
		v[1] *= ilength
		v[2] *= ilength
	}

	return length
}

func (v1 *Vec3) VectorCompare(v2 *Vec3) bool {
	if v1[0] != v2[0] || v1[1] != v2[1] || v1[2] != v2[2] {
		return false
	}
	return true
}

func (v1 *Vec3) DotProduct(v2 *Vec3) float32 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}

func (v1 *Vec3) DotProductV4(v2 [4]float32) float32 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}

func (v1 *Vec3) CrossProduct(v2 *Vec3, cross *Vec3) {
	cross[0] = v1[1]*v2[2] - v1[2]*v2[1]
	cross[1] = v1[2]*v2[0] - v1[0]*v2[2]
	cross[2] = v1[0]*v2[1] - v1[1]*v2[0]
}

func (in1 *Vec3) VectorTransform(in2 *Mat34, out *Vec3) {
	out[0] = in1.DotProductV4(in2[0]) + in2[0][3]
	out[1] = in1.DotProductV4(in2[1]) + in2[1][3]
	out[2] = in1.DotProductV4(in2[2]) + in2[2][3]
}

// rotate by the inverse of the matrix
func (in1 *Vec3) VectorIRotate(in2 *Mat34, out *Vec3) {
	out[0] = in1[0]*in2[0][0] + in1[1]*in2[1][0] + in1[2]*in2[2][0]
	out[1] = in1[0]*in2[0][1] + in1[1]*in2[1][1] + in1[2]*in2[2][1]
	out[2] = in1[0]*in2[0][2] + in1[1]*in2[1][2] + in1[2]*in2[2][2]
}

func (in *Vec3) VectorScale(scale float32, out *Vec3) {
	out[0] = in[0] * scale
	out[1] = in[1] * scale
	out[2] = in[2] * scale
}

func (in1 *Mat34) ConcatTransforms( in2 *Mat34, out *Mat34 ) {
	out[0][0] = in1[0][0] * in2[0][0] + in1[0][1] * in2[1][0] +
				in1[0][2] * in2[2][0]
	out[0][1] = in1[0][0] * in2[0][1] + in1[0][1] * in2[1][1] +
				in1[0][2] * in2[2][1]
	out[0][2] = in1[0][0] * in2[0][2] + in1[0][1] * in2[1][2] +
				in1[0][2] * in2[2][2]
	out[0][3] = in1[0][0] * in2[0][3] + in1[0][1] * in2[1][3] +
				in1[0][2] * in2[2][3] + in1[0][3]
	out[1][0] = in1[1][0] * in2[0][0] + in1[1][1] * in2[1][0] +
				in1[1][2] * in2[2][0]
	out[1][1] = in1[1][0] * in2[0][1] + in1[1][1] * in2[1][1] +
				in1[1][2] * in2[2][1]
	out[1][2] = in1[1][0] * in2[0][2] + in1[1][1] * in2[1][2] +
				in1[1][2] * in2[2][2]
	out[1][3] = in1[1][0] * in2[0][3] + in1[1][1] * in2[1][3] +
				in1[1][2] * in2[2][3] + in1[1][3]
	out[2][0] = in1[2][0] * in2[0][0] + in1[2][1] * in2[1][0] +
				in1[2][2] * in2[2][0]
	out[2][1] = in1[2][0] * in2[0][1] + in1[2][1] * in2[1][1] +
				in1[2][2] * in2[2][1]
	out[2][2] = in1[2][0] * in2[0][2] + in1[2][1] * in2[1][2] +
				in1[2][2] * in2[2][2]
	out[2][3] = in1[2][0] * in2[0][3] + in1[2][1] * in2[1][3] +
				in1[2][2] * in2[2][3] + in1[2][3]
}

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
