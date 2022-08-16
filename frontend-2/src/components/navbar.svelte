<script>
import { session } from "../stores/session";

import Login from "./login.svelte";

export let darkMode = false;

</script>

<nav class="navbar navbar-expand-lg {darkMode ? 'bg-primary navbar-dark' : 'bg-light'}">
    <div class="container-fluid">
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarNav">
            <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                <li class="nav-item">
                    <a href="/" class="nav-link">spAInews</a>
                </li>
            </ul>
            <div class="d-flex" role="search">
                {#if $session !== ''}
                <a href="/account" class="nav-link">Account</a>
                {:else}
                <Login/>
                <a href="/register" class="nav-link">Register</a>
                {/if}
            </div>
        </div>
    </div>
</nav>