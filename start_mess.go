package main

import "fmt"

func startMess() string {
	text := "<b>BOT STARTING</b>"
	for idx := range listCoins {
		text += fmt.Sprintf("\n<code>%s</code>", listCoins[idx].Name)
	}
	text += "\n"

	text += fmt.Sprintf("Base price: <code>%v$</code>\n", base)
	text += fmt.Sprintf("Loss alert: <code>%v</code>\n", lossAlert)
	text += fmt.Sprintf("Profit alert: <code>%v</code>\n", profitAlert)

	return text
}
