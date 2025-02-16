<script lang="ts">
    import type { PageProps } from "./$types";

    let { data }: PageProps = $props();
</script>

<div class="px-8 py-4">
    <ol class="flex list-disc flex-col gap-2">
        {#each data.cart as name}
            <li class="text-lg">{name}</li>
        {/each}
    </ol>
</div>
