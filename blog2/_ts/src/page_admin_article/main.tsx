import { getElementByIdWithException } from "../lib/dom";
import { createRoot } from "react-dom/client";
import Page from "./page";

declare global {
    interface Window {
        env: {
            articleId: string;
            published: boolean;
        }
    }
}

document.addEventListener("DOMContentLoaded", () => {
    const reactAppRoot = createRoot(getElementByIdWithException("react-app"));
    reactAppRoot.render(Page());
});
