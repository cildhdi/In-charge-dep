export interface UserInfo {
    ID: number,
    CreatedAt: string,
    UpdateAt: string,
    DeleteAt: string | null,
    Phone: string,
    Name: string | null,
    Role: number,
    Banned: number,
    Points: number,
    AvatarUrl: string
}
