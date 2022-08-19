<script>
import { session } from "../stores/session";
import { darkMode } from "../stores/dark-mode"

import Login from "./login.svelte";
import Register from "./register.svelte";

let darkModeChecked = $darkMode === 'dark';

$ : {
    darkModeChecked;
    $darkMode = darkModeChecked ? 'dark' : 'light';
}
</script>

<nav class="navbar navbar-expand-lg {darkModeChecked ? 'navbar-dark bg-dark' : 'bg-light'}">
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
                    <li class="nav-item align-middle" style="margin: auto; margin-right: 30px;">
                        <div class="form-check form-switch">
                            <label class="form-check-label">
                                {darkModeChecked ? "Dark" : "Light"}
                                <input class="form-check-input" type="checkbox" role="switch" bind:checked="{darkModeChecked}">
                            </label>
                        </div>
                    </li>
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
