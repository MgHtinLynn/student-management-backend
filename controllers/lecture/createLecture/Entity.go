package createLecture

type InputCreateLecture struct {
	Name    string `json:"name" binding:"required"`
	TutorId int    `json:"tutor_id" binding:"required"`
}
