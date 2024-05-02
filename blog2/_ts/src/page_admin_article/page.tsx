import { useState } from "react";
import { Article } from "../entity/article";
import PageMenu from "./page_menu";
import { Tag, TagId } from "../entity/tag";
import { APIClient } from "../gateway/client";
import PageEditor, { extractCurrentMarkdown } from "./page_editor";
import DialogFile from "../dialog_file/main";
import { File as ArticleFile, FileAndThumbnail, generateFileMarkdownURL } from "../entity/file";

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
    const [isOpenDialogFile, setIsOpenDialogFile] = useState(false);
    const [filesInDialogFile, setFilesInDialogFile] = useState<FileAndThumbnail[]>([]);
    const [isClickablePrevInDialogFile, setIsClickablePrevInDialogFile] = useState(false);
    const [isClickableNextInDialogFile, setIsClickableNextInDialogFile] = useState(true);
    const searchFileQuery: { page: number, size: number, q: string } = {
        page: 0,
        size: 10,
        q: "",
    };
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
    const onClickFile = async () => {
        setIsOpenDialogFile(true);
        await searchAdminFiles();
    }
    const onClickUpload = async (file: File) => {
        console.log(file);
        const r = await apiClient.postAdminFiles(file);
        navigator.clipboard.writeText(generateFileMarkdownURL(r.file));
        setIsOpenDialogFile(false);
    }
    const onClickCopyURLInDialog = async (file: ArticleFile) => {
        navigator.clipboard.writeText(generateFileMarkdownURL(file));
        setIsOpenDialogFile(false);
    }
    async function searchAdminFiles() {
        const r = await apiClient.getAdminFiles(
            searchFileQuery.q,
            searchFileQuery.page,
            searchFileQuery.size,
        );
        setFilesInDialogFile(r.files);
        searchFileQuery.page = r.page;
        searchFileQuery.size = r.size;
        setIsClickablePrevInDialogFile(searchFileQuery.page > 0);
        console.log(r.files.length >= searchFileQuery.size);
        setIsClickableNextInDialogFile(r.files.length >= searchFileQuery.size);
    }
    const onClickSearch = async (q: string) => {
        searchFileQuery.q = q;
        searchFileQuery.page = 0;
        await searchAdminFiles();
    };
    const onClickPrev = async () => {
        searchFileQuery.page--;
        await searchAdminFiles();
    }
    const onClickNext = async () => {
        searchFileQuery.page++;
        await searchAdminFiles();
    }
    return (
        <>
            <DialogFile
                isOpen={isOpenDialogFile}
                isClickableNext={isClickableNextInDialogFile}
                isClickablePrev={isClickablePrevInDialogFile}
                files={filesInDialogFile}
                onClickSearch={onClickSearch}
                onClickCopyURL={onClickCopyURLInDialog}
                onClickUpload={onClickUpload}
                onClickClose={() => {
                    setIsOpenDialogFile(false);
                }}
                onClickNext={onClickNext}
                onClickPrev={onClickPrev}
            ></DialogFile>
            <PageMenu
                article={article}
                selectableTags={selectableTags}
                onClickPublishToggle={onClickPublishToggle}
                onClickAddTag={onClickAddTag}
                onClickDeleteTag={onClickDeleteTag}
                onClickSaveTitle={onClickSaveTitle}
                onClickSaveMarkdown={onClickSaveMarkdown}
                onClickFile={onClickFile}
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