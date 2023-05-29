package service

import (
	"log"

	"github.com/captrep/gin-simple-crud/model/domain/user"
	"github.com/captrep/gin-simple-crud/model/web"
	"github.com/captrep/gin-simple-crud/utils/res"
)

type UserService interface {
	CreateUser(request *web.UserCreateRequest) (*web.UserResponse, *res.Err)
	GetUser() ([]web.UserResponse, *res.Err)
	FindById(*string) (*web.UserResponse, *res.Err)
	UpdateUser(*web.UserUpdateRequest) (*web.UserResponse, *res.Err)
	DeleteUser(*string) *res.Err
}

type UserServiceImpl struct {
	UserRepository user.UserRepository
}

func NewUserService(userRepository user.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (service *UserServiceImpl) CreateUser(request *web.UserCreateRequest) (*web.UserResponse, *res.Err) {
	user := user.User{
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Email:     request.Email,
		Password:  request.Password,
	}

	var err *res.Err
	user, err = service.UserRepository.Save(user)
	if err != nil {
		return nil, err
	}

	return &web.UserResponse{
		Id:        user.Id,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
	}, nil
}

func (service *UserServiceImpl) GetUser() ([]web.UserResponse, *res.Err) {
	var userResponse []web.UserResponse
	result, err := service.UserRepository.GetAll()
	if err != nil {
		return nil, err
	}
	for _, r := range result {
		userResponse = append(userResponse, web.UserResponse{
			Id:        r.Id,
			Firstname: r.Firstname,
			Lastname:  r.Lastname,
			Email:     r.Email,
			CreatedAt: r.CreatedAt,
		})
	}
	return userResponse, nil
}

func (service *UserServiceImpl) FindById(userId *string) (*web.UserResponse, *res.Err) {
	result, err := service.UserRepository.FindById(*userId)
	if err != nil {
		return nil, err
	}
	return &web.UserResponse{
		Id:        result.Id,
		Firstname: result.Firstname,
		Lastname:  result.Lastname,
		Email:     result.Email,
		CreatedAt: result.CreatedAt,
	}, nil
}

func (service *UserServiceImpl) UpdateUser(request *web.UserUpdateRequest) (*web.UserResponse, *res.Err) {
	current, err := service.UserRepository.FindById(request.Id)
	if err != nil {
		log.Println("gagal cari user", err)
		return nil, err
	}
	current.Firstname = request.Firstname
	current.Lastname = request.Lastname
	current.Email = request.Email

	result, errUpdate := service.UserRepository.Update(current)
	if errUpdate != nil {
		return nil, err
	}

	return &web.UserResponse{
		Id:        result.Id,
		Firstname: result.Firstname,
		Lastname:  result.Lastname,
		Email:     result.Email,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (service *UserServiceImpl) DeleteUser(userId *string) *res.Err {
	user, err := service.UserRepository.FindById(*userId)
	if err != nil {
		log.Println("gagal cari user", err)
		return err
	}
	service.UserRepository.Delete(user)
	return nil
}
