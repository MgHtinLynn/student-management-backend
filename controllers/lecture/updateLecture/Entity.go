package updateLecture

type InputUpdateLecture struct {
	ID      int    `json:"id" binding:"required"`
	Name    string `json:"name" binding:"required"`
	TutorId int    `json:"tutor_id" binding:"required"`
}
