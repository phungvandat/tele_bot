package main

import "fmt"

func startMess() string {
	var (
		text       = "<b>BOT STARTING</b>: "
		lengthSub1 = len(listCoins) - 1
		isOnly1    = lengthSub1 == 0
	)
	for idx := range listCoins {
		text += fmt.Sprintf("<code>%s - %s</code>", listCoins[idx].Name, symbol)
		if !isOnly1 && idx != lengthSub1 {
			text += ", "
		}
	}
	text += "\n"

	text += fmt.Sprintf("Base price: <code>%v$</code>. Target %v%% \n", base, profitTargetPercentage)
	text += fmt.Sprintf("Loss alert: <code>%v</code>\n", lossAlert)
	text += fmt.Sprintf("Profit alert: <code>%v</code>\n", profitAlert)

	return text
}
