package v1

import metav1 "github.com/Pis0sion/rblogrus/meta/types"

type Article struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`

	AuthorID           uint64 `json:"authorID"     gorm:"column:authorID" validate:"omitempty"`
	ArticleTitle       string `json:"articleTitle" gorm:"column:articleTitle" validate:"omitempty"`
	ArticleContent     string `json:"articleContent"  gorm:"column:articleContent" validate:"omitempty"`
	ArticleDescription string `json:"articleDescription"  gorm:"column:articleDescription" validate:"omitempty"`
	Status             int    `json:"Status"    gorm:"column:Status" validate:"omitempty"`
}

type ArticleList struct {

	metav1.ListMeta `json:",inline"`

	Items []*Article `json:"items"`
}

func (a *Article) TableName() string {
	return "article"
}
