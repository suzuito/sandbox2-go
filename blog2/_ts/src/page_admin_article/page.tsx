
export default function Page(): JSX.Element {
    return (
        <div>
            hoge
            <h1>foo</h1>
            {b()}
        </div>
    )
}

function b(): JSX.Element {
    return (
        <a href="https://www.example.com">aaaa</a>
    )
}