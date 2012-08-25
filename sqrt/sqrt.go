package sqrt

func sqrt(x float64) float64 {
    z := x
    old_z := 0.0

    for {
        z -= ((z*z) - x) / (2 * z)
        if old_z == z {
            break
        } else {
            old_z = z
        }
    }

    return z
}
