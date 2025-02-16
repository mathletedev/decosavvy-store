<script lang="ts">
    import ProductCard from "$components/ProductCard.svelte";
    import type { PageProps } from "./$types";

    let { data }: PageProps = $props();
</script>

<div class="p-4">
    <h1 class="mb-4 text-2xl font-bold">Products</h1>
    <div class="flex gap-4">
        {#each data.products as product}
            <ProductCard {...product} />
        {/each}
    </div>
</div>
