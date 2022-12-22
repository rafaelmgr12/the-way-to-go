package main

func Season(month int) string {
	var season string
	switch month {
	case 1, 2, 12:
		season = "Winter"
	case 3, 4, 5:
		season = "Spring"
	case 6, 7, 8:
		season = "Summer"
	case 9, 10, 11:
		season = "Autumn"
	default:
		season = "Season unknown"
	}

	return season
}
