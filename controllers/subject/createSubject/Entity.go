package createSubject

type InputCreateSubject struct {
	Name      string `json:"name" binding:"required"`
	TeacherId int    `json:"teacher_id" binding:"required"`
	LectureId int    `json:"lecture_id" binding:"required"`
}
