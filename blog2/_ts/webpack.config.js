const path = require('path');

// 参考にさせていただきました
// https://ics.media/entry/16329/
module.exports = {
    entry: {
        'page_admin_article': './src/page_admin_article/main.ts',
        'component_info_snackbar': './src/component_info_snackbar/main.ts',
        'component_image_modal': './src/component_image_modal/main.ts',
    },
    module: {
        rules: [
            {
                // 拡張子 .ts の場合
                test: /\.ts$/,
                // TypeScript をコンパイルする
                use: 'ts-loader',
            },
        ],
    },
    // import 文で .ts ファイルを解決するため
    // これを定義しないと import 文で拡張子を書く必要が生まれる。
    // フロントエンドの開発では拡張子を省略することが多いので、
    // 記載したほうがトラブルに巻き込まれにくい。
    resolve: {
        // 拡張子を配列で指定
        extensions: [
            '.ts', '.js',
        ],
    },
    output: {
        path: path.resolve(__dirname, '..', 'internal', 'web', '_js'),
        // filename: 'bundle.js',
    },
};