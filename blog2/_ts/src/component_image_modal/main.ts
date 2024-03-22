import { getElementByIdWithException } from "../lib/dom";

class Docs {
    constructor(
        public container = getElementByIdWithException("image-modal"),
        public imageInput: HTMLInputElement = getElementByIdWithException("image-modal-input") as HTMLInputElement,
        public uploadButton = getElementByIdWithException("image-modal-upload-button"),
        public closeButton = getElementByIdWithException("image-modal-close-button"),
    ) { }
}

export class ImageModal {
    private docs: Docs | undefined;
    private articleId: string | undefined;
    constructor() {
        try {
            this.docs = new Docs();
        } catch (e: unknown) {
            this.docs = undefined;
            console.warn(e);
            return;
        }
        this.docs.uploadButton.addEventListener("click", async () => {
            if (this.docs === undefined) {
                return;
            }
            if (this.docs.imageInput.files === null) {
                return;
            }
            if (this.docs.imageInput.files.length <= 0) {
                return;
            }
            if (this.articleId === undefined) {
                return;
            }
            const file = this.docs.imageInput.files[0];
            const formData = new FormData();
            formData.append("image", file);
            const res = await fetch(
                `/admin/articles/${this.articleId}/image`,
                {
                    method: "POST",
                    body: formData,
                },
            );
        });
        this.docs.closeButton.addEventListener("click", () => {
            this.close();
        });
    }
    async open(articleId: string): Promise<void> {
        if (this.docs === undefined) {
            return;
        }
        this.articleId = articleId;
        this.docs.container.style.display = "block";
    }
    async close(): Promise<void> {
        if (this.docs === undefined) {
            return;
        }
        this.docs.container.style.display = "none";
    }
}