import { error } from "@sveltejs/kit";
import { confirmRegister } from "../../../../api/confirm";

export const ssr = false;
export const csr = true;
export async function load({ params }) {
    try {
        const result: void = await confirmRegister(params.code);
        return;
    } catch (e) {
        console.log(e)
    }
    throw error(403, 'Invalid confirmation code')
}