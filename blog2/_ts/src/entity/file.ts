
export type FileId = string;
export type FileType = 'image' | 'video' | 'text' | 'unknown';
export interface File {
    id: FileId,
    type: FileType,
    mediaType: string,
    createdAt: string,
    updatedAt: string,
}
export type FileThumbnailId = string;
export interface FileThumbnail {
    id: FileThumbnailId,
    fileId: FileId,
    mediaType: string,
}
export interface FileAndThumbnail {
    file: File;
    fileThumbnail: FileThumbnail | null;
}

export function getThumbnailUrl(a: FileThumbnail | null): string {
    const baseUrl = window.env.baseUrlFileThumbnail;
    if (a === null) {
        return "";
    }
    if (baseUrl.endsWith("/")) {
        return `${baseUrl}${a.id}`;
    }
    return `${baseUrl}/${a.id}`;
}

export function getFileUrl(a: File): string {
    const baseUrl = window.env.baseUrlFile;
    if (baseUrl.endsWith("/")) {
        return `${baseUrl}${a.id}`;
    }
    return `${baseUrl}/${a.id}`;
}

export function generateFileMarkdownURL(file: File): string {
    const url = getFileUrl(file)
    if (file.type === "image") {
        return `![](${url})`;
    }
    return `[${file.id}](${url})`;
}