package profile

import (
	"github.com/Aoi-hosizora/ahlib-web/xtime"
	"github.com/Aoi-hosizora/ahlib/xcondition"
	"github.com/Aoi-hosizora/ahlib/xentity"
	"github.com/vidorg/vid_backend/src/model/constant"
	"github.com/vidorg/vid_backend/src/model/dto"
	"github.com/vidorg/vid_backend/src/model/param"
	"github.com/vidorg/vid_backend/src/model/po"
)

func addDtoMappers() {
	// userPo -> userDto
	xentity.AddMapper(xentity.NewMapper(&po.User{}, func() interface{} { return &dto.UserDto{} }, func(from interface{}, to interface{}) error {
		user := from.(*po.User)
		userDto := to.(*dto.UserDto)

		userDto.Uid = user.Uid
		userDto.Username = user.Username
		userDto.Gender = user.Gender.String()
		userDto.Profile = user.Profile
		userDto.AvatarUrl = user.AvatarUrl // TODO
		userDto.Birthday = user.Birthday.String()
		userDto.Role = user.Role
		userDto.RegisterTime = xtime.NewJsonDateTime(user.CreatedAt).String()
		return nil
	}))

	// videoPo -> videoDto
	xentity.AddMapper(xentity.NewMapper(&po.Video{}, func() interface{} { return &dto.VideoDto{} }, func(from interface{}, to interface{}) error {
		video := from.(*po.Video)
		videoDto := to.(*dto.VideoDto)

		videoDto.Vid = video.Vid
		videoDto.Title = video.Title
		videoDto.Description = video.Description
		videoDto.CoverUrl = video.CoverUrl // TODO
		videoDto.VideoUrl = video.VideoUrl
		videoDto.UploadTime = xtime.NewJsonDateTime(video.CreatedAt).String()
		videoDto.UpdateTime = xtime.NewJsonDateTime(video.UpdatedAt).String()
		videoDto.Author = xentity.MustMap(video.Author, &dto.UserDto{}).(*dto.UserDto)
		return nil
	}))

	// policyPo -> policyDto
	xentity.AddMapper(xentity.NewMapper(&po.Policy{}, func() interface{} { return &dto.PolicyDto{} }, func(from interface{}, to interface{}) error {
		policy := from.(*po.Policy)
		policyDto := to.(*dto.PolicyDto)

		policyDto.Role = policy.V0
		policyDto.Path = policy.V1
		policyDto.Method = policy.V1
		return nil
	}))
}

func addParamMappers() {
	// userParam -> userPo
	xentity.AddMapper(xentity.NewMapper(&param.UserParam{}, func() interface{} { return &po.User{} }, func(from interface{}, to interface{}) error {
		userParam := from.(*param.UserParam)
		user := to.(*po.User)

		user.Username = userParam.Username
		user.Profile = *userParam.Profile
		user.Gender = constant.ParseSexEnum(userParam.Gender)
		user.Birthday = xcondition.First(xtime.ParseRFC3339Date(userParam.Birthday)).(xtime.JsonDate)
		user.PhoneNumber = userParam.PhoneNumber
		user.AvatarUrl = userParam.AvatarUrl // TODO
		return nil
	}))

	// videoParam -> videoPo
	xentity.AddMapper(xentity.NewMapper(&param.VideoParam{}, func() interface{} { return &po.Video{} }, func(from interface{}, to interface{}) error {
		videoParam := from.(*param.VideoParam)
		video := to.(*po.Video)

		video.Title = videoParam.Title
		video.Description = *videoParam.Description
		video.CoverUrl = videoParam.CoverUrl // TODO
		video.VideoUrl = videoParam.VideoUrl // TODO
		return nil
	}))
}
