package main

import (
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

var DAY = make(map[string]map[string]string)

func main() {

	day := time.Now().AddDate(-1, 0, 0)
	count := 0
	for {
		count++
		day = day.AddDate(0, 0, 1)

		commitAmount := rand.Intn(20) + 5
		for i := 0; i < commitAmount; i++ {
			_, err := exec.Command("git", "commit", "--date", day.Format(time.RFC3339), "--allow-empty", "-m", "heyoooo").CombinedOutput()
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
		}

		log.Println(commitAmount, " >> ", day, " >> TIMESINCE: ", time.Since(day).Hours())
		if time.Since(day).Hours() < 1 {
			log.Println("DONE!")
			break
		}
	}

	// DAY[strconv.Itoa(i)][]

}
