import {error} from "@sveltejs/kit";
import {confirmEmail} from "../../../../api/confirm";

export const ssr = false;
export const csr = true;

export async function load({params}) {
    try {
        const result: void = await confirmEmail(params.code);
        return;
    } catch (e) {
        console.log(e)
    }
    throw error(403, 'Invalid confirmation code or invalid password in email submit')
}