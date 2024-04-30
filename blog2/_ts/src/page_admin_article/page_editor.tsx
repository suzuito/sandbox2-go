import { useEffect, useRef } from "react";
import { getElementByIdWithException } from "../lib/dom";



export default function PageEditor({
    markdown,
}: {
    markdown: string,
}) {
    const divEditor = useRef<HTMLDivElement>(null);
    useEffect(() => {
        if (divEditor.current === null) {
            throw new Error("divEditor.current is null")
        }
        setMarkdownToDivEditor(markdown, divEditor.current);
        appendBrRow(divEditor.current);
    });
    return (
        <div
            id="markdown-editor"
            style={{
                position: "absolute",
                top: 0, bottom: 0,
                width: "49%",
                overflow: "scroll",
                border: "3px solid blueviolet"
            }}
            contentEditable={true}
            ref={divEditor}
            suppressContentEditableWarning={true} // https://mykii.blog/react-warning-contenteditable/
            onInput={() => {
                if (divEditor.current === null) {
                    return;
                }
                modifyDivEditorAfterInput(divEditor.current);
            }}
        >
        </div >
    )
}

export function extractCurrentMarkdown() {
    const e = getElementByIdWithException("markdown-editor");
    const lines = e.querySelectorAll(".cm-line");
    let current: string[] = [];
    for (let i = 0; i < lines.length; i++) {
        const line = lines.item(i) as HTMLElement;
        if (line.innerText === "\n") {
            current.push("");
            continue;
        }
        current.push(line.innerText);
    }
    console.log(current);
    return current.join("\n");
}

function setMarkdownToDivEditor(markdown: string, e: HTMLDivElement) {
    while (e.firstChild) {
        e.removeChild(e.firstChild);
    }
    markdown.split("\n").forEach(l => {
        const cmline = document.createElement("div");
        cmline.className = "cm-line";
        cmline.innerText = l;
        if (l === "") { // 改行のみの行
            cmline.appendChild(document.createElement("br"));
        }
        e.appendChild(cmline);
    });
}

function modifyDivEditorAfterInput(e: HTMLDivElement) {
    appendBrRow(e);
    mergeNestedDivTags(e);
    removeEmptyRow(e);
}

const emptyDiv = (() => {
    const emptyDiv = document.createElement("div");
    emptyDiv.className = "cm-line";
    emptyDiv.appendChild(document.createElement("br"));
    return emptyDiv;
})()

function appendBrRow(e: HTMLDivElement) {
    // エディターに何も入力されていない場合
    // <div class="cm-line"><br/></div>
    // となるように
    if (e.querySelectorAll(".cm-line").length <= 0) {
        e.appendChild(emptyDiv);
        return;
    }
}

function mergeNestedDivTags(e: HTMLDivElement) {
    while (mergeNestedDivTagsCore(e)) { }
}
function mergeNestedDivTagsCore(e: HTMLDivElement): boolean {
    // console.log("M1====")
    // コピーペーストした際に
    // <div class="cm-line">
    //   <div class="cm-line">
    //     hoge
    //   </div>
    // </div>
    // みたいな感じで入れ子構造になってしまうので
    // それを解きほぐす
    for (let i = 0; i < e.children.length; i++) {
        const cmline = e.children.item(i);
        if (cmline === null) {
            break;
        }
        const nestedElems = cmline.querySelectorAll("*")
        // console.log(nestedElems.length, cmline);
        if (nestedElems.length <= 0) {
            continue;
        }
        // 入れ子構造になってしまっているHTMLがある
        if (nestedElems.length === 1 && nestedElems.item(0).tagName.toLowerCase() === "br") {
            // brタグ一個だけなら入れ子になっててもOK
            continue;
        }
        for (let j = 0; j < nestedElems.length; j++) {
            nestedElems.item(j).remove();
        }
        for (let j = 0; j < nestedElems.length; j++) {
            const nestedCMLine = nestedElems.item(j);
            cmline.insertAdjacentElement(
                "afterend",
                nestedCMLine,
            );
        }
        return true;
    }
    return false;
}

function removeEmptyRow(e: HTMLDivElement) {
    // console.log("M2====")
    // なんか知らんが空のdivタグが入ってしまうので
    // 削除する。
    const deletedElems: Element[] = [];
    for (let i = 0; i < e.children.length; i++) {
        const cmline = e.children.item(i);
        if (cmline === null) {
            break;
        }
        if (cmline.children.length > 0) {
            continue;
        }
        if (cmline.textContent !== null && cmline.textContent.length <= 0) {
            deletedElems.push(cmline);
        }
    }
    deletedElems.forEach(e => e.remove());
}