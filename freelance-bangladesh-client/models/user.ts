interface User {
    id: string;
    email: string;
    phone?: string;
    is_verified: boolean;
    name: string;
}

export interface Talent extends User {}
export interface Client extends User {}
