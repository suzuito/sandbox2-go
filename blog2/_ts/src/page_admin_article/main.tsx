import { getElementByIdWithException } from "../lib/dom";
import { createRoot } from "react-dom/client";
import Page from "./page";
import { Article, setupArticle } from "../entity/article";
import { Tag } from "../entity/tag";
import { newAPIClient } from "../infra/gateway/client";

declare global {
    interface Window {
        env: {
            article: Article,
            notAttachedTags: Tag[],
            html: string,
            markdown: string,
            baseUrlFile: string,
            baseUrlFileThumbnail: string,
        }
    }
}

document.addEventListener("DOMContentLoaded", () => {
    setupArticle(window.env.article);
    const reactAppRoot = createRoot(getElementByIdWithException("react-app"));
    const apiClient = newAPIClient();
    reactAppRoot.render(
        <Page
            defaultArticle={window.env.article}
            defaultSelectableTags={window.env.notAttachedTags}
            defaultMarkdown={window.env.markdown}
            defaultHTML={window.env.html}
            apiClient={apiClient}
        ></Page>);
});
