package requests

type UserChangePassword struct {
	Password    string `json:"password" binding:"min=8,required"`
	NewPassword string `json:"newPassword" binding:"min=8,required"`
}
