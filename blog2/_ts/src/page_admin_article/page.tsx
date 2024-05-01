import { useState } from "react";
import { Article } from "../entity/article";
import PageMenu from "./page_menu";
import { Tag, TagId } from "../entity/tag";
import { APIClient } from "../gateway/client";
import PageEditor, { extractCurrentMarkdown } from "./page_editor";

export default function Page({
    defaultArticle,
    defaultSelectableTags,
    defaultMarkdown,
    defaultHTML,
    apiClient,
}: {
    defaultArticle: Article,
    defaultSelectableTags: Tag[],
    defaultMarkdown: string,
    defaultHTML: string,
    apiClient: APIClient,
}) {
    const [article, setArticle] = useState(defaultArticle);
    const [selectableTags, setSelectableTags] = useState(defaultSelectableTags);
    const [html, setHtml] = useState(defaultHTML);
    const [markdown, setMarkdown] = useState(defaultMarkdown);
    const onClickPublishToggle = async () => {
        let newPublished = true;
        if (article.published) {
            newPublished = false;
        }
        setArticle(
            await apiClient.putAdminArticle(article.id, undefined, newPublished)
        );
    };
    const onClickAddTag = async (tagId: TagId) => {
        const r = await apiClient.postAdminArticleTagEditTags(
            article.id,
            [tagId],
            [],
        );
        setArticle(r.article);
        setSelectableTags(r.notAttachedTags);
    };
    const onClickDeleteTag = async (tagId: TagId) => {
        const r = await apiClient.postAdminArticleTagEditTags(
            article.id,
            [],
            [tagId],
        );
        setArticle(r.article);
        setSelectableTags(r.notAttachedTags);
    };
    const onClickSaveTitle = async (title: string) => {
        setArticle(
            await apiClient.putAdminArticle(article.id, title, undefined),
        );
    }
    const onClickSaveMarkdown = async () => {
        const res = await apiClient.putAdminArticleMarkdown(
            article.id,
            extractCurrentMarkdown(),
        );
        setHtml(res.html);
        setMarkdown(res.markdown);
    }
    return (
        <>
            <PageMenu
                article={article}
                selectableTags={selectableTags}
                onClickPublishToggle={onClickPublishToggle}
                onClickAddTag={onClickAddTag}
                onClickDeleteTag={onClickDeleteTag}
                onClickSaveTitle={onClickSaveTitle}
                onClickSaveMarkdown={onClickSaveMarkdown}
            ></PageMenu>
            <div style={{
                position: "absolute",
                top: 156, bottom: 6,
                left: 6, right: 6,
                overflow: "scroll",
                display: "flex",
                flexDirection: "row",
                border: "1px solid green",
            }}>
                <PageEditor markdown={markdown}></PageEditor>
                <div style={{
                    position: "absolute",
                    top: 0, bottom: 0,
                    width: "49%",
                    right: 0,
                    overflow: "scroll",
                    border: "3px solid blueviolet"
                }} dangerouslySetInnerHTML={{ __html: html }}>
                </div>
            </div>
        </>
    )
}