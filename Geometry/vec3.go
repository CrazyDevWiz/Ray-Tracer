package Geometry

type Vec3 struct {
	x float64
	y float64
	z float64
}

func (vector Vec3) add(other Vec3) Vec3 {
	return Vec3{vector.x + other.x, vector.y + other.y, vector.z + other.z}
}

func (vector Vec3) sub(other Vec3) Vec3 {
	return Vec3{vector.x - other.x, vector.y - other.y, vector.z - other.z}
}

func (vector Vec3) mul(other Vec3) Vec3 {
	return Vec3{vector.x * other.x, vector.y * other.y, vector.z * other.z}
}
