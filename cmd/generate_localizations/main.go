package main

import (
	"localizer/pkg/localization"
	"localizer/pkg/logger"
)

func main() {
	err := localization.Update("../localizations/")
	if err != nil {
		logger.Error.Fatalln("Could not update languages:", err)
	}

}
