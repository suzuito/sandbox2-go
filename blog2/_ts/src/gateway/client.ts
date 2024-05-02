import { Article, ArticleId } from "../entity/article";
import { File as ArticleFile, FileAndThumbnail, FileThumbnail } from "../entity/file";
import { Tag, TagId } from "../entity/tag";

export interface APIClient {
    putAdminArticle(
        articleId: ArticleId,
        title: string | undefined,
        published: boolean | undefined,
    ): Promise<Article>;
    putAdminArticleMarkdown(
        articleId: ArticleId,
        markdown: string,
    ): Promise<{ html: string, markdown: string }>;
    postAdminArticleTagEditTags(
        articleId: ArticleId,
        addTagId: TagId[],
        deleteTagId: TagId[],
    ): Promise<{ article: Article, notAttachedTags: Tag[] }>;
    postAdminFiles(
        file: File,
    ): Promise<{ file: ArticleFile, fileThumbnail: FileThumbnail | undefined }>;
    getAdminFiles(
        q: string,
        page: number,
        size: number,
    ): Promise<{ page: number, size: number, files: FileAndThumbnail[] }>
}