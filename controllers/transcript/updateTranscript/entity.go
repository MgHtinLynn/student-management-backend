package updateTranscript

type InputUpdateTranscript struct {
	ID        int    `json:"id" binding:"required"`
	StudentId int    `json:"student_id" binding:"required"`
	LectureID int    `json:"lecture_id" binding:"required"`
	FilePath  string `json:"file_path"`
}
