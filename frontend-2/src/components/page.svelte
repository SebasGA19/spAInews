<script>
import Footer from "./footer.svelte";
import Navbar from "./navbar.svelte";

let darkMode = false;
</script>

<button on:click="{() => darkMode = !darkMode}">{darkMode ? "Light" : "Dark"}</button>
{#if darkMode}
<link rel="stylesheet" href="/vendor/bootstrap-5.2.0-dist/css/bootstrap.min-slate.css"/>
{:else}
<link rel="stylesheet" href="/vendor/bootstrap-5.2.0-dist/css/bootstrap.min.css"/>
{/if}
<Navbar bind:darkMode/>
<slot></slot>
<Footer/>