<script lang="ts">
    import { SERVER_URL } from "$lib/config";
    import { Icon, Trash } from "svelte-hero-icons";
    import type { PageProps } from "./$types";

    let { data }: PageProps = $props();

    const removeFromCart = async (productId: string) => {
        await fetch(`${SERVER_URL}/api/remove-from-cart`, {
            credentials: "include",
            method: "POST",
            body: productId,
        });

        location.reload();
    };
</script>

<div class="p-4">
    <h1 class="mb-4 text-2xl font-bold">My Cart</h1>
    <div class="mb-4 flex flex-col gap-2">
        {#each data.products as product}
            <div class="flex items-center gap-8 text-lg">
                <p>{product.name}: {product.price}</p>
                <button class="h-8" onclick={() => removeFromCart(product.id)}>
                    <Icon src={Trash} />
                </button>
            </div>
        {/each}
    </div>
    <p class="text-xl">
        Total: ${data.products
            .reduce((acc, product) => acc + product.price, 0)
            .toFixed(2)}
    </p>
</div>
