import { Article, ArticleId } from "../../entity/article";
import { Tag, TagId } from "../../entity/tag";
import { APIClient } from "../../gateway/client";


class APIClientImpl implements APIClient {
    async putAdminArticle(
        articleId: ArticleId,
        title: string | undefined,
        published: boolean | undefined,
    ): Promise<Article> {
        return await (await fetch(
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
        )).json();
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
        return await (await fetch(
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
        )).json();
    }
}

export function newAPIClient(): APIClient {
    return new APIClientImpl();
}