import { Tag } from "./tag";

export type ArticleId = string;
export interface Article {
    id: ArticleId,
    title: string,
    published: boolean,
    publishedAt: string,
    publishedAtAsDate: Date | undefined,
    tags: Tag[],
}

export function setupArticle(a: Article) {
    try {
        const d = new Date(Date.parse(a.publishedAt));
        a.publishedAtAsDate = d;
    } catch (e: unknown) {
        // Do nothing
    }
}