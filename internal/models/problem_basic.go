package models

import (
	"gorm.io/gorm"
)

type ProblemBasic struct {
	ID        uint           `gorm:"primarykey;" json:"id"`
	CreatedAt MyTime         `json:"created_at"`
	UpdatedAt MyTime         `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;" json:"deleted_at"`

	Identity          string             `gorm:"column:identity;type:varchar(36);" json:"identity"`                  // 问题表的唯一标识
	ProblemCategories []*ProblemCategory `gorm:"foreignKey:problem_id;references:id" json:"problem_categories"`      // 关联问题分类表
	Title             string             `gorm:"column:title;type:varchar(255);" json:"title"`                       // 文章标题
	Content           string             `gorm:"column:content;type:text;" json:"content"`                           // 文章正文
	MaxRuntime        int                `gorm:"column:max_runtime;type:int(11);" json:"max_runtime"`                // 最大运行时长
	MaxMem            int                `gorm:"column:max_mem;type:int(11);" json:"max_mem"`                        // 最大运行内存
	TestCases         []*TestCase        `gorm:"foreignKey:problem_identity;references:identity;" json:"test_cases"` // 关联测试用例表
	PassNum           int64              `gorm:"column:pass_num;type:int(11);" json:"pass_num"`                      // 通过次数
	SubmitNum         int64              `gorm:"column:submit_num;type:int(11);" json:"submit_num"`                  // 提交次数
	OriginCode        string             `gorm:"column:origin_code;type:varchar(255)" json:"origin_code"`            // 初始代码
}

func (table *ProblemBasic) TableName() string {
	return "problem_basic"
}

func GetProblemList(keyword, categoryIdentity string) *gorm.DB {
	// model（）是去查哪张表
	// .Distinct("problem_basic.id")：这个部分表示查询结果去重，只返回不同的 problem_basic 表中的 id 字段。
	//.Select(...)：这个部分指定了查询需要返回的字段列表。这些字段包括 problem_basic 表中的 id, identity, title, max_runtime, max_mem, pass_num, submit_num, created_at, updated_at, deleted_at 字段。
	// 通过 .Preload() 方法还预加载了相关联的模型数据，包括 ProblemCategories 和 TestCases。
	// .Where("title like ? OR content like ? ", "%"+keyword+"%", "%"+keyword+"%")：这个部分添加了查询条件，限定了只返回标题或内容中包含指定关键词的记录。
	tx := DB.Model(new(ProblemBasic)).Distinct("`problem_basic`.`id`").Select("DISTINCT(`problem_basic`.`id`), `problem_basic`.`identity`, "+
		"`problem_basic`.`title`, `problem_basic`.`max_runtime`, `problem_basic`.`max_mem`, `problem_basic`.`pass_num`, "+
		"`problem_basic`.`submit_num`, `problem_basic`.`created_at`, `problem_basic`.`updated_at`, `problem_basic`.`deleted_at` ").Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").Preload("TestCases").
		Where("title like ? OR content like ? ", "%"+keyword+"%", "%"+keyword+"%")
	if categoryIdentity != "" {
		tx.Joins("RIGHT JOIN problem_category pc on pc.problem_id = problem_basic.id").
			Where("pc.category_id = (SELECT cb.id FROM category_basic cb WHERE cb.identity = ? )", categoryIdentity)
	}
	return tx.Order("problem_basic.id DESC")
}
