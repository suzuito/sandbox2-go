package bmysql

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/suzuito/sandbox2-go/blog/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func newDB() (*gorm.DB, error) {
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Duration(0), // Slow SQL threshold
			LogLevel:                  logger.Info,      // Log level
			IgnoreRecordNotFoundError: true,             // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,            // Disable color
		},
	)
	db, err := gorm.Open(
		mysql.Open("root:@tcp(127.0.0.1:3306)/blog_for_unit_test?charset=utf8mb4&parseTime=True"),
		&gorm.Config{
			Logger: dbLogger,
		},
	)
	return db, err
}

func cleanUpDB(t *testing.T, db *gorm.DB) {
	mustExecSQL(t, db, "DELETE FROM articles_search_index")
	mustExecSQL(t, db, "DELETE FROM mapping_articles_source_articles")
	mustExecSQL(t, db, "DELETE FROM mapping_articles_tags")
	mustExecSQL(t, db, "DELETE FROM tags")
	mustExecSQL(t, db, "DELETE FROM articles")
}

func timeDate(year int, month time.Month, day int, hour int, min int, sec int, nsec int, loc *time.Location) *time.Time {
	a := time.Date(year, month, day, hour, min, sec, nsec, loc)
	return &a
}

func mustExecSQL(t *testing.T, db *gorm.DB, sqlStr string) {
	if err := db.Exec(sqlStr).Error; err != nil {
		t.Errorf("SQL is failed : %+v\n", err)
	}
}

func mustCreate(t *testing.T, db *gorm.DB, v any) {
	if err := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(v).Error; err != nil {
		t.Errorf("Create is failed : %+v\n", err)
	}
}

func setUpTestData(t *testing.T, db *gorm.DB) {
	dummyDateTime := time.Date(2023, 1, 2, 3, 4, 5, 0, time.UTC)

	mustCreate(t, db, tableArticle{
		ID:                            "article01",
		Version:                       1,
		Title:                         "articleTitle01",
		Description:                   "articleDescription01",
		Date:                          time.Date(2023, 1, 1, 1, 1, 1, 1, time.UTC),
		Tags:                          []tableTag{},
		MappingArticlesSourceArticles: tableMappingArticlesSourceArticles{},
		PublishedAt:                   &dummyDateTime,
		Model: gorm.Model{
			CreatedAt: dummyDateTime,
			UpdatedAt: dummyDateTime,
		},
	})
	mustCreate(t, db, tableArticle{
		ID:                            "article01",
		Version:                       2,
		Title:                         "articleTitle01",
		Description:                   "articleDescription01",
		Date:                          time.Date(2023, 1, 2, 1, 1, 1, 1, time.UTC),
		Tags:                          []tableTag{},
		MappingArticlesSourceArticles: tableMappingArticlesSourceArticles{},
		PublishedAt:                   &dummyDateTime,
		Model: gorm.Model{
			CreatedAt: dummyDateTime,
			UpdatedAt: dummyDateTime,
		},
	})
	mustCreate(t, db, tableArticle{
		ID:                            "article02",
		Version:                       1,
		Title:                         "articleTitle02",
		Description:                   "articleDescription02",
		Date:                          time.Date(2023, 1, 3, 1, 1, 1, 1, time.UTC),
		Tags:                          []tableTag{},
		MappingArticlesSourceArticles: tableMappingArticlesSourceArticles{},
		PublishedAt:                   &dummyDateTime,
		Model: gorm.Model{
			CreatedAt: dummyDateTime,
			UpdatedAt: dummyDateTime,
		},
	})
	mustCreate(t, db, tableTag{ID: "tag01"})
	mustCreate(t, db, tableTag{ID: "tag02"})
	mustCreate(t, db, tableMappingArticlesTags{
		ArticleID: "article01", ArticleVersion: 1, TagID: "tag01",
	})
	mustCreate(t, db, tableMappingArticlesTags{
		ArticleID: "article01", ArticleVersion: 2, TagID: "tag02",
	})
	mustCreate(t, db, tableMappingArticlesSourceArticles{
		ArticleID: "article01", ArticleVersion: 1,
		ArticleSourceID: "src01", ArticleSourceVersion: "sha01",
		Meta: entity.ArticleSourceMeta{
			URL: "https://www.example.com",
		},
	})
	mustCreate(t, db, tableMappingArticlesSourceArticles{
		ArticleID: "article01", ArticleVersion: 2,
		ArticleSourceID: "src01", ArticleSourceVersion: "sha02",
		Meta: entity.ArticleSourceMeta{
			URL: "https://www.example.com",
		},
	})
}
