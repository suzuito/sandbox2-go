import { useRef } from "react";
import { File as ArticleFile, FileAndThumbnail, getFileUrl, getThumbnailUrl } from "../entity/file";

export default function DialogFile({
    isOpen,
    isClickablePrev,
    isClickableNext,
    files,
    onClickSearch,
    onClickCopyURL,
    onClickClose,
    onClickUpload,
    onClickNext,
    onClickPrev,
}: {
    isOpen: boolean,
    isClickablePrev: boolean,
    isClickableNext: boolean,
    files: FileAndThumbnail[],
    onClickSearch: (q: string) => void,
    onClickNext: () => void,
    onClickPrev: () => void,
    onClickCopyURL: (file: ArticleFile) => void,
    onClickClose: () => void,
    onClickUpload: (file: File) => void,
}) {
    const refInputFile = useRef<HTMLInputElement>(null);
    const refInputQuery = useRef<HTMLInputElement>(null);
    return (
        <dialog style={{
            zIndex: 100,
        }} open={isOpen}>
            <div>
                <input type="text" ref={refInputQuery}></input>
                <button onClick={() => {
                    if (refInputQuery.current === null) {
                        return;
                    }
                    onClickSearch(refInputQuery.current.value);
                    console.log(refInputQuery.current.value);
                }}>検索する</button>
            </div>
            <div>
                <button disabled={!isClickablePrev} onClick={() => { onClickPrev() }}>前へ</button>
                <button disabled={!isClickableNext} onClick={() => { onClickNext() }}>次へ</button>
            </div>
            <div>
                {files.map(file => {
                    return (
                        <div style={{
                            display: "flex",
                            flexDirection: "row",
                        }} key={file.file.id}>
                            <img width={50} height={50} src={getThumbnailUrl(file.fileThumbnail)}></img>
                            <div>
                                <button onClick={() => {
                                    onClickCopyURL(file.file);
                                }}>URLをコピーする</button>
                            </div>
                            <div>
                                <a href={getFileUrl(file.file)} target="_blank">{file.file.id}</a>
                            </div>
                        </div>
                    )
                })}
            </div>
            <div>
                <input
                    type="file"
                    ref={refInputFile}
                ></input>
                <button onClick={() => {
                    if (refInputFile.current === null) {
                        return;
                    }
                    if (refInputFile.current.files === null) {
                        return;
                    }
                    const file = refInputFile.current.files.item(0);
                    if (file === null) {
                        return;
                    }
                    onClickUpload(file);
                }}>アップロード</button>
            </div>
            <div>
                <button onClick={() => {
                    onClickClose();
                }}>閉じる</button>
            </div>
        </dialog>
    )
}