package admin

import (
	"github.com/purpleshock/acc/admin/repository"
)

func (s *AdminService) Login(mail string, password string) (*repository.Admin, error) {
	admin, err := s.adminRepository.FindByMail(mail)
	if IsAdminNotExist(err) {
		return nil, err
	}

	if admin.Hashed != s.encrypt.Digest(password) {
		return nil, NewInvalidPassword()
	}

	return admin, nil
}
