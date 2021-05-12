package model

// Worker 工人
type Worker struct {

	// Name 姓名
	Name string `json:"name"`

	// Gender 性别
	Gender string

	// Age 年龄
	Age int64 `json:"age,omitempty"`

	// Site 网址
	Site string `json:"blog"`

	// Job 工作
	Job string `json:"-"`

	// title 头衔
	title string
}

// SetTitle 设置title
func (w *Worker) SetTitle(title string) {
	w.title = title
}

// GetTitle 获取title
func (w *Worker) GetTitle() string {
	return w.title
}
