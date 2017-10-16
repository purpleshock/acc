package admin

import (
	"github.com/purpleshock/acc/admin/repository"
)

func (s *AdminService) Register(mail string, password string) (*repository.Admin, error) {
	_, err := s.adminRepository.FindByMail(mail)
	if !IsAdminNotExist(err) {
		return nil, err
	}

	encryptedPwd := s.encrypt.Digest(password)
	return s.adminRepository.Create(mail, encryptedPwd)
}
