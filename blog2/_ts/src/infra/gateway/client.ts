import { Article, ArticleId, setupArticle } from "../../entity/article";
import { File as ArticleFile, FileAndThumbnail, FileThumbnail } from "../../entity/file";
import { Tag, TagId } from "../../entity/tag";
import { APIClient } from "../../gateway/client";


class APIClientImpl implements APIClient {
    async putAdminArticle(
        articleId: ArticleId,
        title: string | undefined,
        published: boolean | undefined,
    ): Promise<Article> {
        const article = await (await fetch(
            `/api/admin/articles/${articleId}`,
            {
                method: "PUT",
                body: JSON.stringify({
                    title,
                    published,
                }),
                headers: {
                    "Content-Type": "application/json",
                },
            },
        )).json() as Article;
        setupArticle(article);
        return article;
    }
    async putAdminArticleMarkdown(
        articleId: ArticleId,
        markdown: string,
    ): Promise<{ html: string, markdown: string }> {
        return await (await fetch(
            `/api/admin/articles/${articleId}/markdown`,
            {
                method: "PUT",
                body: markdown,
                headers: {
                    "Content-Type": "text/plain",
                },
            },
        )).json();
    }
    async postAdminArticleTagEditTags(
        articleId: ArticleId,
        addTagId: TagId[],
        deleteTagId: TagId[],
    ): Promise<{ article: Article, notAttachedTags: Tag[] }> {
        const r = await (await fetch(
            `/api/admin/articles/${articleId}/edit-tags`,
            {
                method: "POST",
                body: JSON.stringify({
                    add: addTagId,
                    delete: deleteTagId,
                }),
                headers: {
                    "Content-Type": "application/json",
                },
            },
        )).json() as { article: Article, notAttachedTags: Tag[] };
        setupArticle(r.article);
        return r;
    }
    async postAdminFiles(
        file: File,
    ): Promise<{ file: ArticleFile, fileThumbnail: FileThumbnail | undefined }> {
        const formData = new FormData();
        formData.append("file", file);
        const r = await (await fetch(
            `/api/admin/files`,
            {
                method: "POST",
                body: formData,
                headers: {
                    "X-File-Type": file.type,
                }
            },
        )).json();
        return r;
    }
    async getAdminFiles(
        q: string,
        page: number,
        size: number,
    ): Promise<{ page: number, size: number, files: FileAndThumbnail[] }> {
        const r = await (await fetch(
            `/api/admin/files?q=${q}&page=${page}&size=${size}`,
            {
                method: "GET",
            },
        )).json();
        return r;
    }
}

export function newAPIClient(): APIClient {
    return new APIClientImpl();
}