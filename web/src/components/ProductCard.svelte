<script lang="ts">
    import { SERVER_URL } from "$lib/config";
    import type { Product } from "$types/product";
    import { Icon, PlusCircle } from "svelte-hero-icons";

    interface Props extends Product {
        signedIn: boolean;
    }

    let { id, image, name, description, price, signedIn }: Props = $props();

    const addToCart = async () => {
        if (!signedIn) {
            alert("Please sign in!");
            return;
        }

        await fetch(`${SERVER_URL}/api/add-to-cart`, {
            credentials: "include",
            method: "POST",
            body: id,
        });

        alert(`Added ${name} to cart!`);
    };
</script>

<div class="flex w-64 flex-col gap-2 bg-stone-200 p-4">
    <img src={image} alt="item" />
    <h3 class="text-xl font-bold text-amber-950">{name}</h3>
    <p>{description}</p>
    <div class="flex-grow"></div>
    <div class="flex justify-between">
        <p class="italic underline">${price.toFixed(2)}</p>
        <button class="h-8" onclick={addToCart}>
            <Icon src={PlusCircle} />
        </button>
    </div>
</div>
