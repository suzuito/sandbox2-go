import { getElementByIdWithException } from "../lib/dom";



export default function PageEditor({
    markdown,
}: {
    markdown: string,
}) {
    return (
        <textarea
            id="markdown-editor"
            style={{
                position: "absolute",
                top: 0, bottom: 0,
                width: "49%",
                overflow: "scroll",
                border: "3px solid blueviolet"
            }}
            defaultValue={markdown}
        >
        </textarea >
    )
}

export function extractCurrentMarkdown() {
    const e = getElementByIdWithException("markdown-editor") as HTMLTextAreaElement;
    return e.value;
}