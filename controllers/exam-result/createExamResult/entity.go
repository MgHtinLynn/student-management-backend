package createExamResult

type InputCreateExamResult struct {
	Status    string `json:"status" binding:"required"`
	StudentId int    `json:"student_id" binding:"required"`
	SubjectID int    `json:"subject_id" binding:"required"`
	FilePath  string `json:"file_path"`
}
