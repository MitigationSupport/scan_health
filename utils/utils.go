package utils

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

func ColorPrintf(format string) {
	_, _ = color.New().Print(format)
}

func IsStringInStringArray(input string, list []string) bool {
	for _, item := range list {
		if input == item {
			return true
		}
	}

	return false
}

func StringToFloat(input string) (float64, error) {
	return strconv.ParseFloat(input, 64)
}

func FormatDuration(duration time.Duration) string {
	str := duration.String()

	if duration.Hours() > 24 {
		days := int(duration.Hours() / 24)
		remainingHours := duration.Hours() - float64(days*24)
		str = fmt.Sprintf("%dd %dh%s", days, int(remainingHours), strings.Split(str, "h")[1])
	}

	str = strings.Replace(str, "h", "h ", 1)
	str = strings.Replace(str, "m", "m ", 1)

	return str
}

func FormatHumanDurationDays(duration time.Duration) string {
	if duration.Hours() > 24 {
		days := int(duration.Hours() / 24)
		suffix := "day"
		if days > 1 {
			suffix = "days"
		}

		return fmt.Sprintf("over %d %s", days, suffix)
	}

	return "less than a day"
}

func PrintTitle(title string) {
	println()
	color.HiCyan(title)
	fmt.Println(strings.Repeat("=", len(title)))
}

func Top5StringList(items []string) string {
	var processedValues []string
	var itemStrings []string
	countOfTop5Processed := 0

	sort.Strings(items)

	for _, item := range items {
		if !IsStringInStringArray(item, processedValues) {
			count := 0

			for _, found := range items {
				if found == item {
					count++
				}
			}

			processedValues = append(processedValues, item)

			if count == 1 {
				itemStrings = append(itemStrings, fmt.Sprintf("%q", item))
			} else {
				itemStrings = append(itemStrings, fmt.Sprintf("%q (x%d instances)", item, count))
			}

			if len(itemStrings) < 6 {
				countOfTop5Processed += count
			}
		}
	}

	if len(itemStrings) == 1 {
		return itemStrings[0]
	}

	if len(itemStrings) < 6 {
		return strings.Join(itemStrings, ", ")
	}

	remaining := len(items) - countOfTop5Processed

	otherPluralString := ""
	if remaining > 1 {
		otherPluralString = "s"
	}

	return fmt.Sprintf("%s and %d other%s", strings.Join(itemStrings[0:5], ", "), remaining, otherPluralString)
}

func ErrorAndExit(message string, err error) {
	color.HiRed(fmt.Sprintf("Error: %s\n", message))

	if err != nil {
		fmt.Printf("%v\n", err)
	}

	os.Exit(1)
}

func ErrorAndExitWithUsage(message string) {
	color.HiRed(fmt.Sprintf("Error: %s", message))
	print("\nUsage:\n")
	flag.PrintDefaults()
	os.Exit(1)
}
