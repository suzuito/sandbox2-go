export class HTMLElementNotFoundError extends Error {
    constructor(id: string) {
        super(`element is not found #${id}`);
    }
}

export function getElementByIdWithException(id: string): HTMLElement {
    const e = document.getElementById(id)
    if (e === null) {
        throw new HTMLElementNotFoundError(id);
    }
    return e;
}