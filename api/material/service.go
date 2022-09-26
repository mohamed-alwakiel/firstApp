package material

// get all materials
func GetMaterialsService() ([]Material, error) {
	return FindMaterials()
}

// create new material
func CreateMaterialService(input Input) (Material, error) {
	return Store(input)
}
