package usecases

import (
	"micromango/pkg/common"
	"micromango/pkg/common/errors"
	"micromango/pkg/services/reading/entity"
	"reflect"
	"testing"
)

type PageRepositoryMock struct{}

func (p *PageRepositoryMock) GetPage(pageId string) (entity.Page, error) {
	if pageId == "0" {
		return entity.Page{}, errors.ThrowNotFound("not found")
	}
	return entity.Page{PageId: pageId, Number: 1, ChapterId: "1", MangaId: "1"}, nil
}

func (p *PageRepositoryMock) SavePage(page entity.Page) (entity.Page, error) {
	page.PageId = "0"
	return page, nil
}

type StaticServiceMock struct{}

func (s *StaticServiceMock) UploadPage(_ string, _ string, _ common.File) (string, error) {
	return "path/to/file", nil
}

func TestPage_AddPage(t *testing.T) {
	type fields struct {
		Repository PageRepository
		Static     StaticService
	}
	type args struct {
		dto entity.AddPageDto
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes entity.Page
		wantErr bool
	}{
		{
			name: "Add page without image",
			fields: fields{
				Repository: &PageRepositoryMock{},
				Static:     &StaticServiceMock{},
			},
			args: args{dto: entity.AddPageDto{MangaId: "1", ChapterId: "2", Number: 3}},
			wantRes: entity.Page{
				PageId:    "0",
				MangaId:   "1",
				ChapterId: "2",
				Number:    3,
				Image:     "",
			},
			wantErr: false,
		},
		{
			name: "Add page with image",
			fields: fields{
				Repository: &PageRepositoryMock{},
				Static:     &StaticServiceMock{},
			},
			args: args{dto: entity.AddPageDto{MangaId: "1", ChapterId: "2", Number: 3, Image: &common.File{}}},
			wantRes: entity.Page{
				PageId:    "0",
				MangaId:   "1",
				ChapterId: "2",
				Number:    3,
				Image:     "path/to/file",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Page{
				Repository: tt.fields.Repository,
				Static:     tt.fields.Static,
			}
			gotRes, err := p.AddPage(tt.args.dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("AddPage() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestPage_GetPage(t *testing.T) {
	type fields struct {
		Repository PageRepository
		Static     StaticService
	}
	type args struct {
		pageId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Page
		wantErr bool
	}{
		{
			name: "Get page 1",
			fields: fields{
				Repository: &PageRepositoryMock{},
				Static:     &StaticServiceMock{},
			},
			args:    args{pageId: "1"},
			want:    entity.Page{PageId: "1", Number: 1, ChapterId: "1", MangaId: "1"},
			wantErr: false,
		},
		{
			name: "Get page not found",
			fields: fields{
				Repository: &PageRepositoryMock{},
				Static:     &StaticServiceMock{},
			},
			args:    args{pageId: "0"},
			want:    entity.Page{},
			wantErr: true,
		},
		{
			name: "Get page 2",
			fields: fields{
				Repository: &PageRepositoryMock{},
				Static:     &StaticServiceMock{},
			},
			args:    args{pageId: "2"},
			want:    entity.Page{PageId: "2", Number: 1, ChapterId: "1", MangaId: "1"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Page{
				Repository: tt.fields.Repository,
				Static:     tt.fields.Static,
			}
			got, err := p.GetPage(tt.args.pageId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPage() got = %v, want %v", got, tt.want)
			}
		})
	}
}
