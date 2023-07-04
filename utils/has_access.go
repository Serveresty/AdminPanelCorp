package utils

func CheckAccess(current []string, target []string) bool {
	var current_manager, current_admin, target_manager, target_admin bool
	for _, elem1 := range current {
		if elem1 == "admin" {
			current_admin = true
		}
		if elem1 == "manager" {
			current_manager = true
		}
	}

	for _, elem2 := range target {
		if elem2 == "admin" {
			target_admin = true
		}
		if elem2 == "manager" {
			target_manager = true
		}
	}

	switch {
	case current_admin == true && target_admin == true:
		return false
	case current_admin == true && target_admin == false:
		return true
	case current_admin == false && target_admin == true:
		return false
	case current_manager == true && target_manager == true:
		return false
	case current_manager == true && target_manager == false:
		return true
	case current_manager == false && target_manager == true:
		return false
	default:
		return false
	}
}
