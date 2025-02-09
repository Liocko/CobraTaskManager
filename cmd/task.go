package cmd

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

var TaskList = []Task{
	//{"TestTask", "Testing adding tasks", "✗"},
	//{"SecondTask", "Another test task", "✓"},
}
