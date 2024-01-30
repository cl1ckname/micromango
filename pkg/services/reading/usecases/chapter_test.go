package usecases

import (
	commonerr "micromango/pkg/common/errors"
	"micromango/pkg/common/utils"
	"micromango/pkg/services/reading/entity"
	"reflect"
	"testing"
)

var testChapters = map[string]entity.Chapter{
	"1": {
		ChapterId: "1",
		MangaId:   "1",
		Number:    1,
		Title:     "title",
		Pages:     []entity.Page{{PageId: "1", MangaId: "1", ChapterId: "1", Number: 1}},
	},
	"2": {
		ChapterId: "2",
		MangaId:   "1",
		Number:    2,
		Title:     "title2",
		Pages:     []entity.Page{{PageId: "2", MangaId: "1", ChapterId: "1", Number: 2}},
	},
	"3": {
		ChapterId: "3",
		MangaId:   "1",
		Number:    3,
		Title:     "title3",
		Pages:     []entity.Page{{PageId: "3", MangaId: "1", ChapterId: "1", Number: 3}},
	},
}

// Mock for chapter repository
type ChapterRepositoryMock struct{}

func (c *ChapterRepositoryMock) GetContent(_ string) ([]entity.Chapter, error) {
	var chapters []entity.Chapter
	for _, chapter := range []string{"1", "2", "3"} {
		chapters = append(chapters, testChapters[chapter])
	}
	return chapters, nil
}

func (c *ChapterRepositoryMock) GetChapter(chapterId string) (entity.Chapter, error) {
	chapter, ok := testChapters[chapterId]
	if !ok {
		return entity.Chapter{}, &commonerr.ErrNotFound{}
	}
	return chapter, nil
}

func (c *ChapterRepositoryMock) SaveChapter(chapter entity.Chapter) (entity.Chapter, error) {
	chapter.ChapterId = "1"
	return chapter, nil
}

// Mock for activity service
type ActivityServiceMock struct{}

func (a *ActivityServiceMock) GetReadChapters(_ string, _ string) ([]string, error) {
	return []string{"1", "2"}, nil
}

func TestChapter_AddChapter(t *testing.T) {
	type fields struct {
		Repository ChapterRepository
		Activity   ActivityService
	}
	type args struct {
		chapter entity.Chapter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Chapter
		wantErr bool
	}{
		// Write tests here
		{
			name: "Add chapter",
			fields: fields{
				Repository: &ChapterRepositoryMock{},
				Activity:   &ActivityServiceMock{},
			},
			args:    args{chapter: entity.Chapter{MangaId: "1", Number: 3, Title: "Chapter 1"}},
			want:    entity.Chapter{MangaId: "1", ChapterId: "1", Number: 3, Title: "Chapter 1"},
			wantErr: false,
		},
		{
			name: "Add chapter 2",
			fields: fields{
				Repository: &ChapterRepositoryMock{},
				Activity:   &ActivityServiceMock{},
			},
			args:    args{chapter: entity.Chapter{MangaId: "1", Number: 2, Title: "Chapter 2"}},
			want:    entity.Chapter{MangaId: "1", ChapterId: "1", Number: 2, Title: "Chapter 2"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chapter{
				Repository: tt.fields.Repository,
				Activity:   tt.fields.Activity,
			}
			got, err := c.AddChapter(tt.args.chapter)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddChapter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddChapter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChapter_GetChapter(t *testing.T) {
	type fields struct {
		Repository ChapterRepository
		Activity   ActivityService
	}
	type args struct {
		chapterId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Chapter
		wantErr bool
	}{
		{
			name: "Get chapter",
			fields: fields{
				Repository: &ChapterRepositoryMock{},
				Activity:   &ActivityServiceMock{},
			},
			args: args{chapterId: "1"},
			want: entity.Chapter{MangaId: "1", ChapterId: "1", Number: 1, Title: "title", Pages: []entity.Page{{
				PageId:    "1",
				MangaId:   "1",
				ChapterId: "1",
				Number:    1,
			}}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chapter{
				Repository: tt.fields.Repository,
				Activity:   tt.fields.Activity,
			}
			got, err := c.GetChapter(tt.args.chapterId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetChapter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetChapter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChapter_GetMangaContent(t *testing.T) {
	type fields struct {
		Repository ChapterRepository
		Activity   ActivityService
	}
	type args struct {
		mangaId string
		userId  *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.ChapterHead
		wantErr bool
	}{
		{
			name: "Get manga content with user",
			fields: fields{
				Repository: &ChapterRepositoryMock{},
				Activity:   &ActivityServiceMock{},
			},
			args: args{mangaId: "1", userId: utils.Ptr("user")},
			want: []entity.ChapterHead{
				{ChapterId: "1", Number: 1, Title: "title", Pages: uint32(1), Read: true},
				{ChapterId: "2", Number: 2, Title: "title2", Pages: uint32(1), Read: true},
				{ChapterId: "3", Number: 3, Title: "title3", Pages: uint32(1), Read: false},
			},
			wantErr: false,
		},
		{
			name: "Get manga content without user",
			fields: fields{
				Repository: &ChapterRepositoryMock{},
				Activity:   &ActivityServiceMock{},
			},
			args: args{mangaId: "1", userId: nil},
			want: []entity.ChapterHead{
				{ChapterId: "1", Number: 1, Title: "title", Pages: uint32(1), Read: false},
				{ChapterId: "2", Number: 2, Title: "title2", Pages: uint32(1), Read: false},
				{ChapterId: "3", Number: 3, Title: "title3", Pages: uint32(1), Read: false},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chapter{
				Repository: tt.fields.Repository,
				Activity:   tt.fields.Activity,
			}
			got, err := c.GetMangaContent(tt.args.mangaId, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMangaContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMangaContent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChapter_UpdateChapter(t *testing.T) {
	type fields struct {
		Repository ChapterRepository
		Activity   ActivityService
	}
	type args struct {
		chapterId string
		data      entity.UpdateChapterDto
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Chapter
		wantErr bool
	}{
		{
			name: "Update chapter",
			fields: fields{
				Repository: &ChapterRepositoryMock{},
				Activity:   &ActivityServiceMock{},
			},
			args: args{chapterId: "1", data: entity.UpdateChapterDto{Number: utils.Ptr[float32](2), Title: utils.Ptr("Chapter 1")}},
			want: entity.Chapter{MangaId: "1", ChapterId: "1", Number: 2, Title: "Chapter 1", Pages: []entity.Page{{
				PageId:    "1",
				MangaId:   "1",
				ChapterId: "1",
				Number:    1,
			}}},
			wantErr: false,
		},
		{
			name: "Not found for update chapter",
			fields: fields{
				Repository: &ChapterRepositoryMock{},
				Activity:   &ActivityServiceMock{},
			},
			args:    args{chapterId: "5", data: entity.UpdateChapterDto{}},
			want:    entity.Chapter{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chapter{
				Repository: tt.fields.Repository,
				Activity:   tt.fields.Activity,
			}
			got, err := c.UpdateChapter(tt.args.chapterId, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateChapter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateChapter() got = %v, want %v", got, tt.want)
			}
		})
	}
}
