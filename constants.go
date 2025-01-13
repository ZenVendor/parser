package parser

const (
	INVALID_ARG = iota
	V_ADD
	V_COMPLETE
	V_CONFIGURE
	V_COUNT
	V_DELETE
	V_HELP
	V_LIST
	V_REOPEN
	V_SHOW
	V_UPDATE
	V_VERSION
	K_COMMENT
	K_DUEDATE
	K_GROUP
	K_ID
	K_LONG
	K_PRIORITY
	K_SHORT
	A_ALL
	A_CLOSED
	A_DUE
	A_LOCAL
	A_OPEN
	A_OVERDUE
	A_RESET
	A_SUBTASK
)

func verbMap() map[string]int {
	return map[string]int{
		"add":       V_ADD,
		"a":         V_ADD,
		"complete":  V_COMPLETE,
		"c":         V_COMPLETE,
		"configure": V_CONFIGURE,
		"count":     V_COUNT,
		"delete":    V_DELETE,
		"help":      V_HELP,
		"h":         V_HELP,
		"list":      V_LIST,
		"l":         V_LIST,
		"reopen":    V_REOPEN,
		"show":      V_SHOW,
		"s":         V_SHOW,
		"update":    V_UPDATE,
		"u":         V_UPDATE,
		"version":   V_VERSION,
		"v":         V_VERSION,
	}
}
func argMap() map[string]int {
	return map[string]int{
		"--all":     A_ALL,
		"-a":        A_ALL,
		"--closed":  A_CLOSED,
		"-c":        A_CLOSED,
		"--due":     A_DUE,
		"-d":        A_DUE,
		"--local":   A_LOCAL,
		"--open":    A_OPEN,
		"-o":        A_OPEN,
		"--overdue": A_OVERDUE,
		"-od":       A_OVERDUE,
		"--reset":   A_RESET,
		"--subtask": A_SUBTASK,
		"-s":        A_SUBTASK,
	}
}
func kwargMap() map[string]int {
	return map[string]int{
		"--comment":  K_COMMENT,
		"-c":         K_COMMENT,
		"--due":      K_DUEDATE,
		"-d":         K_DUEDATE,
		"--group":    K_GROUP,
		"-g":         K_GROUP,
		"--id":       K_ID,
		"--long":     K_LONG,
		"-l":         K_LONG,
		"--priority": K_PRIORITY,
		"-p":         K_PRIORITY,
		"--short":    K_SHORT,
		"-s":         K_SHORT,
	}
}
func verbValueMap() map[int]int {
	return map[int]int{
		V_ADD:      K_SHORT,
		V_COMPLETE: K_ID,
		V_DELETE:   K_ID,
		V_REOPEN:   K_ID,
		V_SHOW:     K_ID,
		V_UPDATE:   K_ID,
	}
}

func verbValidArgMap() map[int][]int {
	return map[int][]int{
		V_ADD:       {K_DUEDATE, K_GROUP, K_LONG, K_PRIORITY, K_SHORT, K_ID, A_SUBTASK},
		V_COMPLETE:  {K_ID, K_COMMENT},
		V_CONFIGURE: {A_LOCAL, A_RESET},
		V_COUNT:     {A_ALL, A_CLOSED, A_DUE, A_OPEN, A_OVERDUE},
		V_DELETE:    {K_ID},
		V_HELP:      {},
		V_LIST:      {A_ALL, A_CLOSED, A_DUE, A_OPEN, A_OVERDUE},
		V_REOPEN:    {K_ID},
		V_SHOW:      {K_ID},
		V_UPDATE:    {K_DUEDATE, K_GROUP, K_LONG, K_PRIORITY, K_SHORT, K_ID, A_SUBTASK},
		V_VERSION:   {},
	}
}

func kwargValidatorMap() map[int]func(string) (interface{}, error) {
	return map[int]func(string) (interface{}, error){
		K_COMMENT: validateString,
		K_DUEDATE: validateDate,
		K_GROUP:   validateGroup,
		K_ID:      validateInt,
		K_LONG:    validateString,
		K_SHORT:   validateShort,
	}
}
