import { SERVER_URL } from "$lib/config";
import type { Product } from "$types/product";
import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch }) => {
    const res = await fetch(`${SERVER_URL}/api/cart`, {
        credentials: "include",
    });

    const data = await res.json();

    return {
        products: data as Product[],
    };
};
