<script lang="ts">
    import { SERVER_URL } from "$lib/config";
    import type { User } from "$types/user";
    import { Icon, MagnifyingGlass } from "svelte-hero-icons";
    import AppBarMenu from "./AppBarMenu.svelte";

    export let user: User | null;

    let search = "";
    let showMenu = false;
    let menuAnchor: HTMLElement | null;

    const toggleMenu = () => {
        showMenu = !showMenu;
    };

    const closeMenu = () => {
        showMenu = false;
    };

    const signIn = () => {
        window.location.href = `${SERVER_URL}/auth/google`;
    };
</script>

<nav class="flex h-16 items-center bg-amber-950 pr-8">
    <a class="w-80 text-center text-2xl text-white" href="/">DecoSavvy</a>
    <div class="relative ml-8 w-96">
        <input
            class="w-full rounded-full bg-white py-2 pl-10"
            placeholder="Search"
            bind:value={search}
        />
        <Icon class="absolute top-0 left-0 ml-2 w-6" src={MagnifyingGlass} />
    </div>
    <div class="grow"></div>
    {#if user}
        <button
            class="h-10 w-10 rounded-full bg-white p-0 hover:shadow"
            on:click={toggleMenu}
            bind:this={menuAnchor}
        >
            <img
                class="rounded-full"
                src={user.avatar}
                alt="avatar"
                width="100%"
                height="100%"
            />
        </button>
        {#if showMenu && menuAnchor}
            <AppBarMenu
                x={menuAnchor.offsetLeft + menuAnchor.offsetWidth / 2}
                close={closeMenu}
            >
                <a href={`${SERVER_URL}/signout/google`}>Sign out</a>
            </AppBarMenu>
        {/if}
    {:else}
        <button
            class="h-10 rounded-full bg-white px-4 text-base hover:shadow"
            on:click={signIn}
        >
            Sign in
        </button>
    {/if}
</nav>
