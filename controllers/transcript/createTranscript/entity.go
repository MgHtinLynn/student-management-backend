package createTranscript

type InputCreateTranscript struct {
	StudentId int    `json:"student_id" binding:"required"`
	LectureID int    `json:"lecture_id" binding:"required"`
	FilePath  string `json:"file_path"`
}
