interface Article {
    title: string;
    description: string;
    maintext: string;
    authors: string[];
    category: string;
    date_publish: string;
    source_domain: string;
    url: string;
};

interface Account {
    "username": string;
    "email": string;
    "creation-date": string;
};

interface News {
    "current-page": number;
    "number-of-results": number;
    "news": Article[];
}

interface Words {
    automatic: boolean;
    words: string[];
}

interface SearchFilter {
    "start-date"?: string,
    "end-date"?: string,
    "sources"?: string[],
    "maintext-words"?: string[],
    "title-words"?: string[],
    "old-first"?: boolean
}

export type { Article, Account, News, Words, SearchFilter };
