package service

import (
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/Aoi-hosizora/ahlib/xstatus"
	"github.com/vidorg/vid_backend/lib/xgorm"
	"github.com/vidorg/vid_backend/src/model/param"
	"github.com/vidorg/vid_backend/src/model/po"
	"github.com/vidorg/vid_backend/src/provide/sn"
	"gorm.io/gorm"
)

type VideoService struct {
	db             *gorm.DB
	common         *CommonService
	userService    *UserService
	orderbyService *OrderbyService
}

func NewVideoService() *VideoService {
	return &VideoService{
		db:             xdi.GetByNameForce(sn.SGorm).(*gorm.DB),
		common:         xdi.GetByNameForce(sn.SCommonService).(*CommonService),
		userService:    xdi.GetByNameForce(sn.SUserService).(*UserService),
		orderbyService: xdi.GetByNameForce(sn.SOrderbyService).(*OrderbyService),
	}
}

func (v *VideoService) QueryAll(pp *param.PageOrderParam) ([]*po.Video, int32, error) {
	total := int64(0)
	rdb := v.db.Model(&po.Video{}).Count(&total)
	if rdb.Error != nil {
		return nil, 0, rdb.Error
	}

	videos := make([]*po.Video, 0)
	rdb = xgorm.WithDB(v.db).Pagination(pp.Limit, pp.Page).Model(&po.Video{}).Order(v.orderbyService.Video(pp.Order)).Find(&videos)
	if rdb.Error != nil {
		return nil, 0, rdb.Error
	}

	return videos, int32(total), nil
}

func (v *VideoService) QueryByUid(uid uint64, pp *param.PageOrderParam) ([]*po.Video, int32, error) {
	author, err := v.userService.QueryByUid(uid)
	if err != nil {
		return nil, 0, err
	} else if author == nil {
		return nil, 0, nil
	}

	total := int64(0)
	rdb := v.db.Model(&po.Video{}).Where(&po.Video{AuthorUid: uid}).Count(&total)
	if rdb.Error != nil {
		return nil, 0, rdb.Error
	}

	videos := make([]*po.Video, 0)
	rdb = xgorm.WithDB(v.db).Pagination(pp.Limit, pp.Page).Model(&po.Video{}).Order(v.orderbyService.Video(pp.Order)).Where("author_uid = ?", uid).Find(&videos)
	if rdb.Error != nil {
		return nil, 0, rdb.Error
	}

	return videos, int32(total), nil
}

func (v *VideoService) QueryByVid(vid uint64) (*po.Video, error) {
	video := &po.Video{}
	rdb := v.db.Model(&po.Video{}).Where("vid = ?", vid).First(video)
	if rdb.RowsAffected == 0 {
		return nil, nil
	} else if rdb.Error != nil {
		return nil, rdb.Error
	}

	return video, nil
}

func (v *VideoService) QueryCountByUids(uids []uint64) ([]int32, error) {
	if len(uids) == 0 {
		return []int32{}, nil
	}

	counts := make([]*_IdCntScanResult, 0)
	where := v.common.BuildOrExpr("author_uid", uids)
	rdb := v.db.Model(&po.Video{}).Select("author_uid as id, count(*) as cnt").Where(where).Group("author_uid").Scan(&counts)
	if rdb.Error != nil {
		return nil, rdb.Error
	}

	bucket := make(map[uint64]int32)
	for _, r := range counts {
		bucket[r.Id] = r.Cnt
	}
	out := make([]int32, len(uids))
	for idx, uid := range uids {
		if cnt, ok := bucket[uid]; ok {
			out[idx] = cnt
		}
	}
	return out, nil
}

func (v *VideoService) Existed(vid uint64) (bool, error) {
	cnt := int64(0)
	rdb := v.db.Model(&po.Video{}).Where("vid = ?", vid).Count(&cnt)
	if rdb.Error != nil {
		return false, rdb.Error
	}

	return cnt > 0, nil
}

func (v *VideoService) Insert(pa *param.InsertVideoParam, uid uint64) (xstatus.DbStatus, error) {
	video := pa.ToVideoPo()
	video.AuthorUid = uid

	rdb := v.db.Model(&po.Video{}).Create(video)
	return xgorm.CreateErr(rdb)
}

func (v *VideoService) Update(vid uint64, video *param.UpdateVideoParam) (xstatus.DbStatus, error) {
	rdb := v.db.Model(&po.Video{}).Where("vid = ?", vid).Updates(video.ToMap())
	return xgorm.UpdateErr(rdb)
}

func (v *VideoService) Delete(vid uint64) (xstatus.DbStatus, error) {
	rdb := v.db.Model(&po.Video{}).Where("vid = ?", vid).Delete(&po.Video{})
	return xgorm.DeleteErr(rdb)
}
