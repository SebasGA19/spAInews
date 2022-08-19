<script>
import { darkMode } from "../stores/dark-mode";
import Footer from "./footer.svelte";
import Navbar from "./navbar.svelte";

</script>

{#if $darkMode === 'dark'}
<link rel="stylesheet" href="/vendor/bootstrap-5.2.0-dist/css/bootstrap.min-darkly.css"/>
<link rel="stylesheet" href="/css/dark.css"/>
{:else}
<link rel="stylesheet" href="/vendor/bootstrap-5.2.0-dist/css/bootstrap.min-litera.css"/>
<link rel="stylesheet" href="/css/light.css"/>
{/if}
<Navbar/>
<slot></slot>
<Footer/>