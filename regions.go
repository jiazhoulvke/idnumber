package idnumber

import (
	"regexp"
	"strings"

	"golang.org/x/xerrors"
)

func init() {
	var err error
	if regions, err = parseRegions(regionsData); err != nil {
		panic(err)
	}
}

var regions []Region

//Region 地区信息
type Region struct {
	//Code 地区行政代码
	Code string `json:"code"`
	//Name 地区名
	Name string `json:"name"`
	//ParentRegions 上级地区
	ParentRegions []Region `json:"parent_regions,omitempty"`
}

//Regions 地区列表
func Regions() []Region {
	return regions
}

var linePattern = regexp.MustCompile(`^(\d+)\s+(.+)$`)

func parseRegions(s string) ([]Region, error) {
	s = strings.TrimSpace(s)
	regions := make([]Region, 0, strings.Count(s, "\n")+1)
	for _, line := range strings.Split(s, "\n") {
		line = strings.TrimSpace(line)
		result := linePattern.FindStringSubmatch(line)
		// 返回的数据应该是 result[1] 为地区行政代码，result[2] 为地区名
		if len(result) != 3 || len(result[1]) != 6 {
			return nil, xerrors.Errorf("invalid region: %s", line)
		}
		region := Region{
			Code:          result[1],
			Name:          result[2],
			ParentRegions: make([]Region, 0),
		}
		// 如果后四位不是 0000,则说明有上级地区
		// 三个特例除外：
		// 710000	台湾省
		// 810000	香港特别行政区
		// 820000	澳门特别行政区
		if region.Code[2:] != "0000" {
			level1Code := region.Code[0:2] + "0000"
			level2Code := region.Code[0:4] + "00"
			for _, r := range regions {
				if r.Code == level1Code && region.Code != r.Code {
					region.ParentRegions = append(region.ParentRegions, r)
				}
				if r.Code == level2Code && region.Code != r.Code {
					region.ParentRegions = append(region.ParentRegions, r)
				}
			}
		}
		if len(region.ParentRegions) == 0 {
			region.ParentRegions = nil
		}
		regions = append(regions, region)
	}
	return regions, nil
}
