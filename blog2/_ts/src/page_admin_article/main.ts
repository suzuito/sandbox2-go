import { Snackbar, SnackbarType } from "../component_info_snackbar/main";
import { getElementByIdWithException } from "../lib/dom";

declare global {
    interface Window {
        env: {
            articleId: string;
            published: boolean;
        }
    }
}

document.addEventListener("DOMContentLoaded", () => {
    const snackbar = new Snackbar();
    const docSaveMarkdownButton = getElementByIdWithException("save-markdown");
    docSaveMarkdownButton.addEventListener("click", async () => {
        const markdown = getElementByIdWithException("markdown-editor").innerText;
        const res = await fetch(
            `/admin/articles/${window.env.articleId}/markdown`,
            {
                method: "PUT",
                headers: { 'Content-Type': 'application/json' },
                body: markdown,
            }
        );
        if (res.status === 200) {
            document.location.reload();
            return;
        }

        snackbar.open(SnackbarType.ERROR, "更新エラー", "Markdownの更新に失敗しました");
    });

    const docPublishToggle = getElementByIdWithException("publish-toggle");
    docPublishToggle.addEventListener("click", async () => {
        let message = "";
        let method = "";
        if (window.env.published === true) {
            message = "ドラフトへ戻しますか？";
            method = "DELETE";
        } else {
            message = "公開しますか？";
            method = "POST";
        }
        const result = confirm(message);
        if (!result) {
            return;
        }
        const res = await fetch(
            `/admin/articles/${window.env.articleId}/publish`,
            {
                method,
                headers: { 'Content-Type': 'application/json' },
            }
        );
        if (res.status === 200) {
            document.location.reload();
            return;
        }
        snackbar.open(SnackbarType.ERROR, "更新エラー", "ステータスの更新に失敗しました");
    });

    const docSaveTitleButton = getElementByIdWithException("save-title");
    const docArticleTitle = getElementByIdWithException("article-title") as HTMLInputElement;
    docSaveTitleButton.addEventListener("click", async () => {
        const res = await fetch(
            `/admin/articles/${window.env.articleId}`,
            {
                method: "PUT",
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    title: docArticleTitle.value,
                }),
            }
        );
        if (res.status === 200) {
            document.location.reload();
            return;
        }
        snackbar.open(SnackbarType.ERROR, "更新エラー", "タイトルの更新に失敗しました");
    });
});