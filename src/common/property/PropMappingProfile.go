package property

import (
	"github.com/vidorg/vid_backend/src/model/dto"
	"github.com/vidorg/vid_backend/src/model/po"
)

type PropMappingProfile struct {
	Mappings []*PropMapping
}

func CreatePropMappingProfile() *PropMappingProfile {
	mappings := make([]*PropMapping, 0)

	mappings = append(mappings, NewPropMapping(&dto.UserDto{}, &po.User{}, map[string]*PropMappingValue{
		"Uid":          NewPropMappingValue([]string{"Uid"}, false),
		"Username":     NewPropMappingValue([]string{"Username"}, false),
		"Sex":          NewPropMappingValue([]string{"Sex"}, false),
		"BirthTime":    NewPropMappingValue([]string{"BirthTime"}, false),
		"RegisterTime": NewPropMappingValue([]string{"CreateAt"}, false),
	}))

	mappings = append(mappings, NewPropMapping(&dto.UserDto{}, &po.User{}, map[string]*PropMappingValue{
		"Vid":         NewPropMappingValue([]string{"Vid"}, false),
		"Title":       NewPropMappingValue([]string{"Title"}, false),
		"Description": NewPropMappingValue([]string{"Description"}, false),
		"BirthTime":   NewPropMappingValue([]string{"BirthTime"}, false),
		"UploadTime":  NewPropMappingValue([]string{"CreateAt"}, false),
		"UpdateTime":  NewPropMappingValue([]string{"UpdateAt"}, false),
		"Uid":         NewPropMappingValue([]string{"AuthorUid"}, false),
	}))

	return &PropMappingProfile{
		Mappings: mappings,
	}
}

func (p *PropMappingProfile) GetPropertyMapping(dtoModel interface{}, poModel interface{}) *PropMapping {
	for _, m := range p.Mappings {
		if m.DtoModel == dtoModel && m.PoModel == poModel {
			return m
		}
	}
	return nil
}
