<script lang="ts">
    import type { Product } from "$types/product";
    import { Icon, PlusCircle } from "svelte-hero-icons";

    let { id, image, name, description, price }: Product = $props();

    const addToCart = () => {
        console.log("Add " + id);
    };
</script>

<div class="flex w-64 flex-col gap-2 bg-red-100 p-4">
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
