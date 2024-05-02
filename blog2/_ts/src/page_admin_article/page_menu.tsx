import { useRef } from "react";
import { Article } from "../entity/article";
import { Tag, TagId } from "../entity/tag";

export default function PageMenu({
    article,
    selectableTags,
    onClickPublishToggle,
    onClickDeleteTag,
    onClickAddTag,
    onClickSaveTitle,
    onClickSaveMarkdown,
    onClickFile,
}: {
    article: Article,
    selectableTags: Tag[],
    onClickPublishToggle: () => void,
    onClickDeleteTag: (tagId: TagId) => void,
    onClickAddTag: (tagId: TagId) => void,
    onClickSaveTitle: (newTitle: string) => void,
    onClickSaveMarkdown: () => void,
    onClickFile: () => void,
}) {
    let currentSelectedTagId: string = ""
    const refInputTitle = useRef<HTMLInputElement>(null);
    return (
        <div style={{
            position: "absolute",
            top: 6, height: 150,
            left: 6, right: 6,
            overflow: "scroll",
            // border: "1px solid blue",
        }}>
            <div style={{
                position: "relative",
                height: 24,
                // border: "1px solid red"
            }}>
                <a href="./">記事一覧へ</a>
            </div>
            <div style={{
                position: "relative",
                height: 24,
                // border: "1px solid red"
            }}>
                <button
                    onClick={() => {
                        onClickPublishToggle()
                    }}
                >{article.published ? "公開中" : "ドラフト"}</button>
                {article.published ?
                    (<span>公開日 {article.publishedAtAsDate?.toLocaleString()}</span>) : (<></>)
                }
            </div>
            <div style={{
                position: "relative",
                height: 24,
                // border: "1px solid red"
            }}>
                <input
                    type="text"
                    style={{ width: "80%" }}
                    defaultValue={article.title}
                    ref={refInputTitle}
                />
                <button onClick={() => {
                    if (refInputTitle.current === null) {
                        return;
                    }
                    onClickSaveTitle(refInputTitle.current.value);
                }}>
                    保存
                </button>
            </div>
            <div style={{
                position: "relative",
                height: 24,
                overflow: "scroll",
                whiteSpace: "nowrap",
                border: "1px solid red",
            }}>
                {article.tags.map((tag) => {
                    return (
                        <span style={{
                            border: "1px solid black",
                            marginRight: 6
                        }} key={tag.id}>
                            {tag.name}
                            <button onClick={() => {
                                onClickDeleteTag(tag.id);
                            }}>x</button>
                        </span>
                    )
                })}
            </div>
            <div style={{
                position: "relative",
                height: 24,
                overflow: "scroll",
                border: "1px solid red",
            }}>
                <select onInput={(event) => {
                    const selectedOption = event.currentTarget.selectedOptions[0]
                    if (selectedOption === undefined) {
                        return;
                    }
                    currentSelectedTagId = selectedOption.value;
                }}>
                    <option key="" value="">タグを選択する</option>
                    {selectableTags.map((tag) => {
                        return (
                            <option
                                key={tag.id}
                                value={tag.id}>{tag.name}
                            </option>
                        )
                    })}
                </select>
                <button
                    onClick={() => {
                        if (currentSelectedTagId === "") {
                            return;
                        }
                        onClickAddTag(currentSelectedTagId);
                    }}
                >
                    タグを追加する
                </button>
            </div>
            <div style={{
                position: "relative",
                height: 24,
                // border: "1px solid red"
            }}>
                <button onClick={() => {
                    onClickFile()
                }}>ファイルをアップロードする</button>
                <button onClick={() => {
                    onClickSaveMarkdown()
                }}>Markdownを保存する</button>
            </div>
        </div>
    )
}