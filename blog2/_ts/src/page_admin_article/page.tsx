export default function Page(): JSX.Element {
    return (
        <>
            <div style={{
                position: "absolute",
                top: 6, height: 150,
                left: 6, right: 6,
                overflow: "scroll",
                border: "1px solid blue",
            }}>
                Header
            </div>
            <div style={{
                position: "absolute",
                top: 156, bottom: 6,
                left: 6, right: 6,
                overflow: "scroll",
                display: "flex",
                flexDirection: "row",
                border: "1px solid green",
            }}>
                <div style={{
                    position: "absolute",
                    top: 0, bottom: 0,
                    width: "49%",
                    overflow: "scroll",
                    border: "3px solid blueviolet"
                }}>Markdown</div>
                <div style={{
                    position: "absolute",
                    top: 0, bottom: 0,
                    width: "49%",
                    right: 0,
                    overflow: "scroll",
                    border: "3px solid blueviolet"
                }}>Preview</div>
            </div>
        </>
    )
}