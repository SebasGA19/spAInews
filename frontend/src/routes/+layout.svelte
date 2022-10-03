<script lang="ts">
	import { session } from '../stores/session';
	import { darkMode } from '../stores/dark-mode';
	import Header from '../components/header.svelte';
	import Footer from '../components/footer.svelte';
	import { onMount } from 'svelte';
	import { accountDetails } from '../api/account';

	onMount(() => {
		if ($session !== '') {
			accountDetails($session)
				.then()
				.catch((error) => {
					alert('Session expired');
					$session = '';
				});
		}
	});
</script>

{#if $darkMode === 'dark'}
	<link rel="stylesheet" href="/vendor/bootstrap-5.2.0-dist/css/bootstrap.min-darkly.css" />
	<link rel="stylesheet" href="/css/dark.css" />
{:else}
	<link rel="stylesheet" href="/vendor/bootstrap-5.2.0-dist/css/bootstrap.min-litera.css" />
	<link rel="stylesheet" href="/css/light.css" />
{/if}
<Header />
<div style="min-height: 80vh; margin-top: 20px; margin-bottom: 20px;">
	<slot />
</div>
<Footer />
