<script lang="ts">
    import "../app.css";
    import AppBar from "$components/AppBar.svelte";
    import type { LayoutProps } from "./$types";

    let { children, data }: LayoutProps = $props();
</script>

<AppBar user={data.user} />

<div class="h-[calc(100vh-64px)]">
    {@render children()}
</div>
