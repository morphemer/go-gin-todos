export interface IUser {
    ID: string | number
    Username: string
    Email: string
    UpdatedAt: string | Date
    CreatedAt: string | Date
    Articles?: IArticle[]
}

export interface IArticle {
    ID: string | number
    Title: string
    Content: string
    UpdatedAt: string | Date
    CreatedAt: string | Date
}
