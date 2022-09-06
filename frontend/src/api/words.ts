import type { Words } from "src/common/types";
import { api_baseurl } from "../config";

export async function updateWords(words: Words, session: string) {
    const response = await fetch(
        api_baseurl + "/words",
        {
            method: 'POST',
            mode: 'cors',
            cache: 'no-cache',
            headers: {
                'Session': session,
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(words)
        }
    );
    if (response.status === 200){
        return;
    }
    const data = await response.json();
    throw data["message"];
}

export async function getWords(session: string): Promise<Words> {
    const response = await fetch(
        api_baseurl + "/words",
        {
            method: 'GET',
            mode: 'cors',
            cache: 'no-cache',
            headers: {
                'Session': session
            }
        }
    );
    const data = await response.json();
    if (response.status === 200){
        return data;
    }
    throw data["message"];
}