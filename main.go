package main

import (
	"encoding/hex"
	"fmt"
	"github.com/Skate-Org/AVS/lib/logging"
	"github.com/Skate-Org/AVS/lib/on-chain/avs"
	"github.com/Skate-Org/AVS/relayer/db/skateapp/disk"
)

func main() {
	logger := logging.NewLoggerWithConsoleWriter()

	pendingTasks, _ := fetchPendingTasks()

	disk.SkateAppDB.Exec(fmt.Sprintf(`DELETE FROM %s WHERE taskId=?`, disk.SignedTaskSchema), 46)
	count := 0
	for _, task := range pendingTasks {
		// if task.ChainType == 1 && task.ChainId == 0 && task.TaskId == 30 {
		count += 1
		logger.Info("", "count", count, "task", task)
		// }
	}

	bytes := avs.TaskData("hello", "0x37D191232D6655D82a7ae6159E8d9D55F303E6B2", 1, 0)
	hString := hex.EncodeToString(bytes)
	logger.Info("string", "s", hString)
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
