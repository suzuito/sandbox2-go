import { getElementByIdWithException } from "../lib/dom";

class Docs {
    constructor(
        public container = getElementByIdWithException("info-snackbar"),
        public title = getElementByIdWithException("info-snackbar-title"),
        public description = getElementByIdWithException("info-snackbar-description"),
        public buttonOK = getElementByIdWithException("info-snackbar-button-ok"),
    ) { }
}

export class Snackbar {
    private docs: Docs | undefined;
    constructor() {
        try {
            this.docs = new Docs();
        } catch (e: unknown) {
            this.docs = undefined;
            console.warn(e);
            return;
        }

        this.docs.buttonOK.addEventListener("click", () => {
            if (this.docs === undefined) {
                return;
            }
            this.docs.container.style.display = "none";
        });
    }

    open(
        type: SnackbarType,
        title: string,
        description: string,
    ) {
        if (this.docs === undefined) {
            return;
        }
        this.docs.container.style.display = "block";
        switch (type) {
            case SnackbarType.WARN:
                this.docs.container.style.backgroundColor = "yellow";
            case SnackbarType.ERROR:
                this.docs.container.style.backgroundColor = "rgb(200,9,0)";
            default:
                this.docs.container.style.backgroundColor = "#fff";
        }
        this.docs.title.innerText = title;
        this.docs.description.innerText = description;
    }
}

export enum SnackbarType {
    INFO,
    WARN,
    ERROR,
}
