package reminder

import (
	"fmt"
	"time"
)

func StartReminder(taskID int, description string, due time.Time, notifyCh chan string) {
	timeUntilDue := time.Until(due)

	if timeUntilDue > 0 {
		time.Sleep(timeUntilDue)
	}

	message := fmt.Sprintf("Reminder: Task %d - %s is due!", taskID, description)
	notifyCh <- message
}
