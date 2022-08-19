<script>
import { session } from "../stores/session";

import Login from "./login.svelte";
import Register from "./register.svelte";

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
                <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                    {#if $session !== ''}
                    <div class="dropdown">
                        <button class="btn btn-secondary dropdown-toggle" type="button" data-bs-toggle="dropdown" aria-expanded="false">
                          Usuario
                        </button>
                        <ul class="dropdown-menu">
                          <li><a class="dropdown-item" href="/settings">Settings</a></li>
                          <li><a class="dropdown-item" href="#" on:click={handle_logout}>Logout</a></li>
                        </ul>
                      </div>
                    {:else}
                    <li class="nav-item">
                        <Login/>
                    </li>
                    <li class="nav-item">
                        <Register/>
                    </li>
                    {/if}
                </ul>
            </div>
        </div>
    </div>
</nav>
