package updateExamResult

type InputUpdateExamResult struct {
	ID        int    `json:"id" binding:"required"`
	Status    string `json:"status" binding:"required"`
	StudentId int    `json:"student_id" binding:"required"`
	SubjectID int    `json:"subject_id" binding:"required"`
	FilePath  string `json:"file_path"`
}
