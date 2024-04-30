import { Tag } from "./tag";

export type ArticleId = string;
export interface Article {
    id: ArticleId,
    title: string,
    published: boolean,
    tags: Tag[],
}