package helper

func VariantClasses(variant string) string {
	switch variant {
	case "primary":
		return "text-white bg-blue-600 hover:bg-blue-700 focus:ring-blue-500"
	case "secondary":
		return "text-blue-700 bg-blue-100 hover:bg-blue-200 focus:ring-blue-300"
	case "outline":
		return "text-blue-600 border border-blue-600 bg-transparent hover:bg-blue-50 focus:ring-blue-400"
	case "ghost":
		return "text-blue-600 bg-transparent hover:bg-blue-50 focus:ring-blue-300"
	case "danger":
		return "text-white bg-red-600 hover:bg-red-700 focus:ring-red-500"
	case "white-primary":
		return "text-blue-700 bg-white hover:bg-neutral-100 focus:ring-blue-300"
	case "white-outline":
		return "text-white border border-white/60 bg-white/10 hover:bg-white/20 focus:ring-white/40"
	default:
		return "text-white bg-blue-600 hover:bg-blue-700 focus:ring-blue-500"
	}
}
