import type { News, SearchFilter } from "src/common/types";
import { api_baseurl } from "../config";

export async function news(pages: number): Promise<News> {

    const response = await fetch(
        api_baseurl + `/news/${pages}`,
        {
            method: 'GET',
            mode: 'cors',
            cache: 'no-cache'
        }
    )
    const data = await response.json();
    if (response.status === 200) {
        return data;
    }
    throw data["message"];
}

export async function search(page: number, searchFilter: SearchFilter): Promise<News> {
    const response = await fetch(
        api_baseurl + `/search/${page}`,
        {
            method: 'POST',
            mode: 'cors',
            cache: 'no-cache',
            headers: {
                'Content-Type': 'application/json;charset=utf-8'
            },
            body: JSON.stringify(searchFilter)
        }
    )
    const data = await response.json();
    if (response.status === 200) {
        return data;
    }
    throw data["message"];
}