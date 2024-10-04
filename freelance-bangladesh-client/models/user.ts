interface User {
    ID: string;
    Email: string;
    Phone?: string;
    IsVerified: boolean;
    Name: string;
}

export interface Talent extends User {}
export interface Client extends User {}
