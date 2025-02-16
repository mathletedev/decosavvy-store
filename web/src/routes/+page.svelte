<script lang="ts">
    import ProductCard from "$components/ProductCard.svelte";
    import type { PageProps } from "./$types";

    let { data }: PageProps = $props();
</script>

<div class="flex gap-4 p-4">
    {#each data.products as product}
        <ProductCard {...product} />
    {/each}
</div>
