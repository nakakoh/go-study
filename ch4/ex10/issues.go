package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"go-study.example.com/ch4/github/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	var withinMonthIssues []github.Issue // 一ヶ月未満
	var withinYearIssues []github.Issue  // 一年未満
	var overYearIssues []github.Issue    // 一年以上前

	now := time.Now()
	oneMonthAgo := now.AddDate(0, -1, 0)
	oneYearAgo := now.AddDate(-1, 0, 0)

	for _, item := range result.Items {
		if item.CreatedAt.After(oneMonthAgo) {
			withinMonthIssues = append(withinMonthIssues, *item)
		} else if item.CreatedAt.After(oneYearAgo) {
			withinYearIssues = append(withinYearIssues, *item)
		} else {
			overYearIssues = append(overYearIssues, *item)
		}
	}
	printIssues(&withinMonthIssues, "created within one month.")
	printIssues(&withinYearIssues, "created within one year.")
	printIssues(&overYearIssues, "created over one year.")

}

func printIssues(issues *[]github.Issue, title string) {
	fmt.Println(title)
	for _, item := range *issues {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
