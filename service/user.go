package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hfs1988/jagaad_test/adapter"
	"github.com/hfs1988/jagaad_test/common"
	"github.com/hfs1988/jagaad_test/config"
	"github.com/hfs1988/jagaad_test/entity"
)

type userService struct {
	http adapter.HTTPAdapter
	csv  adapter.CSVAdapter
	conf config.Config
}

func GetUserService(http adapter.HTTPAdapter, csv adapter.CSVAdapter, conf config.Config) *userService {
	return &userService{
		http: http,
		csv:  csv,
		conf: conf,
	}
}

func (s *userService) FetchUsers() ([]entity.User, error) {
	var users []entity.User
	err := s.http.Get(s.conf.ThirdParty.URL, &users)
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *userService) WriteUsers(users []entity.User) error {
	var (
		dataUsers       = make([][]string, len(users))
		dataUserTags    = make([][]string, len(users))
		dataUserFriends = make([][]string, len(users))
	)

	for k, v := range users {
		dataUsers[k] = append(dataUsers[k], v.ID, fmt.Sprintf("%d", v.Index), v.GUID, fmt.Sprintf("%v", v.IsActive), v.Balance, fmt.Sprintf("%d", k), fmt.Sprintf("%d", k))
		dataUserTags[k] = append(dataUserTags[k], v.Tags...)

		var friends []string
		for _, friend := range v.Friends {
			friends = append(friends, fmt.Sprintf("%d:%s", friend.ID, friend.Name))
		}
		dataUserFriends[k] = append(dataUserFriends[k], friends...)
	}

	err := s.csv.Write(s.conf.Data.UserFilename, dataUsers)
	if err != nil {
		return err
	}

	err = s.csv.Write(s.conf.Data.UserTagsFilename, dataUserTags)
	if err != nil {
		return err
	}

	err = s.csv.Write(s.conf.Data.UserFriendsFilename, dataUserFriends)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) GetUsers(tags []string) ([]entity.User, error) {
	var users []entity.User

	dataUsers, err := s.csv.Read(s.conf.Data.UserFilename)
	if err != nil {
		return users, err
	}

	dataUserTags, err := s.csv.Read(s.conf.Data.UserTagsFilename)
	if err != nil {
		return users, err
	}

	dataUserFriends, err := s.csv.Read(s.conf.Data.UserFriendsFilename)
	if err != nil {
		return users, err
	}

	for _, v := range dataUsers {
		isUserExist := true
		tagsIndex, err := strconv.Atoi(v[5])
		if err != nil {
			return users, err
		}

		userTagsMap := common.ConvertToMap(dataUserTags[tagsIndex])
		for i := 0; i < len(tags); i++ {
			if !common.IsExist(userTagsMap, tags[i]) {
				isUserExist = false
			}
		}
		if !isUserExist {
			continue
		}

		user := entity.User{}
		user.ID = v[0]
		index, err := strconv.Atoi(v[1])
		if err != nil {
			return users, err
		}
		user.Index = index
		user.GUID = v[2]
		isActive, err := strconv.ParseBool(v[3])
		if err != nil {
			return users, err
		}
		user.IsActive = isActive
		user.Balance = v[4]
		user.Tags = dataUserTags[tagsIndex]

		friendsIndex, err := strconv.Atoi(v[6])
		if err != nil {
			return users, err
		}
		for _, raw := range dataUserFriends[friendsIndex] {
			splittedRaw := strings.Split(raw, ":")
			id, err := strconv.Atoi(splittedRaw[0])
			if err != nil {
				return users, err
			}
			friend := entity.UserFriend{
				ID:   id,
				Name: splittedRaw[1],
			}
			user.Friends = append(user.Friends, friend)
		}

		users = append(users, user)
	}

	return users, nil
}
