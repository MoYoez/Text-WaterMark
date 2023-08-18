package addWaterMark

import (
	"fmt"
	"testing"
)

func TestAddMark(t *testing.T) {
	rawName := "我能吞下玻璃而不傷身體"
	encName := "WaterMark"
	encText := AddWaterMarkToText(rawName, encName)
	//	fmt.Print(encText)
	output := GetWaterMark(encText)
	fmt.Print(output)
	if output != encName {
		t.Fatal("work err? rawName: ", rawName, "\nencName:", encName)
	}
}
