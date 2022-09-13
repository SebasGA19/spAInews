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
export type { Article, Account, News, Words };
