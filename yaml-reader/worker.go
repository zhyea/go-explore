package main

//Workshop 作坊
type Workshop struct {
	W1 Worker `mapstructure:"worker1"`
	W2 Worker `mapstructure:"worker2"`
}

//Worker 工人
type Worker struct {
	Name string
	Age  int64
}
