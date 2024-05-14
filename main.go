package main

import (
	"fmt"
	"time"

	// "github.com/Skate-Org/AVS/lib/logging"
	"github.com/Skate-Org/AVS/relayer/db/skateapp/disk"
	libExec "github.com/Skate-Org/AVS/lib/exec"
)

func main() {

	binary := "node"
	args := []string{"solana_client/index.js", "getMessage", "1"}
	libExec.ExecBin(time.Duration(15), binary, args...)

	// logger := logging.NewLoggerWithConsoleWriter()
	// pendingTasks, _ := fetchPendingTasks()

	// disk.SkateAppDB.Exec(fmt.Sprintf(`DELETE FROM %s WHERE taskId=?`, disk.SignedTaskSchema), 49)
	// count := 0
	// for _, task := range pendingTasks {
	// 	// if task.ChainType == 1 && task.ChainId == 0 && task.TaskId == 30 {
	// 	count += 1
	// 	logger.Info("", "count", count, "task", task)
	// 	// }
	// }
}

func fetchPendingTasks() ([]disk.SignedTask, error) {
	query := fmt.Sprintf(`
    SELECT *
    FROM %s s
    WHERE NOT EXISTS (
        SELECT 1 FROM %s c
        WHERE c.taskId = s.taskId AND c.chainId = s.chainId AND c.chainType = s.chainType
    )
  `, disk.SignedTaskSchema, disk.CompletedTaskSchema)
	rows, err := disk.SkateAppDB.Query(query)

	var pendingTasks []disk.SignedTask
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task disk.SignedTask
		var entryid int

		err := rows.Scan(
			&entryid, &task.TaskId, &task.Message, &task.Initiator,
			&task.ChainId, &task.ChainType, &task.Hash, &task.Operator, &task.Signature,
		)
		if err != nil {
			return nil, err
		}
		pendingTasks = append(pendingTasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return pendingTasks, nil
}
