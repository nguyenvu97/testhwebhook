package Service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm-tutorial/Dao"
	"gorm-tutorial/Model"
	"gorm.io/gorm"
)

type UserService struct {
	userRepository *Dao.UserRepository
}

func NewUserService(db *gorm.DB) *UserService {
	userRepository := Dao.NewUserRepository(db)

	return &UserService{
		userRepository: userRepository,
	}
}

// AddUser thêm một người dùng mới.
func (us *UserService) AddUser(user *Model.User) error {
	return us.userRepository.AddUser(user)
}

// FindByID tìm một người dùng theo ID.
func (us *UserService) FindByID(id uint) (*Model.UserDto, error) {
	result, err := us.userRepository.FindByID(id)
	if err != nil {
		// Xử lý lỗi từ repository nếu có
		return nil, err
	}

	// Kiểm tra nếu không tìm thấy người dùng
	if result == nil {
		return nil, errors.New("User not found")
	}
	userDto := Model.MapUserToDTO(result)
	return userDto, nil
}

// Update cập nhật thông tin của một người dùng.
func (us *UserService) Update(user *Model.User) error {
	return us.userRepository.Update(user)
}

// SearchUser tìm kiếm một người dùng theo tên người dùng.
func (us *UserService) SearchUser(email string) (*Model.User, error) {
	return us.userRepository.SreachUser(email)
}
func (us *UserService) GetAllUser() (*[]Model.User, error) {
	user, err := us.userRepository.GetAllUser()
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (us *UserService) Login(email string, password string) (*Model.User, error) {
	user, err := us.userRepository.FindByEmail(email)
	if err != nil {
		return nil, errors.New("User not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("Invalid password")
	}
	return user, nil
}
